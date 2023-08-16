package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nimbo1999/go-multithreading-exercise/internal/dto"
	"github.com/nimbo1999/go-multithreading-exercise/internal/entity"
	"github.com/nimbo1999/go-multithreading-exercise/internal/services"
	"github.com/nimbo1999/go-multithreading-exercise/pkg/utils"
)

type CepHandler struct {
}

func (handler *CepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := utils.FormatCep(chi.URLParam(r, "cep"))
	if len(cep) != 9 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Invalid cep: %s", cep)))
		return
	}

	viacepChannel := make(chan entity.ViaCep)
	cdnApiCepChannel := make(chan entity.CdnApiCep)
	errChannel := make(chan error)
	encoder := json.NewEncoder(w)

	viaCepService := services.NewViaCepService(viacepChannel, errChannel)
	cdnApiCepService := services.NewCdnApiCepService(cdnApiCepChannel, errChannel)

	go viaCepService.GetCep(cep)
	go cdnApiCepService.GetCep(cep)

	select {
	case data := <-viacepChannel:
		w.WriteHeader(http.StatusOK)
		encoder.Encode(dto.ViaCepOutput{
			Cep:     data,
			Message: "Dados retornados da Api, ViaCep",
		})
		return
	case data := <-cdnApiCepChannel:
		w.WriteHeader(http.StatusOK)
		encoder.Encode(dto.CdnApiCepOutput{
			Cep:     data,
			Message: "Dados retornados da Api, CdnApiCep",
		})
		return
	case err := <-errChannel:
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	case <-time.After(time.Second):
		w.WriteHeader(http.StatusRequestTimeout)
		w.Write([]byte("Request timeout"))
		return
	}
}
