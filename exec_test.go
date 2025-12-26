package main

import "testing"

func TestRunCommand(t *testing.T) {
	if err := runCommand("firefox", []string{"http://localhost:33689"}, 10); err != nil {
		t.Log(err.Error())
	}
}
