package main

import "github.com/cixtor/middleware"

func main() {
	app := NewApp("reviews.db")

	router := middleware.New()

	router.Port = "8080"
	router.ReadTimeout = 5
	router.WriteTimeout = 10

	router.STATIC("public", "/assets")

	router.POST("/reviews/save", app.ReviewsSave)
	router.GET("/reviews/list", app.ReviewsList)
	router.GET("/admin", app.Admin)
	router.GET("/", app.Index)

	router.ListenAndServe()
}
