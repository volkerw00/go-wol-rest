package main

import (
	"bufio"
	"os"

	flags "github.com/jessevdk/go-flags"
)

var options struct {
	BroadcastIP string `short:"b" long:"broadcastIp" description:"The default IPv4 broadcast adress to advertise the magic packet to" required:"true"`
}

func parseArgs() error {
	parser := flags.NewParser(&options, flags.Default)
	_, error := parser.Parse()
	if error != nil {
		f := bufio.NewWriter(os.Stdout)
		defer f.Flush()
		parser.WriteHelp(f)
		return error
	}
	return nil
}
