package models

import (
	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	Judul        string
	Nama_penulis string
	Tahun        string
}
