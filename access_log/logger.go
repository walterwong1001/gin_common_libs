package accesslog

import (
	"context"
	"fmt"
)

// Logger is an interface that defines a logging method.
// Implementations of this interface are responsible for processing the log data.
type Logger interface {
	// Log logs the provided metrics.
	// The context can be used to pass additional information or to manage timeouts.
	// Returns an error if logging fails.
	Log(ctx context.Context, metric map[string]any) error
}

// ConsoleLog is a simple logger that prints the metrics to the console.
// Useful for debugging or local development.
type ConsoleLog struct{}

// Log prints the metrics to the console.
// It iterates over the key-value pairs in the metrics map and prints each pair.
// Returns nil as there are no errors expected in console logging.
func (l *ConsoleLog) Log(_ context.Context, metric map[string]any) error {
	for k, v := range metric {
		fmt.Printf("%s: %v \n", k, v)
	}
	return nil
}
