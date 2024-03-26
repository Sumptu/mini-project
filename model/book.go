package model

import (
	"gorm.io/gorm"
)
type Book struct {
	Model
	ISBN      string         `json:"isbn"`
	Penulis   string         `json:"penulis"`
	Tahun     uint           `json:"tahun"`
	Judul     string         `json:"judul"`
	Gambar    string         `json:"gambar"`
	Stok      uint           `json:"stok"`
}

func (bk *Book) Create(db *gorm.DB) error {
	err := db.
		Model(Book{}).
		Create(&bk).
		Error

	if err != nil {
		return err
	}

	return nil
}

func (bk *Book) GetByID(db *gorm.DB) (Book, error) {
	res := Book{}


	err := db.
	Model(Book{}).
	Where("id = ?", bk.Model.ID).
	Take(&res).
	Error

	if err != nil {
		return Book{}, err
	}

	return res, nil
}

func (bk *Book) GetAll(db *gorm.DB) ([]Book, error) {
	res := []Book{}

	err := db.
	Model(Book{}).
	Find(&res).
	Error

	if err != nil {
		return []Book{}, err
	}

	return res, nil
}

func (bk *Book) UpdateOne(db *gorm.DB) error {
	return db.
		Model(Book{}).
		Select("isbn", "penulis", "judul", "tahun", ).
		Where("id = ?", bk.Model.ID).
		Updates(map[string]any{
			"isbn": bk.ISBN,
			"penulis": bk.Penulis,
			"tahun": bk.Tahun,
			"judul": bk.Judul,
			"gambar": bk.Gambar,
			"stok": bk.Stok,
		}).
		Error
}

func (bk *Book) DeleteByID(db *gorm.DB) error {
	return db.
	Model(Book{}).
	Where("id = ?", bk.Model.ID).
	Delete(&bk).
	Error
}