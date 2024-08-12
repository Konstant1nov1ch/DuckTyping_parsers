package main

import (
	"example_parsers/parsers"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Задаем директорию с файлами для парсинга
	dir := "./fsm"

	// Создаем список парсеров
	parsersList := []parsers.FileParser{
		parsers.ConfigParser{},
		parsers.UserDataParser{},
		parsers.NotifyParser{},
	}

	for _, parser := range parsersList {
		fmt.Println("инициализация парсера - ", parser.GetId())
	}

	// Читаем файлы из директории и парсим их
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			for _, parser := range parsersList {
				//Тот самый Parse
				err := parser.Parse(path)
				if err != nil {
					//fmt.Printf("Ошибка обработки файла %s через парсер %T: %v\n", path, parser, err)
				} else {
					fmt.Printf("Файл %s успешно обработан через парсер %T\n", path, parser)
					break // Выходим из цикла если файл успешно обработан
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}
}
