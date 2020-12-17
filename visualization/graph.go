package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	pusher "github.com/pusher/pusher-http-go"
	"sync"

	//"fmt"
	//"math/rand"
	"net/http"
	"sorting-algorithms/Randomizer"
)

// We register the Pusher client
var client = pusher.Client{
	AppID:   "1119581",
	Key:     "9515c01265248c7e86e8",
	Secret:  "06516341e457fe5988fd",
	Cluster: "us2",
	Secure:  true,
}

// visitsData is a struct
type visitsData struct {
	Pages int
	Count int
}

type arrayElement struct {
	Value int
	Label string
}

type arrayStruct struct {
	Elements [1000]arrayElement
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the HTTP routes
	e.File("/", "visualization/public/index.html")
	e.File("/style.css", "visualization/public/style.css")
	e.File("/app.js", "visualization/public/app.js")
	e.GET("/visualize", visualize)

	// Start server
	e.Logger.Fatal(e.Start("localhost:9000"))
}

func visualize(c echo.Context) error {

	array := Randomizer.RandomArray(10)

	var waitGroup sync.WaitGroup

	for _, number := range array {
		waitGroup.Add(1)
		go addNumber(number, &waitGroup)
	}
	waitGroup.Wait()

	client.Trigger("arrayVisualization", "update", 0)

	return c.String(http.StatusOK, "Simulation begun")
}

func addNumber( pNumber int, pWaitGroup *sync.WaitGroup) {
	defer pWaitGroup.Done()
	arrayValue := arrayElement{
		Value: pNumber,
		Label: "",
	}
	client.Trigger("arrayVisualization", "addNumber", arrayValue)
}