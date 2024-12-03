package usecase

import (
	"context"
	"errors"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("debug 1")
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no data category"),
	}
}
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	fmt.Printf("Category: %+v\n", category)
	
	fmt.Printf("Error: %+v\n", err)
	

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
	res, err := c.categoryRepository.DeleteCategoryByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("record not found"),
			}
		}
		return "", &helper.ErrorStruct{
			Err: err,
		}
	}

	fmt.Printf("res: %+v\n", res)

	return "Succeed to DELETE data", nil
}