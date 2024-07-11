package models

import "gorm.io/gorm"

type History struct {
	gorm.Model

	Command string
}
