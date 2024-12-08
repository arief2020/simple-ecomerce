package utils

import "tugas_akhir_example/internal/pkg/dto"

func IsIDExist(data []*dto.CityResp, id string) bool {
	for _, region := range data {
		if region.Id == id {
			return true
		}
	}
	return false
}
