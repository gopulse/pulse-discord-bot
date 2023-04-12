package helpers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func compileAndRun(code string) (string, error) {
	f, err := os.CreateTemp("", "goworkshop*.go")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Printf("failed to remove temporary file: %v", err)
		}
	}(f.Name())

	_, err = f.WriteString(code)
	if err != nil {
		return "", fmt.Errorf("failed to write to temporary file: %w", err)
	}

	cmd := exec.Command("go", "run", f.Name())
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to compile and run code: %w", err)
	}

	return out.String(), nil
}
