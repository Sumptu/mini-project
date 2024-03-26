package model_test

import (
	"fmt"
	"mini-project/config"
	"mini-project/model"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using global env")
	}
	config.OpenDB()
}

func TestCreateCar(t *testing.T) {
	Init()

	err := ImportDataFromCSV("sample_books.csv")
    if err != nil {
        t.Fatalf("Failed to import data from CSV: %v", err)
    }

	bookData := model.Book{
		 ISBN: "978-33",
	}

	err = bookData.Create(config.Mysql.DB)

	assert.Nil(t, err)

	fmt.Println(bookData.ID)
}

func TestGetByID(t *testing.T) {
	Init()

	carData := model.Car{
		Model: model.Model{
			ID: 1,
		},
	}

	data, err := carData.GetByID(config.Mysql.DB)
	assert.Nil(t, err)

	fmt.Println(data)
}

func TestGetAll(t *testing.T) {
	Init()

	carData := model.Car{
		Nama: "toyota",
		Tipe: "supra",
		Tahun: "1999",
	}

	err := carData.Create(config.Mysql.DB)
	assert.Nil(t, err)

	res, err := carData.GetAll(config.Mysql.DB)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(res), 1)

	fmt.Println(res)
}

func TestUpdateOne(t *testing.T) {
	Init()

	carData := model.Car{
		Model: model.Model{
			ID: 3,
		},
		Nama: "avanza",
		Tipe: "mobil",
		Tahun: "1999",
	}
	
	err := carData.UpdateOne(config.Mysql.DB)
	assert.Nil(t, err)

}

func TestDeleteByID(t *testing.T) {
	Init()

	carData := model.Car{
		Model: model.Model{
			ID: 2,
		},
	}

	err := carData.DeleteByID(config.Mysql.DB)
	assert.Nil(t, err)
	
}