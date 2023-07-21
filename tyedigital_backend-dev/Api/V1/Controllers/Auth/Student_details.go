package auth

import (
	"log"
	"strconv"

	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
	mysql "github.com/Kushmanda-Tech/tyedigital_backend/Db/MySql"
	"github.com/gin-gonic/gin"
)

func PostUserDetails() gin.HandlerFunc {
	fn := func(c *gin.Context) {

	}
	return gin.HandlerFunc(fn)
}

func GetStudentDetails() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if err := c.ShouldBindJSON(&config.Input); err != nil {
			log.Println("config.Input ERROR ", err.Error())
			c.JSON(200, gin.H{
				"status":  config.INPUTERROR,
				"message": err.Error(),
			})
		} else {
			int_userID, _ := strconv.Atoi(config.Input.UserID)
			user_details_using_userID := mysql.Get_Student_Details_By_User_Id_FrDb(int_userID)
			if len(user_details_using_userID) == 0 {
				c.JSON(200, gin.H{
					"message": "No users found",
					"status":  204,
				})
			} else {
				c.JSON(200, gin.H{
					"message": "Users getting successfully",
					"data":    user_details_using_userID,
					"status":  200,
				})
			}
		}
	}
	return gin.HandlerFunc(fn)
}

func PutVerifyEmail() gin.HandlerFunc {
	fn := func(c *gin.Context) {

	}
	return gin.HandlerFunc(fn)
}
