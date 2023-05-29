package app

import (
	"bufio"
	"fmt"
	"os"

	//"strings"

	"flag"

	"github.com/ahsar/cli-chat/internal/chat"
	"github.com/ahsar/cli-chat/internal/logger"
)

// verbose 模式
var v bool

func Init() {
	flag.BoolVar(&v, "v", false, "verbose 模式")
	flag.Parse()

	level := "info"
	if v {
		level = "debug"
	}
	logger.Init(level)
}

func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Text()

		if line == "^D" {
			fmt.Println("Exiting program...")
			break
		}

		extractCmd(line)
		fmt.Print("> ")
	}
}

// extract cmd
func extractCmd(s string) {
	fmt.Println("now receive cmd:", s)

	switch s {
	case "login":
		chat.Run()
	case "friends":
		chat.Friends()
	case "groups":
		chat.Groups()
	}
}
