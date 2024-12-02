package usecase

import (
	"context"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"
)

type CategoryUseCase interface {
	GetAllCategory(ctx context.Context) ([]*dto.CategoryResp, *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, id uint) (*dto.CategoryResp, *helper.ErrorStruct)
	CreateCategory(ctx context.Context, data dto.CategoryReq) (*dto.CategoryResp, *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, id uint, data dto.CategoryReq) (*dto.CategoryResp, *helper.ErrorStruct)
	DeleteCategoryByID(ctx context.Context, id uint) (string, *helper.ErrorStruct)
}

type CategoryUseCaseImpl struct {
	categoryRepository repository.CategoryRepository
}


func NewCategoryUseCase(categoryRepository repository.CategoryRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		categoryRepository: categoryRepository,
	}
}

func (c *CategoryUseCaseImpl) GetAllCategory(ctx context.Context) ([]*dto.CategoryResp, *helper.ErrorStruct) {
	categories, err := c.categoryRepository.GetAllCategory(ctx)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	var categoriesResp []*dto.CategoryResp
	for _, category := range categories {
		categoriesResp = append(categoriesResp, &dto.CategoryResp{
			ID:          category.ID,
			NamaCategory: category.NamaCategory,
		})
	}

	return categoriesResp, nil
}

func (c *CategoryUseCaseImpl) GetCategoryByID(ctx context.Context, id uint) (*dto.CategoryResp, *helper.ErrorStruct) {
	category, err := c.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	return &dto.CategoryResp{
		ID:          category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (c *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data dto.CategoryReq) (*dto.CategoryResp, *helper.ErrorStruct) {
	dataReq := entity.Category{
		NamaCategory: data.NamaCategory,	
	}
	category, err := c.categoryRepository.CreateCategory(ctx, dataReq)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	return &dto.CategoryResp{
		ID:          category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (c *CategoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, id uint, data dto.CategoryReq) (*dto.CategoryResp, *helper.ErrorStruct) {
	dataReq := entity.Category{
		NamaCategory: data.NamaCategory,	
	}
	category, err := c.categoryRepository.UpdateCategoryByID(ctx, id, dataReq)
	if err != nil {
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	return &dto.CategoryResp{
		ID:          category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (c *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, id uint) (string, *helper.ErrorStruct) {
	_, err := c.categoryRepository.DeleteCategoryByID(ctx, id)
	if err != nil {
		return "", &helper.ErrorStruct{
			Err: err,
		}
	}

	return "Succeed to DELETE data", nil
}