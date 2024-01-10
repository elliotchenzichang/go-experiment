package temporal

import (
	"context"
	"go.temporal.io/sdk/activity"
)

func SayHelloWorldActivity(ctx context.Context) error {
	activity.GetLogger(ctx).Info("Hello World")
	return nil
}
