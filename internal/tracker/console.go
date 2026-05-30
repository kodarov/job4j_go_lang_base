package tracker

import (
	"bufio"
	"fmt"
	"os"
)

type ConsoleInput struct{}

func (c ConsoleInput) Get() string {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	return scanner.Text()
}

type ConsoleOutput struct{}

func (c ConsoleOutput) Out(text string) {
	fmt.Println(text)
}
