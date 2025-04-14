package main

import (
	"log"

	"github.com/abhinavkale-dev/golang-auth/controllers"
	"github.com/abhinavkale-dev/golang-auth/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := models.InitPrisma(); err != nil {
		log.Fatalf("failed to initialize Prisma client: %v", err)
	}

	defer models.PrismaClient.Disconnect()

	r := gin.Default()

	store := cookie.NewStore([]byte("secret-session-key"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.POST("/signout", controllers.SignOut)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
