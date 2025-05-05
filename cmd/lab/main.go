package main
//fdfsfsdfsd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countWords(s string) int {
	words := strings.Fields(s)
	return len(words)
}

func main() {
	var inputStr string
	var scanner *bufio.Scanner

	// Проверяем, есть ли файл с входными данными
	if _, err := os.Stat("/app/input.txt"); err == nil {
		// Чтение данных из файла
		file, err := os.Open("/app/input.txt")
		if err != nil {
			fmt.Println("Ошибка открытия файла:", err)
			return
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)
		if scanner.Scan() {
			inputStr = scanner.Text()
		}
	} else {
		// Интерактивный ввод
		fmt.Println("Используется стандартный ввод")
		scanner = bufio.NewScanner(os.Stdin)
		fmt.Print("Введите строку> ")
		if scanner.Scan() {
			inputStr = scanner.Text()
		}
	}

	wordCount := countWords(inputStr)
	fmt.Printf("Количество слов в строке: %d\n", wordCount)
}
