package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/sevlyar/go-daemon"
)

func main() {
	// Main arguments parser
	parser := argparse.NewParser("commands", "Available sosimple commands")

	// Init command
	initCmd := parser.NewCommand("init", "Initialize sosimple node listener")
	//initCmdPort := initCmd.String("p", "port", &argparse.Options{Required: false, Help: "Port to allocate"})

	// Daemon command
	daemonCmd := parser.NewCommand("daemon", "Start daemon process")

	// Shutdown command
	shutdownCmd := parser.NewCommand("shutdown", "Shutdown sosimple node listener")

	// Ping command
	pingCmd := parser.NewCommand("ping", "Ping connected Node")
	pingCmdNode := pingCmd.String("n", "node", &argparse.Options{Required: true, Help: "Node to ping"})

	// Send command
	sendCmd := parser.NewCommand("send", "Send a message")
	sendCmdMessage := sendCmd.String("m", "message", &argparse.Options{Required: true, Help: "Message to send"})

	// Log command
	logCmd := parser.NewCommand("log", "View messages")
	//logCmdMessageOnly := logCmd.Flag("", "message-only", &argparse.Options{Required: false, Help: "Show only messages"})

	// List command
	listCmd := parser.NewCommand("list", "List all your friends")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	if initCmd.Happened() {
		initialize()
		log.Print("Node successfully initialized")
	} else if daemonCmd.Happened() {
		log.Print("Starting daemon process")
		cntxt := &daemon.Context{
			PidFileName: "sample.pid",
			PidFilePerm: 0644,
			LogFileName: "sample.log",
			LogFilePerm: 0640,
			WorkDir:     "./", // TODO $HOME directory
			Umask:       027,
			Args:        []string{"[go-daemon sample]"},
		}

		d, err := cntxt.Reborn()
		if err != nil {
			log.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()

		log.Print("- - - - - - - - - - - - - - -")
		log.Print("daemon started")

		launchServer() // TODO add optional port number
	} else if shutdownCmd.Happened() {
		log.Print("Shutting down")
		// TODO
		log.Print("not yet implemented")
	} else if pingCmd.Happened() {
		for i := 0; i < 6; i++ {
			ping(pingCmdNode)
		}
	} else if sendCmd.Happened() {
		send(sendCmdMessage)
	} else if logCmd.Happened() {
		log.Print("not yet implemented")
	} else if listCmd.Happened() {
		list()
	}
}