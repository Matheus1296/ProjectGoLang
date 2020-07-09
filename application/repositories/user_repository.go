package repositories

import (
	"log"
	"github.com/jinzhu/gorm"
	"github.com/Matheus1296/ProjectGoLang/domain"
)

type UserRepository interface{
	Insert(user *domain.User) (*domain.User, error)
}

type UserRepositoryDb struct{
	Db *gorm.DB
}

func (repo UserRepositoryDb) Inser(user *domain.User) (*domain.User,error)  {
	err:= user.Prepare()

	if err != nil{
		log.Fatalf("error during the user validation %v",err)
		return user,err
	}

	err= repo.Db.Create(user).Error

	if err!=nil{
		log.Fatalf("error to persis user: %v",err)
		return user,err
	}

	return user,nil
}