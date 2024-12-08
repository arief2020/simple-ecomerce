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

type UserUseCase interface {
	GetMyProfile(ctx context.Context, id uint) (*dto.UserResp, *helper.ErrorStruct)
	UpdateMyProfile(ctx context.Context, id uint, params dto.UpdateUser) (string, *helper.ErrorStruct)

	GetMyAlamat(ctx context.Context, id uint, params dto.FiltersAlamat) ([]*dto.AlamatResp, *helper.ErrorStruct)

	CreateMyNewAlamat(ctx context.Context, id uint, params dto.InserAlamatReq) (int, *helper.ErrorStruct)

	GetMyAlamatById(ctx context.Context, id uint, idAlamat uint) (*dto.AlamatResp, *helper.ErrorStruct)
	UpdateMyAlamatById(ctx context.Context, id uint, idAlamat uint, params dto.UpdateAlamatReq) (string, *helper.ErrorStruct)
	DeleteMyAlamatById(ctx context.Context, id uint, idAlamat uint) (string, *helper.ErrorStruct)
}

type UserUseCaseImpl struct {
	userrepository         repository.UsersRepository
	provinceCityRepository repository.ProvinceCityRepository
}

func NewUserUseCase(userrepository repository.UsersRepository, provinceCityRepository repository.ProvinceCityRepository) UserUseCase {
	return &UserUseCaseImpl{
		userrepository:         userrepository,
		provinceCityRepository: provinceCityRepository,
	}

}

func (uc *UserUseCaseImpl) GetMyProfile(ctx context.Context, id uint) (*dto.UserResp, *helper.ErrorStruct) {
	// Panggil repository untuk mendapatkan data user
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyProfile")
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	dataProvince, errRepo := uc.provinceCityRepository.GetProvinceByID(ctx, user.IdProvinsi)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyProfile")
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	dataCity, errRepo := uc.provinceCityRepository.GetCityByID(ctx, user.IdKota)

	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyProfile")
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	birthDate := utils.FormatDate(user.TanggalLahir)

	// Map entity.User ke dto.UserResp
	resp := &dto.UserResp{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		Email:        user.Email,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		TanggalLahir: birthDate,
		IdKota:       dataCity,
		IdProvinsi:   dataProvince,
		IsAdmin:      user.IsAdmin,
		// Lakukan mapping lainnya jika diperlukan
	}
	return resp, nil
}

func (uc *UserUseCaseImpl) UpdateMyProfile(ctx context.Context, id uint, params dto.UpdateUser) (string, *helper.ErrorStruct) {
	// Panggil repository untuk mendapatkan data user
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		fmt.Println("debug updatemyprofile 1")
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at UpdateMyProfile")
		return "", &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	TanggalLahirParsed, err := utils.ParseDate(params.TanggalLahir)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at UpdateMyProfile")
		return "", &helper.ErrorStruct{
			Err:  err,
			Code: fiber.StatusBadRequest,
		}
	}

	user.Nama = params.Nama
	user.NoTelp = params.NoTelp
	user.TanggalLahir = TanggalLahirParsed
	user.Tentang = params.Tentang
	user.Pekerjaan = params.Pekerjaan
	user.Email = params.Email
	user.IdProvinsi = params.IdProvinsi
	user.IdKota = params.IdKota

	resUpdate, errRepo := uc.userrepository.UpdateUserById(ctx, id, user)
	if errRepo != nil {
		fmt.Println("debug updatemyprofile 2")
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at UpdateMyProfile")
		return "", &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return resUpdate, nil
}

func (uc *UserUseCaseImpl) GetMyAlamat(ctx context.Context, id uint, params dto.FiltersAlamat) ([]*dto.AlamatResp, *helper.ErrorStruct) {
	// Panggil repository untuk mendapatkan data user
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyAlamat")
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	alamat, errRepo := uc.userrepository.GetAlamatByUserId(ctx, user.ID, params)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyAlamat")
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	resp := make([]*dto.AlamatResp, len(alamat))
	for i, a := range alamat {
		resp[i] = &dto.AlamatResp{
			Id:           a.ID,
			JudulAlamat:  a.JudulAlamat,
			NamaPenerima: a.NamaPenerima,
			NoTelp:       a.NoTelp,
			DetailAlamat: a.DetailAlamat,
		}
	}

	return resp, nil
}

func (uc *UserUseCaseImpl) CreateMyNewAlamat(ctx context.Context, id uint, params dto.InserAlamatReq) (int, *helper.ErrorStruct) {
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get User By Id")
		return 0, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	dataEntity := entity.Alamat{
		JudulAlamat:  params.JudulAlamat,
		NamaPenerima: params.NamaPenerima,
		NoTelp:       params.NoTelp,
		DetailAlamat: params.DetailAlamat,
		UserID:       user.ID,
	}

	alamat, errRepo := uc.userrepository.InsertAlamat(ctx, dataEntity)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at Create New Alamat")
		return 0, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return int(alamat.ID), nil
}

func (uc *UserUseCaseImpl) GetMyAlamatById(ctx context.Context, id uint, idAlamat uint) (*dto.AlamatResp, *helper.ErrorStruct) {
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyAlamatByID")
		return nil, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	// Panggil repository untuk mendapatkan data alamat
	alamat, errRepo := uc.userrepository.GetMyAlamatById(ctx, user.ID, idAlamat)
	if errRepo != nil {
		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Alamat If Not Found")
			return nil, &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("alamat not found"),
			}
		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at GetMyAlamatByID")
		return nil, &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	resp := &dto.AlamatResp{
		Id:           alamat.ID,
		JudulAlamat:  alamat.JudulAlamat,
		NamaPenerima: alamat.NamaPenerima,
		NoTelp:       alamat.NoTelp,
		DetailAlamat: alamat.DetailAlamat,
	}

	return resp, nil
}

func (uc *UserUseCaseImpl) UpdateMyAlamatById(ctx context.Context, id uint, idAlamat uint, params dto.UpdateAlamatReq) (string, *helper.ErrorStruct) {
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get User By Id")
		return "", &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	alamat, errRepo := uc.userrepository.GetMyAlamatById(ctx, user.ID, idAlamat)
	if errRepo != nil {

		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Alamat")
			return "", &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("record not found"),
			}

		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Alamat By Id")
		return "", &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	dataEntity := entity.Alamat{
		JudulAlamat:  alamat.JudulAlamat,
		NamaPenerima: params.NamaPenerima,
		NoTelp:       params.NoTelp,
		DetailAlamat: params.DetailAlamat,
		UserID:       user.ID,
	}

	_, errRepo = uc.userrepository.UpdateMyAlamatById(ctx, user.ID, idAlamat, dataEntity)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at UpdateMyAlamatByID")
		return "", &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return "Succeed to Update data", nil
}

func (uc *UserUseCaseImpl) DeleteMyAlamatById(ctx context.Context, id uint, idAlamat uint) (string, *helper.ErrorStruct) {
	user, err := uc.userrepository.GetUserById(ctx, id)
	if err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at DeleteMyAlamatByID")
		return "", &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("user not found"),
		}
	}

	// Panggil repository untuk mendapatkan data alamat
	_, errRepo := uc.userrepository.GetMyAlamatById(ctx, user.ID, idAlamat)
	if errRepo != nil {
		if errors.Is(errRepo, gorm.ErrRecordNotFound) {
			helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Not Found Alamat")
			return "", &helper.ErrorStruct{
				Code: fiber.StatusNotFound,
				Err:  errors.New("record not found"),
			}

		}
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at DeleteMyAlamatByID")
		return "", &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	// Hapus data alamat dari database
	_, errRepo = uc.userrepository.DeleteMyAlamatById(ctx, user.ID, idAlamat)
	if errRepo != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error at DeleteMyAlamatByID")
		return "", &helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusBadRequest,
		}
	}

	return "Succeed to Delete data", nil
}
