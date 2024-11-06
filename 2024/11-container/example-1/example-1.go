package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// go run main.go run <cmd> <args>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("only 'run' is supported.")
	}
}

func getPID() {
	pid := os.Getpid()
	parentpid := os.Getppid()
	fmt.Printf("The parent process id of %v is %v\n", pid, parentpid)
}

func run() {
	fmt.Printf("Running %v as process %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags: syscall.CLONE_NEWUTS,
		// Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
