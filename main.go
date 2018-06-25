package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Model is a test model to insert
type Model struct{
	ID string `gorm:"column:id"`
	RecTime time.Time `gorm:"column:rec_time"`
}

// TableName for model
func (Model) TableName() string{return "model"}

// BeforeCreate model hook
func (m *Model) BeforeCreate() error {
	m.RecTime = time.Now()
	return nil
}

func main(){
	db, err := gorm.Open("postgres", "host=localhost port=5432 dbname=test_gorm_db sslmode=disable user=postgres")
	if err != nil {
		panic(err)
	}

	modelToCreate := &Model{
		ID: "34f620bf-b456-4397-a434-cc157b2b8633",
	}

	if err := db.Debug().Create(&modelToCreate).Error; err != nil {
		fmt.Println(">>>>> CREATE:", err)
	}

	var i int
	if err := db.Debug().Model(Model{}).Count(&i).Error; err != nil {
		fmt.Println(">>>>> COUNT:", err)
	}

	fmt.Printf("%d rows found\n", i)

	fmt.Println("That's all.")
}
