package chapter

import (
	"log"
	"strconv"

	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"
	mysql "github.com/Kushmanda-Tech/tyedigital_backend/Db/MySql"

	"github.com/gin-gonic/gin"
)

func GetChapterDetails() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		if err := c.ShouldBindJSON(&config.Input); err != nil {
			log.Println("config.Input ERROR ", err.Error())
			c.JSON(200, gin.H{
				"status":  config.INPUTERROR,
				"message": err.Error(),
			})
		} else {
			int_chapter_id, _ := strconv.Atoi(config.Input.ChapterId)
			chapter_details := mysql.GetChapterDetailsFromDb(int_chapter_id)
			if err != nil {
				log.Println(err.Error())
				c.JSON(200, gin.H{
					"message": "Sorry we can't get this Chapter details ",
					"status":  451,
				})
			} else {
				c.JSON(200, gin.H{
					"message": "Chapter details getting successfully",
					"data":    chapter_details,
					"status":  200,
				})
			}
		}

	}
	return gin.HandlerFunc(fn)
}
