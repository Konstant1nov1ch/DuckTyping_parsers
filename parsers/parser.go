package parsers

// FileParser - интерфейс для парсинга файлов
type FileParser interface {
	Parse(filePath string) error
	GetId() string
}
