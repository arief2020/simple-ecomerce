package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
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
		fmt.Println("debug 1")
		fmt.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Telp atau kata sandi salah"),
		}
	}

	fmt.Println("debug 1 success")

	if errRepo != nil {
		// helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at GetAllUsers : %s", errRepo.Error()), errRepo)
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at GetAllUsers : %s", errRepo.Error()))
		fmt.Println("debug 2")
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	fmt.Println("debug 2 success")

	isValid := utils.CheckPasswordHash(params.KataSandi, resRepo.KataSandi)
	if !isValid {
		fmt.Println("debug 3")
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("No Telp atau kata sandi salah"),
		}
	}

	dataProvince, errRepo := alc.provinceCityRepository.GetProvinceByID(ctx, resRepo.IdProvinsi)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	dataCity, errRepo := alc.provinceCityRepository.GetCityByID(ctx, resRepo.IdKota)
	if errRepo != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	fmt.Println("debug 3 success")

	tokenInit := utils.NewToken(utils.DataClaims{
		ID:    fmt.Sprint(resRepo.ID),
		Email: resRepo.Email,
		IsAdmin: resRepo.IsAdmin,
	})

	token, errToken := tokenInit.Create()
	if errToken != nil {
		fmt.Println("debug 4")
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errToken,
		}
	}

	// layout := "02/01/2006" // Define the layout you expect the date to be in
	// tanggalLahirParsed, errParse := time.Parse(resRepo.TanggalLahir, layout)
	// if errParse != nil {
	// 	log.Println("Error parsing date:", errParse)
	// 	err = &helper.ErrorStruct{
	// 		Code: fiber.StatusBadRequest,
	// 		Err:  errParse,
	// 	}
	// 	return
	// }

	fmt.Println("debug 4 success")

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
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resCityByProvRepo, _ := alc.provinceCityRepository.GetAllCitiesByProvinceID(ctx, params.IdProvinsi)
	if resCityByProvRepo == nil {
		return res, &helper.ErrorStruct{
			Err:  errors.New("data kota tidak ditemukan"),
			Code: fiber.StatusNotFound,
		}
	}

	isCityExist := utils.IsIDExist(resCityByProvRepo, params.IdKota)
	if !isCityExist {
		return res, &helper.ErrorStruct{
			Err:  errors.New("data kota tidak ditemukan"),
			Code: fiber.StatusNotFound,
		}
	}

	// resProvRepo, _ := alc.provinceCityRepository.GetProvinceByID(ctx, params.IdProvinsi)
	// if resProvRepo == nil {
	// 	return res, &helper.ErrorStruct{
	// 		Err:  errors.New("Data Provinsi Tidak Ditemukan"),
	// 		Code: fiber.StatusNotFound,
	// 	}
	// }

	// resKotaRepo, _ := alc.provinceCityRepository.GetCityByID(ctx, params.IdKota)
	// if resKotaRepo == nil {
	// 	return res, &helper.ErrorStruct{
	// 		Err:  errors.New("Data Kota Tidak Ditemukan"),
	// 		Code: fiber.StatusNotFound,
	// 	}
	// }
	hashPass, errHash := utils.HashPassword(params.KataSandi)
	if errHash != nil {
		log.Println(errHash)
		err = &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errHash,
		}
		return
	}

	TanggalLahirParse, errParse := utils.ParseDate(params.TanggalLahir)
	if err != nil {
		log.Println(err)
		err = &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errParse,
		}
		return
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
		// helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateUsers : %s", errRepo.Error()), errRepo)
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
		// helper.Logger(helper.LoggerLevelError, fmt.Sprintf("Error at CreateToko : %s", errRepo.Error()), errRepo)
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, fmt.Sprintf("Error at CreateToko : %s", errRepo.Error()))
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}



	return resRepo.ID, nil
}
