package tracker

import (
	"bufio"
	"fmt"
	"os"
)

type Input interface {
	Get() string
}

type Output interface {
	Out(text string)
}

type ConsoleInput struct{}

func (c ConsoleInput) Get() string {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		// block IO
	}
	return scanner.Text()
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}
