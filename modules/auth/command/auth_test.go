package command_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gin-money-manager-api/modules/auth/command"
	"gin-money-manager-api/modules/auth/dto"
	"gin-money-manager-api/modules/user/entity"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserFinder struct {
	mock.Mock
}

func (m *MockUserFinder) FindByUsername(username string) (*entity.User, error) {
	args := m.Called(username)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entity.User), args.Error(1)
}

func TestAuthCommand_Handler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockUserFinder)

	hashed, _ := bcrypt.GenerateFromPassword([]byte("password123"), 10)

	user := &entity.User{
		ID:       uuid.New(),
		Username: "admin",
		Password: string(hashed),
	}

	mockRepo.
		On("FindByUsername", "admin").
		Return(user, nil)
	cmd := command.NewAuthCommand(mockRepo)

	body := dto.LoginRequest{
		Username: "admin",
		Password: "password123",
	}

	jsonBody, _ := json.Marshal(body)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/auth",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	cmd.Handler(c)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}
