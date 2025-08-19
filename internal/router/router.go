package router

import (
	"fmt"
	"os"
	"svi-be/internal/handler"
	"svi-be/internal/model"
	"svi-be/internal/repository"
	"svi-be/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&model.Posts{})

	repo := repository.NewArticleRepository(db)
	svc := service.NewArticleService(repo)
	h := handler.NewArticleHandler(svc)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:  []string{"Origin", "Content-Type", "Authorization", "User-Agent", "X-Device-Info", "X-Longitude", "X-Latitude", "X-Source-System"},
		ExposeHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	api := r.Group("/api/article")
	{

		api.POST("/", h.CreateArticle)
		api.GET("/", h.GetAll)
		api.GET("/:id", h.GetDetail)
		api.PUT("/delete/:id", h.DeleteArtikel)
		api.PUT("/update/:id", h.UpdateArtikel)

	}

	return r
}
