package app

import (
	"fmt"

	"go.temporal.io/sdk/workflow"
)

func SendGreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s ! \n", name), nil
}
