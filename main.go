package main

import (
	"api-go/threads"
	"database/sql"
	"log"
	"os"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()

	// * Initialize database connection
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// * Initialize Clerk client
	client, err := clerk.NewClient("sk_test_EXcHIiBhyjSv1FKxPpo3g0QkMzZg5OSJYctcVFAWrx")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client)

	// * Initialize Gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// * Define a middleware function to check for authorization headers
	// authorize := func(c *gin.Context) {
	// 	sessionToken := c.GetHeader("Authorization")
	// 	if sessionToken == "" {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
	// 		c.Abort()
	// 		return
	// 	}
	// 	sessionToken = strings.TrimPrefix(sessionToken, "Bearer ")
	// 	_, err := client.VerifyToken(sessionToken)

	// 	if err != nil {
	// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
	// 		c.Abort()
	// 		return
	// 	}

	// 	c.Next()
	// }
	// router.Use(authorize)

	router.GET("/threads", threads.GetThreads(db))
	router.GET("/threads/:public_id", threads.GetThread(db))

	// * Run the server
	router.Run(":8080")
}
