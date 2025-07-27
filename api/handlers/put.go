package api_handlers

import (
	"github.com/IliaKoshkin/cloud-app/models"
)

func Put(key string, value string) error {
	models.Store[key] = value
	return nil
}
