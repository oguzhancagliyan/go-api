package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yemek-sepeti-case-api/src/models"
	"yemek-sepeti-case-api/src/services"
)

type KeyValueController struct {
	keyValueService services.IKeyValueService
}

func NewKeyValueController(keyValueService services.IKeyValueService) *KeyValueController {
	return &KeyValueController{keyValueService: keyValueService}
}

func (c *KeyValueController) Get(w http.ResponseWriter, key string) {
	respBody := models.NewGetKeyResponseDto()

	if key == "" {
		respBody.Error = models.NewErrorDto(models.KeyEmptyErrorMessage)
		body, err := models.CreateGetKeyResponseDtoBody(respBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	serviceResponse := c.keyValueService.Get(key)

	if serviceResponse.Message != "" {
		respBody.Error = models.NewErrorDto(serviceResponse.Message)
		body, err := models.CreateGetKeyResponseDtoBody(respBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(serviceResponse.StatusCode)
		w.Write(body)
		return
	}

	respBody.Data = models.NewData(serviceResponse.Key, serviceResponse.Value)

	body, err := models.CreateGetKeyResponseDtoBody(respBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(serviceResponse.StatusCode)
	w.Write(body)
}

func (c *KeyValueController) Set(w http.ResponseWriter, r *http.Request) {
	respBody := models.NewAddKeyResponseDto()
	var addKeyDto models.AddKeyDto
	err := json.NewDecoder(r.Body).Decode(&addKeyDto)
	if err != nil {
		respBody.Error = models.NewErrorDto(models.ModelIsNotValidMessage)
		body, err := models.CreateAddKeyResponseDtoBody(respBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	if addKeyDto.Key == "" {
		respBody.Error = models.NewErrorDto(models.KeyEmptyErrorMessage)
		body, err := models.CreateAddKeyResponseDtoBody(respBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(body)
		return
	}

	serviceResponse := c.keyValueService.Set(addKeyDto)

	if serviceResponse.Message != "" {
		respBody.Error = models.NewErrorDto(serviceResponse.Message)
		body, err := models.CreateAddKeyResponseDtoBody(respBody)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(serviceResponse.StatusCode)
		w.Write(body)
		return
	}

	respBody.Data = models.NewData(serviceResponse.Key, serviceResponse.Value)

	body, err := models.CreateAddKeyResponseDtoBody(respBody)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(serviceResponse.StatusCode)
	w.Write(body)
	w.Header().Set("Location", fmt.Sprintf("http://localhost:8081/%s", addKeyDto.Value))
}

func (c *KeyValueController) Flush(w http.ResponseWriter) {
	c.keyValueService.Flush()
	w.WriteHeader(http.StatusNoContent)
	return
}
