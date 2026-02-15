package daemon

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
)

type Daemon struct {
	pidFile string
}

func NewDaemon(workDir string) *Daemon {
	pidFile := filepath.Join(workDir, ".gitsentry", "gitsentry.pid")
	return &Daemon{pidFile: pidFile}
}

func (d *Daemon) WritePID() error {
	pid := os.Getpid()
	pidStr := strconv.Itoa(pid)
	
	return os.WriteFile(d.pidFile, []byte(pidStr), 0644)
}

func (d *Daemon) ReadPID() (int, error) {
	data, err := os.ReadFile(d.pidFile)
	if err != nil {
		return 0, err
	}
	
	return strconv.Atoi(string(data))
}

func (d *Daemon) IsRunning() bool {
	pid, err := d.ReadPID()
	if err != nil {
		return false
	}
	
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

func (d *Daemon) Stop() error {
	pid, err := d.ReadPID()
	if err != nil {
		return err
	}
	
	process, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	
	if err := process.Signal(os.Interrupt); err != nil {
		return err
	}
	
	return d.RemovePID()
}

func (d *Daemon) RemovePID() error {
	if _, err := os.Stat(d.pidFile); os.IsNotExist(err) {
		return nil
	}
	
	return os.Remove(d.pidFile)
}

func (d *Daemon) Daemonize() error {
	if d.IsRunning() {
		return fmt.Errorf("daemon already running")
	}
	
	return d.WritePID()
}