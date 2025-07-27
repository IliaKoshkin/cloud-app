package api_handlers

import (
	"github.com/IliaKoshkin/cloud-app/models"
)

func Delete(key string) error {
	delete(models.Store, key)
	return nil
}
