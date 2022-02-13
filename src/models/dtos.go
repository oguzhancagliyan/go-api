package models

import (
	"encoding/json"
	"errors"
)

//Error

type ErrorDto struct {
	Message string `json:"message,omitempty"`
}

func NewErrorDto(message string) *ErrorDto {
	return &ErrorDto{Message: message}
}

//Data

type Data struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewData(key string, value string) *Data {
	return &Data{Key: key, Value: value}
}

//Add Key
type AddKeyDto struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AddKeyResponseDto struct {
	Data  *Data     `json:"data,omitempty"`
	Error *ErrorDto `json:"error,omitempty"`
}

func NewAddKeyResponseDto() *AddKeyResponseDto {
	return &AddKeyResponseDto{}
}

func CreateAddKeyResponseDtoBody(dto *AddKeyResponseDto) ([]byte, error) {
	if dto == nil {
		return nil, errors.New("dto can not be nil")
	}

	body, err := json.Marshal(dto)

	if err != nil {
		return nil, err
	}

	return body, nil
}

type AddKeyServiceModel struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
}

//Get Key

func CreateGetKeyResponseDtoBody(dto *GetKeyResponseDto) ([]byte, error) {
	if dto == nil {
		return nil, errors.New("dto can not be nil")
	}

	body, err := json.Marshal(dto)

	if err != nil {
		return nil, err
	}

	return body, nil
}

type GetKeyResponseDto struct {
	Data  *Data     `json:"data,omitempty"`
	Error *ErrorDto `json:"error,omitempty"`
}

func NewGetKeyResponseDto() *GetKeyResponseDto {
	return &GetKeyResponseDto{}
}

type GetKeyServiceResponseModel struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message,omitempty"`
}
