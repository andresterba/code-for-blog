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
	case "child":
		child()
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
	// getPID()

	fmt.Printf("Running %v as process %d\n", os.Args[2:], os.Getpid())

	// cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// This will show errors on macOS!
	// We run a new process by cloning the current one and attach a new namespace with some attributes to it.
	// NEWUTS = Unix Timesharing System (hostname)
	// NEWPID = Give the process a new PID (0) and 1 for the child
	// NEWNS = Mount namespace for /sys, /dev, etc.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		// Cloneflags: syscall.CLONE_NEWUTS,
		// Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	// getPID()

	fmt.Printf("Running %v as process %d\n", os.Args[2:], os.Getpid())

	syscall.Sethostname([]byte("container"))
	syscall.Chroot("/home/aps/share/gocker/debianfs")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(cmd.Run())
}
