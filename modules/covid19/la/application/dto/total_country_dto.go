package dto

import (
	"encoding/json"

	domSchema "github.com/muharihar/d3ta-go/modules/covid19/la/domain/schema"
)

// DisplayCurrentDataByCountryReqDTO represent DisplayCurrentDataByCountryReqDTO
type DisplayCurrentDataByCountryReqDTO struct {
	CountryCode string                `json:"countryCode"`
	Providers   []*domSchema.Provider `json:"providers"`
}

// DisplayCurrentDataByCountryResDTO represent DisplayCurrentDataByCountryResDTO
type DisplayCurrentDataByCountryResDTO struct {
	Query interface{}                             `json:"query"`
	Data  *domSchema.TotalCountryProviderDataList `json:"data"`
}

// ToJSON covert to JSON
func (r *DisplayCurrentDataByCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
