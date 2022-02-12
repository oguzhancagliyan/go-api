package services

import (
	"net/http"
	"testing"
	"yemek-sepeti-case-api/src/configuration"
	"yemek-sepeti-case-api/src/models"
)

func TestGet_Should_Return_KeyEmptyErrorMessage(t *testing.T) {
	key := ""

	db := configuration.NewDatabase()
	keyValueService := NewKeyValueService(db)

	response := keyValueService.Get(key)

	if response.Message != models.KeyEmptyErrorMessage {
		t.Errorf("expected '%s', got '%s'", models.KeyEmptyErrorMessage, response.Message)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("expected '%d', got '%d'", http.StatusBadRequest, response.StatusCode)
	}

	if response.Key != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Key)
	}

	if response.Value != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Value)
	}
}

func TestGet_Should_Return_KeyNotFoundErrorMessage(t *testing.T) {
	key := "aaa"

	db := configuration.NewDatabase()
	keyValueService := NewKeyValueService(db)

	response := keyValueService.Get(key)

	if response.Message != models.KeyNotFoundErrorMessage {
		t.Errorf("expected '%s', got '%s'", models.KeyNotFoundErrorMessage, response.Message)
	}

	if response.StatusCode != http.StatusNotFound {
		t.Errorf("expected '%d', got '%d'", http.StatusNotFound, response.StatusCode)
	}

	if response.Key != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Key)
	}

	if response.Value != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Value)
	}
}

func TestGet_Should_Return_Value_Of_The_Key(t *testing.T) {
	key := "key"
	value := "value"

	db := configuration.NewDatabase()
	db.Db[key] = value
	keyValueService := NewKeyValueService(db)

	response := keyValueService.Get(key)

	if response.Message != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Message)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected '%d', got '%d'", http.StatusOK, response.StatusCode)
	}

	if response.Key != key {
		t.Errorf("expected '%s', got '%s'", key, response.Key)
	}

	if response.Value != value {
		t.Errorf("expected '%s', got '%s'", value, response.Value)
	}
}

func TestSet_Should_Return_KeyEmptyErrorMessage(t *testing.T) {
	addKeyDto := models.AddKeyDto{
		Key:   "",
		Value: "",
	}

	db := configuration.NewDatabase()
	keyValueService := NewKeyValueService(db)

	response := keyValueService.Set(addKeyDto)

	if response.Message != models.KeyEmptyErrorMessage {
		t.Errorf("expected '%s', got '%s'", models.KeyEmptyErrorMessage, response.Message)
	}

	if response.StatusCode != http.StatusBadRequest {
		t.Errorf("expected '%d', got '%d'", http.StatusBadRequest, response.StatusCode)
	}

	if response.Key != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Key)
	}

	if response.Value != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Value)
	}
}

func TestSet_Should_Return_KeyAlreadyExistMessage(t *testing.T) {
	key := "already"
	value := "exist"
	addKeyDto := models.AddKeyDto{
		Key:   key,
		Value: value,
	}

	db := configuration.NewDatabase()
	db.Db[key] = value
	keyValueService := NewKeyValueService(db)

	response := keyValueService.Set(addKeyDto)

	if response.Message != models.KeyAlreadyExist {
		t.Errorf("expected '%s', got '%s'", models.KeyAlreadyExist, response.Message)
	}

	if response.StatusCode != http.StatusConflict {
		t.Errorf("expected '%d', got '%d'", http.StatusConflict, response.StatusCode)
	}

	if response.Key != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Key)
	}

	if response.Value != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Value)
	}
}

func TestSet_Should_Return_CreatedKeyAndValue(t *testing.T) {
	key := "insert key"
	value := "insert value"
	addKeyDto := models.AddKeyDto{
		Key:   key,
		Value: value,
	}

	db := configuration.NewDatabase()
	keyValueService := NewKeyValueService(db)
	response := keyValueService.Set(addKeyDto)

	if response.Message != "" {
		t.Errorf("expected '%s', got '%s'", "", response.Message)
	}

	if response.StatusCode != http.StatusCreated {
		t.Errorf("expected '%d', got '%d'", http.StatusCreated, response.StatusCode)
	}

	if response.Key != key {
		t.Errorf("expected '%s', got '%s'", key, response.Key)
	}

	if response.Value != value {
		t.Errorf("expected '%s', got '%s'", value, response.Value)
	}
}
