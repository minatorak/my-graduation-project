package controller

import (
	"flag"
	"fmt"
	"log"
	"os"
)


func Run() {

	printchainCmd := flag.NewFlagSet("printchain", flag.ExitOnError) //blockchain_go printchain
	switch os.Args[1] {
	case "printchain":
		err := printchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		PrintHelp()
		os.Exit(1)
	}

	if printchainCmd.Parsed() {
		fmt.Print(`Call printChain success`)
	}
}

func PrintHelp() {
	fmt.Println("Usage: ")
}
