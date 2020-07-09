package domain

import (
	"log"
	uuid "github.com/satori/go.uuid"
	"time"
	"github.com/jinzhu/gorm"
)

type Base struct{
	ID string `json:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `json:"create_at" `
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" sql:"index"`
}

func (base *Base) BeforeCreate(scope *gorm.Scope) error  {
	err := scope.SetColumn("CreatedAt", time.Now())

	if err!= nil{
		log.Fatalf("error during obj creating %v",err)
	}
	err = scope.SetColumn("ID", uuid.NewV4().String())

	if err != nil{
		log.Fatalf("error during obj creating %v",err)
	}
	return nil
}