// package tests

// import (
// 	"testing"

// 	"github.com/Zindiks/lookinlabs-test-task/model"
// 	"github.com/Zindiks/lookinlabs-test-task/service"
// 	"github.com/stretchr/testify/assert"
// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"

// 	"strconv"
	
// )

// func setupDatabase() (*gorm.DB, error) {
//     return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// }

// // TestUserService_CreateUser tests the CreateUser method
// func TestUserService_CreateUser(t *testing.T) {
//     db, err := setupDatabase()
//     if err != nil {
//         t.Fatalf("failed to connect to database: %v", err)
//     }

//     // Migrate the schema
//     db.AutoMigrate(&model.User{})

//     service := service.NewUserService(db)

//     user := &model.User{Name: "John Doe", Email: "john@example.com"}

//     // Call the method
//     err = service.CreateUser(user)

//     // Assert that no error occurred
//     assert.NoError(t, err)
// }

// // TestUserService_GetUsers tests the GetUsers method
// func TestUserService_GetUsers(t *testing.T) {
//     db, err := setupDatabase()
//     if err != nil {
//         t.Fatalf("failed to connect to database: %v", err)
//     }

//     // Migrate the schema
//     db.AutoMigrate(&model.User{})

//     service := service.NewUserService(db)

//     users := []model.User{
//         {Name: "John Doe2", Email: "john@example.com"},
//         {Name: "Jane Smith2", Email: "jane@example.com"},
//     }

//     // Insert users into the database
//     db.Create(&users)

//     // Call the method
//     result, err := service.GetUsers()

//     // Assert that no error occurred
//     assert.NoError(t, err)

//     // Assert that the returned users match
//     assert.ElementsMatch(t, users, result)
// }

// // TestUserService_GetUserByID tests the GetUserByID method
// func TestUserService_GetUserByID(t *testing.T) {
//     db, err := setupDatabase()
//     if err != nil {
//         t.Fatalf("failed to connect to database: %v", err)
//     }

//     // Migrate the schema
//     db.AutoMigrate(&model.User{})

//     service := service.NewUserService(db)

//     user := &model.User{Name: "John Doe", Email: "john@example.com"}

//     // Insert user into the database
//     db.Create(user)


	


//     // Call the method

// 	id := strconv.FormatUint(uint64(user.ID), 10)
//     result, err := service.GetUserByID(id)

//     // Assert that no error occurred
//     assert.NoError(t, err)

//     // Assert that the returned user matches
//     assert.Equal(t, user, result)
// }

// // TestUserService_UpdateUser tests the UpdateUser method
// func TestUserService_UpdateUser(t *testing.T) {
//     db, err := setupDatabase()
//     if err != nil {
//         t.Fatalf("failed to connect to database: %v", err)
//     }

//     // Migrate the schema
//     db.AutoMigrate(&model.User{})

//     service := service.NewUserService(db)

//     user := &model.User{Name: "John Doe", Email: "john@example.com"}

//     // Insert user into the database
//     db.Create(user)

//     // Update user information
//     user.Email = "john.doe@example.com"

//     // Call the method
//     err = service.UpdateUser(user)

//     // Assert that no error occurred
//     assert.NoError(t, err)

//     // Retrieve the updated user

// 	id := strconv.FormatUint(uint64(user.ID), 10)

//     updatedUser, err := service.GetUserByID(id)
//     assert.NoError(t, err)
//     assert.Equal(t, "john.doe@example.com", updatedUser.Email)
// }
