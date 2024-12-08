package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/utils"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUsersByEmail(ctx context.Context, email string) (res entity.User, err error)
	CreateUsers(ctx context.Context, data entity.User) (res entity.User, err error)
	GetUserByNoTelp(ctx context.Context, no_telp string) (res entity.User, err error)
	GetUserById(ctx context.Context, id uint) (res entity.User, err error)
	UpdateUserById(ctx context.Context, id uint, data entity.User) (res string, err error)

	GetAlamatByUserId(ctx context.Context, id uint, params dto.FiltersAlamat) (res []entity.Alamat, err error)
	InsertAlamat(ctx context.Context, data entity.Alamat) (res entity.Alamat, err error)
	GetMyAlamatById(ctx context.Context, id uint, idAlamat uint) (res entity.Alamat, err error)
	UpdateMyAlamatById(ctx context.Context, id uint, idAlamat uint, data entity.Alamat) (res string, err error)
	DeleteMyAlamatById(ctx context.Context, id uint, idAlamat uint) (res string, err error)
}

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{
		db: db,
	}
}

func (r *UsersRepositoryImpl) GetUsersByEmail(ctx context.Context, email string) (res entity.User, err error) {
	if err := r.db.Where("email = ?", email).First(&res).WithContext(ctx).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get User By Email")
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) GetUserByNoTelp(ctx context.Context, no_telp string) (res entity.User, err error) {
	if err := r.db.Where("no_telp = ?", no_telp).First(&res).WithContext(ctx).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get User By No Telp")
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) CreateUsers(ctx context.Context, data entity.User) (res entity.User, err error) {
	result := r.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Create User")
		return res, result.Error
	}

	return data, nil
}

func (r *UsersRepositoryImpl) GetUserById(ctx context.Context, id uint) (res entity.User, err error) {

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get User By Id")
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) UpdateUserById(ctx context.Context, id uint, data entity.User) (res string, err error) {
	if err := r.db.Model(&entity.User{}).WithContext(ctx).Where("id = ?", id).Updates(data).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update User By Id")
		return res, err
	}

	return "Succeed to Update data", nil
}

func (r *UsersRepositoryImpl) GetAlamatByUserId(ctx context.Context, id uint, params dto.FiltersAlamat) (res []entity.Alamat, err error) {
	db := r.db

	filter := map[string][]any{}
	if params.JudulAlamat != "" {
		filter["judul_alamat LIKE ?"] = []any{fmt.Sprintf("%%%s%%", params.JudulAlamat)}
	}

	for key, val := range filter {
		db = db.Where(key, val...)
	}

	db = db.Where("id_user = ?", id)

	if err := db.Debug().WithContext(ctx).Find(&res).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get Alamat By User Id")
		return res, err
	}

	return res, nil
}

func (r *UsersRepositoryImpl) InsertAlamat(ctx context.Context, data entity.Alamat) (res entity.Alamat, err error) {
	if err := r.db.Create(&data).WithContext(ctx).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Insert Alamat")
		return res, err
	}
	return data, nil
}

func (r *UsersRepositoryImpl) GetMyAlamatById(ctx context.Context, id uint, idAlamat uint) (res entity.Alamat, err error) {
	if err := r.db.WithContext(ctx).Where("id_user = ? AND id = ?", id, idAlamat).First(&res).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Get My Alamat By Id")
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) UpdateMyAlamatById(ctx context.Context, id uint, idAlamat uint, data entity.Alamat) (res string, err error) {
	if err := r.db.Model(&entity.Alamat{}).WithContext(ctx).Where("id_user = ? AND id = ?", id, idAlamat).Updates(data).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Update My Alamat By Id")
		return res, err
	}
	return "success", nil
}

func (r *UsersRepositoryImpl) DeleteMyAlamatById(ctx context.Context, id uint, idAlamat uint) (res string, err error) {
	if err := r.db.WithContext(ctx).Where("id_user = ? AND id = ?", id, idAlamat).Delete(&entity.Alamat{}).Error; err != nil {
		helper.Logger(utils.GetFunctionPath(), helper.LoggerLevelError, "Error Delete My Alamat By Id")
		return res, err
	}
	return "success", nil
}
