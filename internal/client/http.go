package client

import (
	"encoding/json"
	"net/http"

	"github.com/nimbo1999/go-multithreading-exercise/internal/entity"
)

func NewClient() *http.Client {
	return &http.Client{}
}

func RequestCep[Data entity.ViaCep | entity.CdnApiCep](url string) (Data, error) {
	var data Data
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return data, err
	}
	client := NewClient()
	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
