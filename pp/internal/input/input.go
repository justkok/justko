package input

import (
	"bufio"
	"fmt"
	"os"
)

func Input() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку> ")
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}
