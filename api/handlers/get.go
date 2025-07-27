package api_handlers

import (
	"github.com/IliaKoshkin/cloud-app/internal"
	"github.com/IliaKoshkin/cloud-app/models"
)

func Get(key string) (string, error) {
	value, ok := models.Store[key]
	if !ok {
		return "", internal.ErrorNoSuchKey
	}
	return value, nil
}
