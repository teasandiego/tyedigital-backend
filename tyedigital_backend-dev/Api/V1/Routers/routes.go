package routers

import (
	"log"
	"net/http"
	"time"

	auth "github.com/Kushmanda-Tech/tyedigital_backend/Api/V1/Controllers/Auth"
	chapter "github.com/Kushmanda-Tech/tyedigital_backend/Api/V1/Controllers/Chapter"
	project "github.com/Kushmanda-Tech/tyedigital_backend/Api/V1/Controllers/Project"
	config "github.com/Kushmanda-Tech/tyedigital_backend/Config"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

/* This CORS Function is required for APIs to work on React as well Mobile*/

func CORS(c *gin.Context) {
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

/* This HandleRequest Function is required for Routing APIs to Functions*/

func HandleRequest() {
	log.Println("Before Open  This Server  Run This Command in server 'sudo ufw allow 7070'")
	log.Println("Quit The Server With CONTROL-C.")
	// gin.Mode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORS)
	router.Use(cors.New(cors.Config{
		// AllowOrigins: []string{"https://dev.tyedigital.org"},
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	gin.SetMode(gin.ReleaseMode)
	router.GET("/verify_email", auth.PutVerifyEmail())
	v1 := router.Group(config.API_VERSION)
	{
		authRoute := v1.Group("/User_Authentication")
		{
			authRoute.POST("", auth.PostUserDetails())
			authRoute.GET("sd_get_student_details", auth.GetStudentDetails())
			authRoute.GET("pd_get_project", project.GetProjectDetails())
			authRoute.GET("cd_get_chapter", chapter.GetChapterDetails())

		}

	}
	router.LoadHTMLGlob("ThirdPartyAPI/Email/templates/*.html")
	router.Run(":" + config.SERVER_PORT)
}
