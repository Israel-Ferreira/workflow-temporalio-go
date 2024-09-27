package main

import (
	"fmt"
	"log"

	"github.com/Israel-Ferreira/myfirstwf/app"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	fmt.Println("Criando um Worker com Temporal")

	c, err := client.Dial(client.Options{})

	if err != nil {
		log.Fatalln("Erro ao rodar o Worker")
	}

	defer c.Close()

	w := worker.New(c, "greeting-tasks", worker.Options{})

	w.RegisterWorkflow(app.SendGreetingWorkflow)

	if err = w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Erro ao iniciar o worker")
	}

}
