package windows

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/shirou/gopsutil/v4/process"
)

func EnsureLHMRunning() {
	if isLHMRunning() {
		return
	}
	startLHM()
}

func isLHMRunning() bool {
	ps, _ := process.Processes()
	for _, proc := range ps {
		name, _ := proc.Name()
		if name == "LibreHardwareMonitor.exe" {
			return true
		}
	}
	return false
}

func startLHM() {
	exePath, _ := os.Executable()
	baseDir := filepath.Dir(exePath)

	lhm := filepath.Join(baseDir, "tools", "LibreHardwareMonitor", "LibreHardwareMonitor.exe")

	cmd := exec.Command(lhm)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	err := cmd.Start()
	if err != nil {
		return
	}
}
