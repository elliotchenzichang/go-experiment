package timer

import (
	"context"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
	"time"
)

func SampleTimerWorkflow(ctx workflow.Context) error {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	childCtx, _ := workflow.WithCancel(ctx)
	selector := workflow.NewSelector(ctx)

	timerFuture := workflow.NewTimer(childCtx, 10*time.Second)
	selector.AddFuture(timerFuture, func(f workflow.Future) {
		workflow.ExecuteActivity(ctx, SayHelloWorldActivity).Get(ctx, nil)
	})

	// wait the timer or the order processing to finish
	selector.Select(ctx)
	return nil
}

func SayHelloWorldActivity(ctx context.Context) error {
	activity.GetLogger(ctx).Info("Hello World")
	return nil
}
