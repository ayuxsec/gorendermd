package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

func runCommand(cmdName string, args []string, timeout int) error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(timeout)*time.Second,
	)
	defer cancel()
	cmd := exec.CommandContext(ctx, cmdName, args...)
	err := cmd.Run()

	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("command %s timed out after %ds", cmd.String(), timeout)
	}

	if err != nil {
		return fmt.Errorf("command %s failed: %v", cmd.String(), err)
	}

	return nil
}
