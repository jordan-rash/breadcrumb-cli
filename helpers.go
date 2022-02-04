package breadcrumb

import (
	"os"
	"os/exec"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func sanitizeInput(s string) (string, []string) {
	cmd := strings.Split(strings.TrimSpace(s), " ")[:1][0]
	args := strings.Split(strings.TrimSpace(s), " ")[1:]

	log.Debug(len(s))
	log.Debug(strings.Split(strings.TrimSpace(s), " ")[:1][0])

	return cmd, args
}

func ClearScreen() {
	value := runtime.GOOS
	var cmd *exec.Cmd
	switch value {
	case "unix", "darwin":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cls")
	default:
		log.Error("Terminal not supported")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
