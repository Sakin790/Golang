
package main

import (
  "net/http"
  "time"
  "log"
  "os"
  "os/signal"
  "syscall"
  "context"

  "github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
		c.Header("X-Gin-Handler", "root")
		c.JSON(http.StatusOK, gin.H{"matched": "root", "path": c.Request.URL.Path})
	})

  router.GET("/api/v1/users", func(c *gin.Context) {
		c.Header("X-Gin-Handler", "users")
		c.JSON(http.StatusOK, gin.H{"matched": "users", "path": c.Request.URL.Path})
	})

  router.GET("/api/v1/items/:id", func(c *gin.Context) {
		c.Header("X-Gin-Handler", "item-detail")
		c.JSON(http.StatusOK, gin.H{"matched": "item-detail", "id": c.Param("id")})
	})

  router.Any("/legacy/*filepath", func(c *gin.Context) {
		c.Header("X-Gin-Handler", "legacy-wildcard")
		c.JSON(http.StatusOK, gin.H{"matched": "legacy-wildcard", "path": c.Request.URL.Path})
	})
  

  srv := &http.Server{
    Addr:           ":8080",
    Handler:        router,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

go func() {
    if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
      log.Fatalf("listen: %s\n", err)
    }
  }()
  quit := make(chan os.Signal, 1)
  signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
  <-quit
  log.Println("Shutdown Server ...")

  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  if err := srv.Shutdown(ctx); err != nil {
    log.Println("Server Shutdown:", err)
  }
  log.Println("Server exiting")

}
