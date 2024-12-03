package repository

import (
	"context"
	"fmt"
	"tugas_akhir_example/internal/pkg/dto"
	"tugas_akhir_example/internal/pkg/entity"

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
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) GetUserByNoTelp(ctx context.Context, no_telp string) (res entity.User, err error) {
	// if err := r.db.Where("no_telp = ?", no_telp).First(&res).WithContext(ctx).Error; err != nil {
	// 	return res, err
	// }

	if err := r.db.Raw("SELECT * FROM users WHERE no_telp = ? LIMIT 1", no_telp).Scan(&res).Error; err != nil {
    	return res, err
	}
	return res, nil
}

// func (r *UsersRepositoryImpl) CreateUsers(ctx context.Context, data entity.User) (res entity.User, err error) {
// 	result := r.db.Create(&data).WithContext(ctx)
// 	if result.Error != nil {
// 		return res, result.Error
// 	}

// 	return result, nil
// }

func (r *UsersRepositoryImpl) CreateUsers(ctx context.Context, data entity.User) (res entity.User, err error) {
    result := r.db.Create(&data).WithContext(ctx)
    if result.Error != nil {
        return res, result.Error
    }

    return data, nil // Mengembalikan objek User
}

// func (r *UsersRepositoryImpl) GetUserById(ctx context.Context, id uint) (res entity.User, err error) {
// 	if err := r.db.Where("id = ?", id).First(&res).WithContext(ctx).Error; err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

// func (r *UsersRepositoryImpl) GetUserById(ctx context.Context, id uint) (entity.User, error) {
// 	var user entity.User
// 	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&user).Error; err != nil {
// 		return user, err
// 	}
// 	return user, nil
// }


// func (r *UsersRepositoryImpl) GetUserById(ctx context.Context, id uint) (res entity.User, err error) {
// 	// Validasi id sebelum query
// 	if id == 0 {
// 		return res, errors.New("invalid user ID")
// 	}

// 	// Gunakan query sederhana tanpa tambahan kondisi yang tidak perlu
// 	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&res).Error; err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

func (r *UsersRepositoryImpl) GetUserById(ctx context.Context, id uint) (res entity.User, err error) {
	// Gunakan query manual untuk mengambil data user berdasarkan ID
	query := "SELECT * FROM users WHERE id = ? LIMIT 1"
	if err := r.db.Raw(query, id).Scan(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) UpdateUserById(ctx context.Context, id uint, data entity.User) (res string, err error) {
	sql := `UPDATE users SET nama = ?, email = ?, no_telp = ?, tanggal_lahir = ?, tentang = ?, pekerjaan = ?, id_provinsi = ?, id_kota = ? WHERE id = ?`

	if err := r.db.Exec(sql, data.Nama, data.Email, data.NoTelp, data.TanggalLahir, data.Tentang, data.Pekerjaan, data.IdProvinsi, data.IdKota, id).Error; err != nil {
		return res, err
	}

	return "Succeed to Update data", nil
}

// func (r *UsersRepositoryImpl) GetAlamatByUserId(ctx context.Context, id uint, params dto.FiltersAlamat) (res []entity.Alamat, err error) {
// 	db := r.db

// 	filter := map[string][]any{
// 		"judul_alamat like ?": {fmt.Sprint("%" + params.JudulAlamat+"%")},
// 	}

// 	for key, val := range filter {
// 		db = db.Where(key, val...)
// 	}

// 	if err := db.Debug().WithContext(ctx).Where("id_user = ?", id).Find(&res).Error; err != nil {
// 		return res, err
		
// 	}
// 	// query := "SELECT * FROM alamats WHERE id_user = ?"
// 	// if err := r.db.Raw(query, id).Scan(&res).Error; err != nil {
// 	// 	return res, err
// 	// }
// 	return res, nil
// }

func (r *UsersRepositoryImpl) GetAlamatByUserId(ctx context.Context, id uint, params dto.FiltersAlamat) (res []entity.Alamat, err error) {
    db := r.db

    // Validasi dan buat filter
    filter := map[string][]any{}
    if params.JudulAlamat != "" {
        filter["judul_alamat LIKE ?"] = []any{fmt.Sprintf("%%%s%%", params.JudulAlamat)}
    }

    // Terapkan filter
    for key, val := range filter {
        db = db.Where(key, val...)
    }

    // Tambahkan filter `id_user`
    db = db.Where("id_user = ?", id)

    // Eksekusi query
    if err := db.Debug().WithContext(ctx).Find(&res).Error; err != nil {
        return res, err
    }

    return res, nil
}

func (r *UsersRepositoryImpl) InsertAlamat(ctx context.Context, data entity.Alamat) (res entity.Alamat, err error) {
	if err := r.db.Create(&data).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return data, nil
}

func (r *UsersRepositoryImpl) GetMyAlamatById(ctx context.Context, id uint, idAlamat uint) (res entity.Alamat, err error) {
	// query := "SELECT * FROM alamats WHERE id_user = ? AND id = ?"
	// result := r.db.Raw(query, id, idAlamat).Scan(&res)
	// if result.Error != nil {
	// 	return res, result.Error // Error teknis
	// }
	// if result.RowsAffected == 0 {
	// 	return res, errors.New("alamat not found") // Error jika data kosong
	// }

	if err := r.db.WithContext(ctx).Where("id_user = ? AND id = ?", id, idAlamat).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) UpdateMyAlamatById(ctx context.Context, id uint, idAlamat uint, data entity.Alamat) (res string, err error) {
	query := "UPDATE alamats SET judul_alamat = ?, nama_penerima = ?, no_telp = ?, detail_alamat = ? WHERE id_user = ? AND id = ?"
	if err := r.db.Exec(query, data.JudulAlamat, data.NamaPenerima, data.NoTelp, data.DetailAlamat, id, idAlamat).Error; err != nil {
		return res, err
	}
	return "success", nil
}

func (r *UsersRepositoryImpl) DeleteMyAlamatById(ctx context.Context, id uint, idAlamat uint) (res string, err error) {
	query := "DELETE FROM alamats WHERE id_user = ? AND id = ?"
	if err := r.db.Exec(query, id, idAlamat).Error; err != nil {
		return res, err
	}
	return "success", nil
}