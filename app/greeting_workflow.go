package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SendGreetingWorkflow(ctx workflow.Context, name string) (string, error) {

	var ptGreeting string

	activityOptions := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}

	ctx = workflow.WithActivityOptions(ctx, activityOptions)

	ptGreetingFuture := workflow.ExecuteActivity(ctx, "GreetingInPortugueseActivity", name)

	if err := ptGreetingFuture.Get(ctx, &ptGreeting); err != nil {
		fmt.Println("Erro ao executar ")
	}

	return ptGreeting, nil
}
