package main

import (
	"fmt"
	"net/http"

	instana "github.com/instana/go-sensor"
	"github.com/tullo/otel-workshop/web/fib"
)

func ServeInstana(sensor *instana.Sensor) error {
	span := sensor.Tracer().StartSpan("ServeInstana")
	defer span.Finish()

	mux := http.NewServeMux()
	mux.Handle("/", instana.TracingHandlerFunc(sensor, "/", fib.RootHandler))
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/fib", instana.TracingHandlerFunc(sensor, "/fib/{n}", fib.FibHandler))
	mux.Handle("/fibinternal", instana.TracingHandlerFunc(sensor, "/fibinternal/{n}", fib.FibHandler))

	fmt.Println("Your server is live!\nTry to navigate to: http://127.0.0.1:3000/fib?n=6")
	if err := http.ListenAndServe("127.0.0.1:3000", mux); err != nil {
		return fmt.Errorf("could not start web server: %w", err)
	}

	return nil
}
