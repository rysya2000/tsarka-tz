package usecase

import (
	"bufio"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Find_SubStr(str string) string {
	res := ""

	for i := 0; i < len(str); i++ {
		for j := i; j < len(str); j++ {
			if IsUnique(str[i : j+1]) {
				if len(str[i:j+1]) > len(res) {
					res = str[i : j+1]
				}
			}
		}
	}
	return res
}

func IsUnique(str string) bool {
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(str); j++ {
			if i != j {
				if str[i] == str[j] {
					return false
				}
			}
		}
	}

	return true
}

func Check_Email(str string) []string {
	re := regexp.MustCompile(`Email:\s*[\w\d._%+-]+@[\w\d.-]+\.[\w]{2,}`)

	return re.FindAllString(str, -1)
}

func Find_Self(str string) ([]string, error) {
	slice := []string{}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Пропускаем директории и файлы, которые не являются файлами на Go
		if info.IsDir() || filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Открываем файл для чтения
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Считываем содержимое файла построчно
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			// Ищем все идентификаторы в строке, которые содержат подстроку
			re := regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`)
			identifiers := re.FindAllString(line, -1)
			for _, identifier := range identifiers {
				if strings.Contains(identifier, str) {
					slice = append(slice, identifier)
				}
			}
		}

		return nil
	})

	return slice, err
}
