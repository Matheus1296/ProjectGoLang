package main

import (
	"fmt"
	"log"
	"github.com/Matheus1296/ProjectGoLang/framework/utils"
	"github.com/Matheus1296/ProjectGoLang/application/repositories"
	"github.com/Matheus1296/ProjectGoLang/domain"
)

func main()  {

	db:= utils.ConnectDB()

	user:= domain.User{
		Name:"Teste",
		Email:"teste1@teste.com",
		Password:"123",
	}

	userRepo := repositories.UserRepositoryDb{Db:db}

	result,err := userRepo.Inser(&user)

	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(result.Name,result.Email,result.Token)
}