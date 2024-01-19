package enrich

type apiAgeResponse struct {
	Age int `json:"age"`
}

type apiGenderResponse struct {
	Gender string `json:"gender"`
}

type apiNationalityResponse struct {
	Nationality []struct {
		CountryId string `json:"country_id"`
	} `json:"country"`
}
