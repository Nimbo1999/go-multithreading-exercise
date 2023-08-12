package services

import (
	"fmt"
	"log"

	"github.com/nimbo1999/go-multithreading-exercise/internal/client"
	"github.com/nimbo1999/go-multithreading-exercise/internal/entity"
)

type ViaCepService struct {
	Channel chan<- entity.ViaCep
	Error   chan<- error
}

func NewViaCepService(channel chan<- entity.ViaCep, err chan<- error) ViaCepService {
	return ViaCepService{Channel: channel, Error: err}
}

func (service *ViaCepService) GetCep(cep string) {
	viacep, err := client.RequestCep[entity.ViaCep](fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
	if err != nil {
		log.Println(err)
		service.Error <- err
		return
	}
	service.Channel <- viacep
}
