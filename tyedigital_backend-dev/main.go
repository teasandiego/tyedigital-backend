/*
**************************************************************************************************************************
This code is licensed under Kushmanda Tech LLC license
Property of Kushmanda Tech LLC, Do not distribute to unauthorized person. Do not print / copy without explicit permission.
App name : ARIES 1.0
Purpose : This file contains main settings of goLang Config.
Brief :
Developer : Rajendro Sau.
email : contact@kushmanda.tech Phone : +1 (760) 659-0487
***************************************************************************************************************************
*/
package main

import (
	"log"
	"os"

	routers "github.com/Kushmanda-Tech/tyedigital_backend/Api/V1/Routers"
	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
	mysql "github.com/Kushmanda-Tech/tyedigital_backend/Db/MySql"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Rest of your code...

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	log.Println("DB_PORT: ", os.Getenv("DB_PORT"))
	log.Println("DB_PASSWORD: ", os.Getenv("DB_PASSWORD"))
	log.Println("DB_USERNAME: ", os.Getenv("DB_USERNAME"))
	log.Println("DB_HOST: ", os.Getenv("DB_HOST"))
	log.Println("SERVER_NAME: ", config.DbConfig().SERVER_NAME)
}
func main() {

	// config.CreatDesign()
	mysql.Connect(&gin.Context{})
	routers.HandleRequest()
}
