package threads

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetThreads(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		threads, err := GetMany(db)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, threads)
	}
}

func GetThread(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		publicID := c.Param("public_id")
		thread, err := GetOne(db, publicID)

		if err != nil {
			switch err {
			case sql.ErrNoRows:
				c.JSON(http.StatusNotFound, gin.H{"error": "Thread not found"})
				return
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(http.StatusOK, thread)
	}
}
