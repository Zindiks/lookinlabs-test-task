package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Zindiks/lookinlabs-test-task/controller"
	"github.com/Zindiks/lookinlabs-test-task/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Мок для userOperations
type MockUserService struct {
    mock.Mock
}

func (m *MockUserService) CreateUser(user *model.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func (m *MockUserService) GetUserByID(id string) (*model.User, error) {
    args := m.Called(id)
    if user, ok := args.Get(0).(*model.User); ok {
        return user, args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *MockUserService) GetUsers() ([]model.User, error) {
    args := m.Called()
    if users, ok := args.Get(0).([]model.User); ok {
        return users, args.Error(1)
    }
    return nil, args.Error(1)
}

func (m *MockUserService) UpdateUser(user *model.User) error {
    args := m.Called(user)
    return args.Error(0)
}

func TestCreateUser(t *testing.T) {
    gin.SetMode(gin.TestMode)

    t.Run("success case", func(t *testing.T) {
        mockUserOps := new(MockUserService)
        userCtrl := controller.NewUserController(mockUserOps) 

        user := model.User{
            Name:  "TestUser",
            Email: "testuser@example.com",
        }

        mockUserOps.On("CreateUser", &user).Return(nil)

        reqBody, _ := json.Marshal(user)
        req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        c, _ := gin.CreateTestContext(w)
        c.Request = req

        userCtrl.CreateUser(c)

        assert.Equal(t, http.StatusOK, w.Code)

        var respUser model.User
        err := json.Unmarshal(w.Body.Bytes(), &respUser)
        assert.NoError(t, err)
        assert.Equal(t, user.Name, respUser.Name)
        assert.Equal(t, user.Email, respUser.Email)

        mockUserOps.AssertExpectations(t)
    })

    t.Run("invalid JSON", func(t *testing.T) {
        mockUserOps := new(MockUserService)
        userCtrl := controller.NewUserController(mockUserOps)

        req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBufferString("invalid JSON"))
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        c, _ := gin.CreateTestContext(w)
        c.Request = req

        userCtrl.CreateUser(c)

        assert.Equal(t, http.StatusBadRequest, w.Code)
        var resp map[string]string
        err := json.Unmarshal(w.Body.Bytes(), &resp)
        assert.NoError(t, err)
        assert.Contains(t, resp["error"], "invalid character")
    })

    t.Run("create user fails", func(t *testing.T) {
        mockUserOps := new(MockUserService)
        userCtrl := controller.NewUserController(mockUserOps)

        user := model.User{
            Name:  "TestUser",
            Email: "testuser@example.com",
        }

        mockUserOps.On("CreateUser", &user).Return(errors.New("failed to save user"))

        reqBody, _ := json.Marshal(user)
        req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        c, _ := gin.CreateTestContext(w)
        c.Request = req

        userCtrl.CreateUser(c)

        assert.Equal(t, http.StatusInternalServerError, w.Code)
        var resp map[string]string
        err := json.Unmarshal(w.Body.Bytes(), &resp)
        assert.NoError(t, err)
        assert.Equal(t, "Failed to save user", resp["error"])

        mockUserOps.AssertExpectations(t)
    })

    t.Run("missing required fields", func(t *testing.T) {
        mockUserOps := new(MockUserService)
        userCtrl := controller.NewUserController(mockUserOps)

        reqBody := []byte(`{}`)
        req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(reqBody))
        req.Header.Set("Content-Type", "application/json")

        w := httptest.NewRecorder()
        c, _ := gin.CreateTestContext(w)
        c.Request = req

        userCtrl.CreateUser(c)

        assert.Equal(t, http.StatusBadRequest, w.Code)
        var resp map[string]string
        err := json.Unmarshal(w.Body.Bytes(), &resp)
        assert.NoError(t, err)
        assert.Contains(t, resp["error"], "required")
    })
}


func TestGetUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("user found", func(t *testing.T) {
		mockUserOps := new(MockUserService)
		userCtrl := controller.NewUserController(mockUserOps)

		user := &model.User{
			Model: gorm.Model{ID: 1},
			Name:  "TestUser",
			Email: "testuser@example.com",
		}

		mockUserOps.On("GetUserByID", "1").Return(user, nil)

		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "1"}}
		c.Request = req


		userCtrl.GetUser(c)


		assert.Equal(t, http.StatusOK, w.Code)

		var respUser model.User
		err := json.Unmarshal(w.Body.Bytes(), &respUser)
		assert.NoError(t, err)
		assert.Equal(t, user.Name, respUser.Name)
		assert.Equal(t, user.Email, respUser.Email)

		mockUserOps.AssertExpectations(t)
	})

	t.Run("user not found", func(t *testing.T) {
		mockUserOps := new(MockUserService)
		userCtrl := controller.NewUserController(mockUserOps)


		mockUserOps.On("GetUserByID", "2").Return(nil, gorm.ErrRecordNotFound)

		req := httptest.NewRequest(http.MethodGet, "/users/2", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "2"}}
		c.Request = req

		userCtrl.GetUser(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "User not found", resp["error"])

		mockUserOps.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockUserOps := new(MockUserService)
		userCtrl := controller.NewUserController(mockUserOps)

		mockUserOps.On("GetUserByID", "3").Return(nil, errors.New("database error"))

		req := httptest.NewRequest(http.MethodGet, "/users/3", nil)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Params = gin.Params{{Key: "id", Value: "3"}}
		c.Request = req

		userCtrl.GetUser(c)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		var resp map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "Failed to fetch user", resp["error"])

		mockUserOps.AssertExpectations(t)
	})
}


// func TestGetUsers(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	mockUserOps := new(MockUserService)
// 	userCtrl := controller.NewUserController(mockUserOps)

// 	t.Run("users_found", func(t *testing.T) {
// 		users := []model.User{
// 			{Name: "User1", Email: "user1@example.com"},
// 			{Name: "User2", Email: "user2@example.com"},
// 		}

// 		mockUserOps.On("GetUsers").Return(users, nil)

// 		req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req

// 		userCtrl.GetUsers(c)

// 		assert.Equal(t, http.StatusOK, w.Code)

// 		var response []model.User
// 		err := json.Unmarshal(w.Body.Bytes(), &response)
// 		assert.NoError(t, err)
// 		assert.Equal(t, users, response)

// 		mockUserOps.AssertExpectations(t)
// 	})

// 	t.Run("no_users_found", func(t *testing.T) {
// 		mockUserOps.On("GetUsers").Return([]model.User{}, nil)

// 		req := httptest.NewRequest(http.MethodGet, "/users", nil)
// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req

// 		userCtrl.GetUsers(c)

// 		assert.Equal(t, http.StatusOK, w.Code)

// 		var response []model.User
// 		err := json.Unmarshal(w.Body.Bytes(), &response)
// 		assert.NoError(t, err)
// 		assert.Empty(t, response)

// 		mockUserOps.AssertExpectations(t)
// 	})

// }


// func TestUpdateUser(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("successful update", func(t *testing.T) {
// 		mockUserOps := new(MockUserService)
// 		userCtrl := controller.NewUserController(mockUserOps)

// 		userID := "1"
// 		updatedUser := model.User{Name: "UpdatedName", Email: "updatedemail@example.com"}
// 		userJSON, _ := json.Marshal(updatedUser)

// 		mockUserOps.On("UpdateUser", userID, updatedUser).Return(nil)

// 		req := httptest.NewRequest(http.MethodPut, "/user/"+userID, bytes.NewBuffer(userJSON))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req
// 		c.Params = gin.Params{{Key: "userID", Value: userID}}

// 		userCtrl.UpdateUser(c)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		var resp map[string]string
// 		err := json.Unmarshal(w.Body.Bytes(), &resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "User updated successfully", resp["message"])

// 		mockUserOps.AssertExpectations(t)
// 	})

// 	t.Run("user not found", func(t *testing.T) {
// 		mockUserOps := new(MockUserService)
// 		userCtrl := controller.NewUserController(mockUserOps)

// 		userID := uint(999) 
// 		updatedUser := model.User{Name: "NonExistentUser", Email: "nonexistent@example.com"}
// 		userJSON, _ := json.Marshal(updatedUser)

// 		mockUserOps.On("UpdateUser", userID, updatedUser).Return(gorm.ErrRecordNotFound)

// 		req := httptest.NewRequest(http.MethodPut, "/user/"+string(rune(userID)), bytes.NewBuffer(userJSON))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req
// 		c.Params = gin.Params{{Key: "userID", Value: string(rune(userID))}}

// 		userCtrl.UpdateUser(c)

// 		assert.Equal(t, http.StatusNotFound, w.Code)
// 		var resp map[string]string
// 		err := json.Unmarshal(w.Body.Bytes(), &resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "User not found", resp["error"])

// 		mockUserOps.AssertExpectations(t)
// 	})

// 	t.Run("internal server error", func(t *testing.T) {
// 		mockUserOps := new(MockUserService)
// 		userCtrl := controller.NewUserController(mockUserOps)

// 		userID := uint(1)
// 		updatedUser := model.User{Name: "ServerErrorUser", Email: "servererror@example.com"}
// 		userJSON, _ := json.Marshal(updatedUser)

// 		mockUserOps.On("UpdateUser", userID, updatedUser).Return(errors.New("database error"))

// 		req := httptest.NewRequest(http.MethodPut, "/user/"+string(rune(userID)), bytes.NewBuffer(userJSON))
// 		req.Header.Set("Content-Type", "application/json")
// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)
// 		c.Request = req
// 		c.Params = gin.Params{{Key: "userID", Value: string(rune(userID))}}

// 		userCtrl.UpdateUser(c)

// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		var resp map[string]string
// 		err := json.Unmarshal(w.Body.Bytes(), &resp)
// 		assert.NoError(t, err)
// 		assert.Equal(t, "Failed to update user", resp["error"])

// 		mockUserOps.AssertExpectations(t)
// 	})
// }