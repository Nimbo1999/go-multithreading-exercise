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
	if cdnApiCep.Ok {
		service.Channel <- cdnApiCep
	}
}
