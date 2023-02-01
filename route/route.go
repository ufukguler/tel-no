package routes

import (
	"github.com/labstack/echo/v4"
	"telno/handlers"
)

func UseRoute(e *echo.Echo) {
	api := e.Group("/api")
	{
		api.GET("/phoneNumber", handlers.FindByPhoneNumber)
		api.POST("/phoneNumber", handlers.AddComment)

		api.GET("/image", handlers.GetImagePhoneNumber)
		api.GET("/latest", handlers.GetLatestUpdates)

		api.GET("/sitemap/main.xml", handlers.GetSitemapXML)       // MAIN sitemap file
		api.GET("/sitemap/sitemap.xml.gz", handlers.GetSitemapXML) // MAIN sitemap file
		api.GET("/sitemap/pages.xml", handlers.GetSitemapPages)    // top pages sitemap file
		api.GET("/sitemap/numbers/:id", handlers.GetSitemap)       // segmented numbers
	}

	admin := e.Group("/api/admin")
	{
		admin.GET("/phoneNumber/:number", handlers.FindByPhoneNumberOnlyFalseComments)
		admin.GET("/phoneNumber", handlers.FindPhoneNumbersByPageable)
		admin.GET("/phoneNumberUnchecked", handlers.FindPhoneNumbersCommentUncheckedByPageable)
		admin.GET("/comment", handlers.FindCommentById)
		admin.POST("/comment", handlers.UpdateCommentById)
		admin.POST("/quickComment", handlers.QuickUpdateCommentById)
		admin.DELETE("/comment", handlers.DeleteCommentById)
	}
}
