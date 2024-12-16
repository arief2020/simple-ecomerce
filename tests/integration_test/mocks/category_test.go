package mocks

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"github.com/stretchr/testify/mock"
)

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
