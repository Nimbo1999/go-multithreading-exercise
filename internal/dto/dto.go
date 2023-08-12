package dto

import "github.com/nimbo1999/go-multithreading-exercise/internal/entity"

type ViaCepOutput struct {
	Cep     entity.ViaCep `json:"cep"`
	Message string        `json:"message"`
}

type CdnApiCepOutput struct {
	Cep     entity.CdnApiCep `json:"cep"`
	Message string           `json:"message"`
}
