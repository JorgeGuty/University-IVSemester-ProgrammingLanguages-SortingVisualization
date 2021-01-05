package main

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	pusher "github.com/pusher/pusher-http-go"
	"sorting-visualization/sorting-algorithms/Randomizer"
	"strconv"
	"sync"
	//"fmt"
	//"math/rand"
	"net/http"
)


var globalArray []int

// We register the Pusher client
var client = pusher.Client{
	AppID:   "1119581",
	Key:     "9515c01265248c7e86e8",
	Secret:  "06516341e457fe5988fd",
	Cluster: "us2",
	Secure:  true,
}


type arrayElement struct {
	Value int
	Label string
}

/*
SortIDs:
heap
bubble
insertion
quick
tree
selection
*/
type swapElement struct {
	Index1 int
	Index2 int
	SortID string
}


func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the HTTP routes
	e.File("/", "visualization/public/index.html")
	e.Static("/style.css", "visualization/public/style.css")
	e.Static("/app.js", "visualization/public/app.js")
	e.GET("/visualize/:n", visualize)
	e.GET("/solve", solve)

	// Start server
	e.Logger.Fatal(e.Start(":11000"))
}



func visualize(c echo.Context) error {

	n, err := strconv.Atoi(c.Param("n"))

	if err !=  nil {
		fmt.Println(err)
	}

	array := Randomizer.RandomArray(n)

	globalArray = array

	fmt.Println(globalArray)

	var waitGroup sync.WaitGroup

	for _, number := range array {
		waitGroup.Add(1)
		go addNumber(number, &waitGroup)
		waitGroup.Wait()
	}

	client.Trigger("arrayVisualization", "update", 0)

	return c.String(http.StatusOK, "Visualization done")
}


func addNumber( pNumber int, pWaitGroup *sync.WaitGroup) {
	defer pWaitGroup.Done()
	arrayValue := arrayElement{
		Value: pNumber,
		Label: "",
	}
	client.Trigger("arrayVisualization", "addNumber", arrayValue)
}

func solve(c echo.Context) error {

	bubbleArray := make([]int, len(globalArray))
	insertionArray := make([]int, len(globalArray))
	heapArray := make([]int, len(globalArray))
	quickArray := make([]int, len(globalArray))
	selectionArray := make([]int, len(globalArray))
	treeArray := make([]int, len(globalArray))

	copy(heapArray, globalArray)
	copy(quickArray, globalArray)
	copy(bubbleArray, globalArray)
	copy(insertionArray, globalArray)
	copy(selectionArray, globalArray)
	copy(treeArray, globalArray)

	var waitGroup sync.WaitGroup
	waitGroup.Add(6)

	waitGroup.Done()
	waitGroup.Done()
	waitGroup.Done()
	waitGroup.Done()
	waitGroup.Done()
	waitGroup.Done()

	waitGroup.Wait()

	return c.String(http.StatusOK, "Simulation begun")

}

// Swap
func swap( pIndex1 int, pIndex2 int, sortID string ){
	swapData := swapElement{
		Index1: pIndex1,
		Index2: pIndex2,
		SortID: sortID,
	}

	client.Trigger("arrayVisualization", "swap", swapData)
}