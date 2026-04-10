package main

import (
	"fmt"
	"os"
	"log"
	"context"
	"github.com/urfave/cli/v3"
)

func main() {
	// описываем нашу cli команду и логику ее работы
	// и берем указатель на эту команду
	cmd := &cli.Command{
		// поля name и usage используются для вывода -h
		// данный флаг в явном виде не объявлен, но 
		// библиотека позволяет автоматически его обрабатывать
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(context.Context, *cli.Command) error {
			// непосредственная логика выполнения прграммы (основная логика)
			var message string = "Hello from Hexlet!"
			fmt.Println(message)
			return nil
		},
	}

	// разыменовываем указатель и обращаемся к методу Run 
	// т.е. фактическое выполнение логики описанной выше происходит
	// сейчас внутри условной конструкции if
	// мы получаем ошибку от выполнения команды, и если она есть - выводим ее
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
	
}

