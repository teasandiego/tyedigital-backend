package project

import (
	"log"
	"strconv"

	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
	mysql "github.com/Kushmanda-Tech/tyedigital_backend/Db/MySql"
	"github.com/gin-gonic/gin"
)

func GetProjectDetails() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if err := c.ShouldBindJSON(&config.Input); err != nil {
			log.Println("config.Input ERROR ", err.Error())
			c.JSON(200, gin.H{
				"status":  config.INPUTERROR,
				"message": err.Error(),
			})
		} else {
			int_project_id, _ := strconv.Atoi(config.Input.ProjectId)
			project_details := mysql.GetProjectDetailsFromDb(int_project_id)
			if err != nil {
				log.Println(err.Error())
				c.JSON(200, gin.H{
					"message": "Sorry we can't get this Chapter details ",
					"status":  451,
				})
			} else {
				c.JSON(200, gin.H{
					"message": "Chapter details getting successfully",
					"data":    project_details,
					"status":  200,
				})
			}
		}

	}
	return gin.HandlerFunc(fn)
}
