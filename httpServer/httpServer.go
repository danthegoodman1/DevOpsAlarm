package httpserver

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Server *HTTPServer
)

type HTTPServer struct {
	Echo *echo.Echo
}

func StartHTTPServer() {
	echoInstance := echo.New()
	Server = &HTTPServer{
		Echo: echoInstance,
	}

	Server.Echo.HideBanner = true
	Server.Echo.Use(middleware.Logger())

	Server.registerRoutes()
	SetupTicker() // Start the Alarm status ticker

	log.Println("Starting DevOps on port 80")
	Server.Echo.Logger.Fatal(Server.Echo.Start(":80"))
}

func (s *HTTPServer) registerRoutes() {
	// Health check route
	s.Echo.GET("/hc", healthCheck)

	s.Echo.POST("/alarm", Post_Alarm)
	s.Echo.GET("/ack", Get_Ack)
}

func healthCheck(c echo.Context) error {
	return c.String(200, "Alive!")
}
