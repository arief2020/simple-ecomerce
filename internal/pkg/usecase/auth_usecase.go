package usecase

import (
	"context"
	"errors"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	userdto "tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/repository"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthsUseCase interface {
	Login(ctx context.Context, params userdto.Login) (res userdto.LoginRes, err *helper.ErrorStruct)
	CreateUsers(ctx context.Context, data userdto.CreateUser) (res uint, err *helper.ErrorStruct)
}

type AuthUseCaseImpl struct {
	userrepository repository.UsersRepository
	provinceCityRepository repository.ProvinceCityRepository
	tokoRepository repository.TokoRepository
}

func NewAuthUseCase(userrepository repository.UsersRepository, provinceCityRepository repository.ProvinceCityRepository, tokoRepository repository.TokoRepository) AuthsUseCase {
	return &AuthUseCaseImpl{
		userrepository: userrepository,
		provinceCityRepository: provinceCityRepository,
		tokoRepository: tokoRepository,
	}

}

func (alc *AuthUseCaseImpl) Login(ctx context.Context, params userdto.Login) (res userdto.LoginRes, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.userrepository.GetUserByNoTelp(ctx, params.NoTelp)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found: Get User By NoTelp")
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("no telp atau kata sandi salah"),
		}
	}

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Get User By NoTelp : %s", errRepo.Error()))
		fmt.Println("debug 2")
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}


	isValid := utils.CheckPasswordHash(params.KataSandi, resRepo.KataSandi)
	if !isValid {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Invalid Password")
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("no telp atau kata sandi salah"),
		}
	}

	dataProvince, errRepo := alc.provinceCityRepository.GetProvinceByID(ctx, resRepo.IdProvinsi)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Get Province By ID : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	dataCity, errRepo := alc.provinceCityRepository.GetCityByID(ctx, resRepo.IdKota)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Get City By ID : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	tokenInit := utils.NewToken(utils.DataClaims{
		ID:    fmt.Sprint(resRepo.ID),
		Email: resRepo.Email,
		IsAdmin: resRepo.IsAdmin,
	})

	token, errToken := tokenInit.Create()
	if errToken != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Create Token : %s", errToken.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errToken,
		}
	}
	tanggalLahirFormatted := resRepo.TanggalLahir.Format("02/01/2006")

	res = userdto.LoginRes{
		Email: resRepo.Email,
		Nama:  resRepo.Nama,
		NoTelp: resRepo.NoTelp,
		TanggalLahir: tanggalLahirFormatted,
		Tentang: resRepo.Tentang,
		Pekerjaan: resRepo.Pekerjaan,
		IdProvinsi: dataProvince,
		IdKota: dataCity,
		Token: token,
	}

	return res, nil
}
func (alc *AuthUseCaseImpl) CreateUsers(ctx context.Context, params userdto.CreateUser) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(params); errValidate != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Validate : %s", errValidate.Error()))
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resCityByProvRepo, _ := alc.provinceCityRepository.GetAllCitiesByProvinceID(ctx, params.IdProvinsi)
	if resCityByProvRepo == nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found: Get City By Province ID")
		return res, &helper.ErrorStruct{
			Err:  errors.New("data kota tidak ditemukan"),
			Code: fiber.StatusNotFound,
		}
	}

	isCityExist := utils.IsIDExist(resCityByProvRepo, params.IdKota)
	if !isCityExist {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found: Get City By Province ID")
		return res, &helper.ErrorStruct{
			Err:  errors.New("data kota tidak ditemukan"),
			Code: fiber.StatusNotFound,
		}
	}

	hashPass, errHash := utils.HashPassword(params.KataSandi)
	if errHash != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Hash Password : %s", errHash.Error()))
		
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errHash,
		}
	}

	TanggalLahirParse, errParse := utils.ParseDate(params.TanggalLahir)
	if errParse != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at Parse Date : %s", errParse.Error()))
		
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errParse,
		}
	}


	resRepo, errRepo := alc.userrepository.CreateUsers(ctx, entity.User{
		Email:    params.Email,
		Nama:     params.Name,
		KataSandi: hashPass,
		TanggalLahir: TanggalLahirParse,
		NoTelp:       params.NoTelp,
		JenisKelamin: params.JenisKelamin,
		Pekerjaan:   params.Pekerjaan,
		IdProvinsi:  params.IdProvinsi,
		IdKota:      params.IdKota,
	})
	
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at CreateUsers : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	_, errRepoToko := alc.tokoRepository.CreateToko(ctx, entity.Toko{
		NamaToko: nil,
		UrlFoto:  nil,
		UserID:   resRepo.ID,
	})
	if errRepoToko != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at CreateToko : %s", errRepo))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo.ID, nil
}
