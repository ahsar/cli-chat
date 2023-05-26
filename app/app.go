package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {

	ch := make(chan string)

	go readCmd(ch)

	for str := range ch {
		delCmd(str)
	}
}

func readCmd(ch chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		fmt.Println("run in---")
		ch <- name
	}
}

func delCmd(str string) {
	fmt.Println("your enter is:", str)
}
