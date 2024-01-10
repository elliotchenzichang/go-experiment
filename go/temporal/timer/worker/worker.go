package main

import (
	"go-learn/go/temporal/timer"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"log"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort: client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "timer_queue", worker.Options{
		MaxConcurrentActivityExecutionSize: 3,
	})

	w.RegisterWorkflow(timer.SampleTimerWorkflow)
	w.RegisterActivity(timer.SayHelloWorldActivity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
