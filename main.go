package main

import (
	"log"
	"os"
	"os/signal"

	instana "github.com/instana/go-sensor"
)

func main() {
	sensor := instana.NewSensor("fib-service")
	// sensor := instana.NewSensorWithTracer(
	// 	instana.NewTracerWithOptions(&instana.Options{
	// 		//AgentHost: "127.0.0.1",
	// 		//AgentPort: 42699,
	// 	},
	// 	),
	// )

	sensor.Tracer()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	errCh := make(chan error)

	// Start web server.
	l := log.New(os.Stdout, "", 0)
	go func() {
		errCh <- ServeInstana(sensor)
	}()

	select {
	case <-sigCh:
		l.Println("\ngoodbye")
		return
	case err := <-errCh:
		if err != nil {
			l.Fatal(err)
		}
	}

}
