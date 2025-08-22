package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

// InitConfig initializes the configuration by reading from config/config.yaml file.
func LoadConfig() error {
	// godotenv.Load() по умолчанию ищет файл ".env" в текущей директории
	// и загружает его переменные в окружение.
	err := godotenv.Load()
	if err != nil {
		// Оборачиваем ошибку для предоставления контекста
		return fmt.Errorf("ошибка загрузки .env файла: %w", err)
	}

	// Функция успешно отработала, ошибки нет
	return nil
}
