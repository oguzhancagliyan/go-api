package services

import (
	"net/http"
	"yemek-sepeti-case-api/src/configuration"
	"yemek-sepeti-case-api/src/models"
)

type IKeyValueService interface {
	Set(dto models.AddKeyDto) (response models.AddKeyServiceModel)
	Get(key string) (response models.GetKeyServiceResponseModel)
	Flush()
}

type KeyValueService struct {
	db *configuration.Database
}

func NewKeyValueService(db *configuration.Database) *KeyValueService {
	return &KeyValueService{db: db}
}

func (k *KeyValueService) Set(dto models.AddKeyDto) (response models.AddKeyServiceModel) {

	if dto.Key == "" {
		response.Message = models.KeyEmptyErrorMessage
		response.StatusCode = http.StatusBadRequest
		return
	}

	_, isOk := get(k.db, dto.Key)

	if isOk {
		response.Message = models.KeyAlreadyExist
		response.StatusCode = http.StatusConflict
		return
	}

	k.db.Lock()
	k.db.Db[dto.Key] = dto.Value
	k.db.Unlock()

	response.Key = dto.Key
	response.Value = dto.Value
	response.StatusCode = http.StatusCreated
	return
}

func (k *KeyValueService) Get(key string) (response models.GetKeyServiceResponseModel) {

	if key == "" {
		response.Message = models.KeyEmptyErrorMessage
		response.StatusCode = http.StatusBadRequest
		return
	}

	value, isOk := get(k.db, key)

	if !isOk {
		response.Message = models.KeyNotFoundErrorMessage
		response.StatusCode = http.StatusNotFound
		return
	}

	response.Key = key
	response.Value = value
	response.StatusCode = http.StatusOK
	return
}

func (k *KeyValueService) Flush() {
	k.db.Lock()
	k.db.Db = make(map[string]string)
	k.db.Unlock()
}

func get(db *configuration.Database, key string) (string, bool) {
	db.RLock()
	value, isOk := db.Db[key]
	db.RUnlock()

	return value, isOk
}
