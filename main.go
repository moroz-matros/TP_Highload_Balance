package main

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var HitsCount = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "foo_total",
	Help: "Number of foo successfully processed.",
})


func main() {
	serv := echo.New()

	serv.GET("/", Handle)


	err := prometheus.Register(HitsCount)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		metricsRouter := echo.New()
		metricsRouter.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		log.Fatal(metricsRouter.Start(":8088"))
	}()


	err = serv.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}


}

func Handle(c echo.Context) error{
	defer HitsCount.Add(1)
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
	return c.String(http.StatusOK, "Hello from server")
}