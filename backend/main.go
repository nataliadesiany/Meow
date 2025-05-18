// File ini adalah entry point backend Little Alchemy 2.
// Berisi setup server Gin, middleware CORS, dan routing endpoint pencarian resep.

package main

import (
	"main/controllers" // Import controller pencarian resep
	"net/http"         // Untuk kebutuhan HTTP
	"os"               // Untuk mengambil environment variable (PORT)

	"github.com/gin-gonic/gin" // Framework web Gin
)

// Middleware CORS agar backend bisa diakses dari frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")                                // Izinkan semua origin (untuk pengembangan)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE") // Metode yang diizinkan
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Header yang diizinkan
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")                        // Izinkan kredensial

		if c.Request.Method == "OPTIONS" { // Tangani preflight request CORS
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next() // Lanjutkan ke handler berikutnya
	}
}

func main() {
	r := gin.Default()                              // Inisialisasi Gin
	r.Use(CORSMiddleware())                         // Pasang middleware CORS
	r.POST("/api/search", controllers.SearchRecipe) // Endpoint pencarian resep

	// Ambil PORT dari environment, default ke 8081 jika kosong
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081" // Untuk local development
	}

	// Jalankan server di port yang sesuai
	r.Run(":" + port)
}
