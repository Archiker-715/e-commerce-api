package pg

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var DB *gorm.DB

func Connect() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect to DB: %v\n", err)
	}

	if err := migrate(DB); err != nil {
		log.Fatal("migration err:", err)
	}
	log.Println("Database connected successfully")
	return DB
}

func migrate(DB *gorm.DB) error {
	for _, sqlFilePath := range findSqlFiles() {
		content, err := os.ReadFile(sqlFilePath)
		if err != nil {
			return fmt.Errorf("err when read file %v: %w", sqlFilePath, err)
		}
		if err := DB.Exec(string(content)).Error; err != nil {
			return fmt.Errorf("SQL execution err: %w", err)
		}
	}
	return nil
}

func findSqlFiles() (filePaths []string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("get wd err", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(cwd, "go.mod")); err == nil {
			break
		}
		parent := filepath.Dir(cwd)
		if parent == cwd {
			// Дошли до корня файловой системы, но не нашли
			log.Fatal("project root not found")
		}
		cwd = parent
	}

	mgPath := filepath.Join(cwd, "internal", "repo", "pg", "migrations")

	files, err := os.ReadDir(mgPath)
	if err != nil {
		log.Fatalf("err read files in folder %v: %v\n", mgPath, err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fp := filepath.Join(mgPath, file.Name())
			filePaths = append(filePaths, fp)
		}
	}
	return
}
