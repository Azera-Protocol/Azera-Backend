package tests

import (
	"azera-backend/controllers"
	"azera-backend/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockDB is a mock for the GORM DB
type MockDB struct {
	mock.Mock
}

// Create mocks the Create method of GORM
func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

var mockDB *MockDB
var router *gin.Engine

func init() {
	// Initialize mock DB and router
	mockDB = new(MockDB)
	router = setupRouter()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Assuming you have a way to inject the mock DB into your controller
	r.POST("/users/create", func(c *gin.Context) {
		controllers.CreateUserDetails(c)
	})
	return r
}

func TestCreateUserDetailsSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	user := models.User{
		// Fill with appropriate user data
	}
	mockDB.On("Create", &user).Return(&gorm.DB{}) // Mocking the DB call

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/users/create", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	// Additional assertions to check the response body if necessary
	mockDB.AssertExpectations(t)
}
