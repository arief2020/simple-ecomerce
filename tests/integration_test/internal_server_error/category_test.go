package tests

import (
	"context"
	"errors"
	// "net/http"
	// "net/http/httptest"
	"testing"

	// "tugas_akhir_example/internal/pkg/controller"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/usecase"

	// "github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock Repository
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetAllCategory(ctx context.Context) ([]entity.Category, error) {
	args := m.Called(ctx)
	return args.Get(0).([]entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetCategoryByID(ctx context.Context, id uint) (entity.Category, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) CreateCategory(ctx context.Context, data entity.Category) (entity.Category, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) UpdateCategoryByID(ctx context.Context, id uint, data entity.Category) (entity.Category, error) {
	args := m.Called(ctx, id, data)
	return args.Get(0).(entity.Category), args.Error(1)
}

func (m *MockCategoryRepository) DeleteCategoryByID(ctx context.Context, id uint) (string, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(string), args.Error(1)
}

func TestGetAllCategory_InternalServerError(t *testing.T) {
	mockRepo := new(MockCategoryRepository)

	// Atur mock behavior
	mockRepo.On("GetAllCategory", mock.Anything).
		Return([]entity.Category{}, errors.New("internal server error"))

	// Inisialisasi usecase dengan mock
	categoryUseCase := usecase.NewCategoryUseCase(mockRepo)

	// Eksekusi fungsi yang diuji
	ctx := context.Background()
	categories, err := categoryUseCase.GetAllCategory(ctx)

	// Validasi hasil
	assert.Nil(t, categories) // Karena error, categories harus nil
	assert.NotNil(t, err)
	assert.Equal(t, "internal server error", err.Err.Error())

	mockRepo.AssertExpectations(t)
}
