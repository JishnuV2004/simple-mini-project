package main

import (
	"fmt"
	config "gorm/Config"
	controllers "gorm/Controllers"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB
	config.InitDB()

	r := gin.Default()

	// Load all templates
	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	tmpl = template.Must(template.ParseGlob("templates/admin/dashboard.html"))
tmpl = template.Must(template.ParseGlob("templates/admin/users.html"))

	// Print loaded templates for debugging
	for _, t := range tmpl.Templates() {
		fmt.Println("Loaded template:", t.Name())
	}

	// Set templates for Gin
	r.SetHTMLTemplate(tmpl)

	// Serve static files if needed
	r.Static("/static", "static")

	// Admin routes
	r.GET("/admin/dashboard", controllers.Dashboard) // renders only dashboard.html
	r.GET("/admin/users", controllers.Users)         // renders only users.html

	// Start server
	r.Run(":8080")
}
