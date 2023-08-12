package services

import (
	"fmt"
	"log"

	"github.com/nimbo1999/go-multithreading-exercise/internal/client"
	"github.com/nimbo1999/go-multithreading-exercise/internal/entity"
)

type CdnApiCepService struct {
	Channel chan<- entity.CdnApiCep
	Error   chan<- error
}

func NewCdnApiCepService(channel chan<- entity.CdnApiCep, err chan<- error) CdnApiCepService {
	return CdnApiCepService{Channel: channel, Error: err}
}

func (service *CdnApiCepService) GetCep(cep string) {
	cdnApiCep, err := client.RequestCep[entity.CdnApiCep](fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", cep))
	if err != nil {
		log.Println(err)
		service.Error <- err
		return
	}
	/*
		Sometimes this api return a 429 Http status code, it means that it is receiving
		too many requests.
		For now I'm ignoring it, but later we can verify it and create an error,
		or just don't commit to the channel and wait for ViaCep or Timeout response.
	*/
	service.Channel <- cdnApiCep
}
