package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-module/component"
	gormpgstore "go-module/component/datastore/gorm/postgresql"
	"go-module/component/uploadprovider"
	"go-module/middleware"
	ginauth "go-module/modules/auth/transport/gin"
	ginpost "go-module/modules/post/transport/gin"
	ginupload "go-module/modules/uploadfile/transport"
	ginuser "go-module/modules/user/transport/gin"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	gormDB, err := gormpgstore.NewDB(os.Getenv("DB_CONNECTION_STRING"))
	if err != nil {
		log.Fatalln(err)
	}

	uploadProvider, err :=
		uploadprovider.NewS3Provider(
			os.Getenv("S3_BUCKET"),
			os.Getenv("S3_KEY"),
			os.Getenv("S3_SECRET"),
			os.Getenv("S3_REGION"),
			os.Getenv("S3_URL"))

	if err != nil {
		log.Fatalln(err)
	}

	runGinService(gormDB, uploadProvider)
}

func runGinService(gormDB *gorm.DB, provider component.UploadFileProvider) error {

	router := gin.Default()

	ctx := component.NewAppContext(component.Configuration{
		GormDB:             gormDB,
		UploadFileProvider: provider,
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middleware.Recover(ctx))

	router.POST("/api/v1/auth/register", ginauth.RegisterUser(ctx))
	router.POST("/api/v1/auth/login", ginauth.Login(ctx))

	router.GET("/api/v1/role", middleware.RequiredAuth(ctx, "ADMIN"), ginuser.GetRoles(ctx))
	router.POST("/api/v1/role", middleware.RequiredAuth(ctx, "ADMIN"), ginuser.CreateRole(ctx))
	router.PUT("/api/v1/role/:id", middleware.RequiredAuth(ctx, "ADMIN"), ginuser.UpdateRole(ctx))

	router.POST("/api/v1/user/add-role", middleware.RequiredAuth(ctx, "ADMIN"), ginuser.CreateRoleForUser(ctx))

	router.GET("/api/v1/post/find", ginpost.ListPost(ctx))
	router.GET("/api/v1/post/find/:id", ginpost.GetPost(ctx))
	router.POST("/api/v1/post", middleware.RequiredAuth(ctx, "ADMIN"), ginpost.CreatePost(ctx))
	router.PUT("/api/v1/post/:id", middleware.RequiredAuth(ctx, "ADMIN"), ginpost.UpdatePost(ctx))
	router.DELETE("/api/v1/post/:id", middleware.RequiredAuth(ctx, "ADMIN"), ginpost.DeletePost(ctx))

	router.POST("/api/v1/upload", middleware.RequiredAuth(ctx, "ADMIN"), ginupload.UploadFile(ctx))
	return router.Run(":" + os.Getenv("SERVER_PORT"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
