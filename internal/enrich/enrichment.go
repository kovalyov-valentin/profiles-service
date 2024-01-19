package enrich

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/kovalyov-valentin/profiles-service/internal/config"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

const (
	keyName    = "name"
	keySurname = "surname"
)

type Enrichment interface {
	GetAgeByName(name string) (int, error)
	GetGenderByName(name string) (string, error)
	GetNationalityByName(name string) (string, error)
}

type enrichment struct {
	conf *config.Config
}

func NewEnrichment(conf *config.Config) Enrichment {
	return &enrichment{
		conf: conf,
	}
}

func (e enrichment) GetAgeByName(name string) (int, error) {
	body, err := e.makeGetRequest(e.conf.Api.AgeUrl, keyName, name)
	if err != nil {
		return 0, err
	}

	var response apiAgeResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return 0, fmt.Errorf("failed to decode JSON data: %w", err)
	}

	logrus.Debugf("age is %v", response)

	if response.Age < 0 {
		return 0, errors.New("age not found")
	}

	return response.Age, nil
}

func (e enrichment) GetGenderByName(name string) (string, error) {
	body, err := e.makeGetRequest(e.conf.Api.GenderUrl, keyName, name)
	if err != nil {
		return "", err
	}

	var response apiGenderResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to decode JSON data: %w", err)
	}

	logrus.Debugf("gender: %v", response)

	if response.Gender == "" {
		return "", errors.New("gender not found")
	}

	return response.Gender, nil
}

func (e enrichment) GetNationalityByName(name string) (string, error) {
	body, err := e.makeGetRequest(e.conf.Api.NationalityUrl, keyName, name)
	if err != nil {
		return "", err
	}

	var response apiNationalityResponse

	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to decode JSON data: %w", err)
	}

	logrus.Debugf("nationality: %v", response)

	if len(response.Nationality) == 0 {
		return "", errors.New("nationality not found")
	}

	return response.Nationality[0].CountryId, nil
}

func (e enrichment) makeGetRequest(basicUrl, key, name string) ([]byte, error) {
	req, err := http.NewRequest("GET", basicUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make a GET request: %w", err)
	}

	query := req.URL.Query()
	query.Add(key, name)
	req.URL.RawQuery = query.Encode()

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
