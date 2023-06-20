package types

import "gorm.io/gorm"

type API_KEY struct {
  gorm.Model
  ApiKey string
}
