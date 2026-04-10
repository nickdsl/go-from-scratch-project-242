package code

import (
	"fmt"
	"os"
)

func GetPathSize(path string) (string, error) {
	mySize := int64(0)

	fileInfo, err := os.Lstat(path)
	if err != nil {
		return fmt.Sprintf("%dB", mySize), err
	}
	// если мы тут, то нам получилось без ошибок получить какую то информацию
	// которая находится в fileInfo
	
	// нам нужно проверить каталог ли это

	// если Не каталог, то просто возвращаем размер файла
	if !(fileInfo.IsDir()) {
		mySize += fileInfo.Size()
		return fmt.Sprintf("%dB", mySize), nil
	}
	// если мы доходим до этого места, то мы в ситуации, когда нам на вход подали не файл, а каталог

	// для начала нужно прочитать каталог и получить его содержимое	
	dirEntries, err := os.ReadDir(path)

	// возникли ошибки при чтении каталога
	if err != nil {
		return fmt.Sprintf("%dB", mySize), err
	}

	// если мы тут, ошибок нет, работаем с полученными данными
	for _, file := range dirEntries {
		// если файл в каталоге не директория
		if !(file.IsDir()) {
			// если это не каталог, то получаем информацию о файле
			fileInfo, err := file.Info()
			// обрабатываем ошибку получения информации о файле
			if err != nil {
                                return fmt.Sprintf("%dB", mySize), err
                        }
			// кажется здесь нужно реализовать в будущем обработку отдельно для скрытых файлов	
			mySize += fileInfo.Size()
		}else{
			// в этой ветке кода мы оказываемся если у нас текущий файл, это вложенный каталог
			// по-умолчанию мы его игнорируем пока что
			continue
		}
	}
	return fmt.Sprintf("%dB", mySize), nil
}
