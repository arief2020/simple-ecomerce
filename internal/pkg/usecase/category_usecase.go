package usecase

import (
	"context"
	"errors"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

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
			Err:  err,
			Code: fiber.StatusInternalServerError,
		}
	}

	var categoriesResp []*dto.CategoryResp
	for _, category := range categories {
		categoriesResp = append(categoriesResp, &dto.CategoryResp{
			ID:           category.ID,
			NamaCategory: category.NamaCategory,
		})
	}

	return categoriesResp, nil
}

func (c *CategoryUseCaseImpl) GetCategoryByID(ctx context.Context, id uint) (*dto.CategoryResp, *helper.ErrorStruct) {
	category, err := c.categoryRepository.GetCategoryByID(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Get Category")
			return nil, &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("no data category"),
			}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Get Category")
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	return &dto.CategoryResp{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (c *CategoryUseCaseImpl) CreateCategory(ctx context.Context, data dto.CategoryReq) (*dto.CategoryResp, *helper.ErrorStruct) {
	dataReq := entity.Category{
		NamaCategory: data.NamaCategory,
	}
	category, err := c.categoryRepository.CreateCategory(ctx, dataReq)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create Category")
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	return &dto.CategoryResp{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (c *CategoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, id uint, data dto.CategoryReq) (*dto.CategoryResp, *helper.ErrorStruct) {
	dataReq := entity.Category{
		NamaCategory: data.NamaCategory,
	}
	_, errGetCategory := c.categoryRepository.GetCategoryByID(ctx, id)
	if errGetCategory != nil {
		if errors.Is(errGetCategory, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Get Category")
			return nil, &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("no data category"),
			}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Category By Id")
		return nil, &helper.ErrorStruct{
			Err: errGetCategory,
			Code: fiber.StatusBadRequest,
		}
	}
	category, err := c.categoryRepository.UpdateCategoryByID(ctx, id, dataReq)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update Category")
		return nil, &helper.ErrorStruct{
			Err: err,
		}
	}

	return &dto.CategoryResp{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}, nil
}

func (c *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, id uint) (string, *helper.ErrorStruct) {
	res, err := c.categoryRepository.DeleteCategoryByID(ctx, id)
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Delete Category")
			return "", &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("record not found"),
			}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete Category")
		return "", &helper.ErrorStruct{
			Err: err,
		}
	}

	fmt.Printf("res: %+v\n", res)

	return "Succeed to DELETE data", nil
}
