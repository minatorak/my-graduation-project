package main

import (
	"github.com/minatorak/block-container/src/controller"
	"os"
)


func main() {
	if len(os.Args) < 2 {
		controller.PrintHelp()
		os.Exit(1)
	}
	controller.Run()
}


func run() {
	//fmt.Printf("running %v\n", os.Args[2:])
	//	//
	//	//cmd := exec.Command(os.Args[2], os.Args[3:]...)
	//	//cmd.Stdin = os.Stdin
	//	//cmd.Stdout = os.Stdout
	//	//cmd.Stderr = os.Stderr
	//	//
	//	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	//	//Cloneflags: syscall.CLONE_NEWUTS,
	//	//}
	//	//
	//	//must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}