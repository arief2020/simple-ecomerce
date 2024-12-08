package repository

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"tugas_akhir_example/internal/pkg/dto"
	// "gorm.io/gorm"
)

const provinceCityListProvinceAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/provinces.json"
const provinceCityListCityByProvinceAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/regencies/%s.json"
const provinceCityDetailProvinceAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/province/%s.json"
const provinceCityDetailCityAPI = "https://emsifa.github.io/api-wilayah-indonesia/api/regency/%s.json"


type ProvinceCityRepository interface {
	// ListProvincies(ctx context.Context) (res dto.ListProvResp, err error)
	GetAllProvinces(ctx context.Context, limit, offset int, search string) (res []*dto.ProvinceResp, err error)
	GetAllCitiesByProvinceID(ctx context.Context, provinceid string) (res []*dto.CityResp, err error)
	GetProvinceByID(ctx context.Context, provinceid string) (res *dto.ProvinceResp, err error)
	GetCityByID(ctx context.Context, cityid string) (res *dto.CityResp, err error)
}

// type provCityUseCaseImpl struct {
// 	db *gorm.DB
// }

type ProvinceCityRepositoryImpl struct {
}

func NewProvinceCityRepository() ProvinceCityRepository {
	return &ProvinceCityRepositoryImpl{}
}

func (alr *ProvinceCityRepositoryImpl) GetAllProvinces(ctx context.Context, limit, offset int, search string) (res []*dto.ProvinceResp, err error) {
	resp, err := http.Get(provinceCityListProvinceAPI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	println(len(res))
	search = strings.ToLower(search)
	tmpres := []*dto.ProvinceResp{}
	for _, v := range res {
		if strings.Contains(strings.ToLower(v.Name), search) {
			tmpres = append(tmpres, v)
		}
	}
	res = tmpres
	println(len(res))

	if offset >= len(res) {
		res = nil
	} else {
		endIndex := offset + limit
		if endIndex > len(res) {
			endIndex = len(res)
		}
		res = res[offset:endIndex]
	}

	return res, nil
}

func (alr *ProvinceCityRepositoryImpl) GetAllCitiesByProvinceID(ctx context.Context, provinceid string) (res []*dto.CityResp, err error) {
	resp, err := http.Get(strings.Replace(provinceCityListCityByProvinceAPI, "%s", provinceid, 1))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (alr *ProvinceCityRepositoryImpl) GetProvinceByID(ctx context.Context, provinceid string) (res *dto.ProvinceResp, err error) {
	resp, err := http.Get(strings.Replace(provinceCityDetailProvinceAPI, "%s", provinceid, 1))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (alr *ProvinceCityRepositoryImpl) GetCityByID(ctx context.Context, cityid string) (res *dto.CityResp, err error) {
	resp, err := http.Get(strings.Replace(provinceCityDetailCityAPI, "%s", cityid, 1))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

