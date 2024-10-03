package app

import (
	"context"
	"fmt"
)

func GreetingInPortugueseActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Olá %s", name), nil
}

func GreetingInSpanishActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("!Hola %s", name), nil
}

func GreetingInEnglishActivity(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}
