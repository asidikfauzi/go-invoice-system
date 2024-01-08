package test

import (
	"go-invoice-system/common/database"
	"go-invoice-system/common/helper"
	"os"
	"testing"
)

func TestInitDatabase(t *testing.T) {
	envVars := map[string]string{
		"DB_USERNAME": "root",
		"DB_PASSWORD": "",
		"DB_HOST":     "localhost",
		"DB_PORT":     "3306",
		"DB_DATABASE": "invoice_system",
	}

	originalEnv := saveOriginalEnvValues(envVars)
	defer restoreOriginalEnvValues(envVars, originalEnv)

	db, err := database.InitDatabase()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	if db == nil {
		t.Fatal("Expected non-nil database instance, got nil")
	}

	var result int
	db.Raw("SELECT 1").Scan(&result)

	if result != 1 {
		t.Fatal("Expected query result to be 1, got:", result)
	}
}

func saveOriginalEnvValues(envVars map[string]string) map[string]string {
	originalEnv := make(map[string]string)
	for key, value := range envVars {
		originalEnv[key] = helper.GetEnv(key)
		SetEnv(key, value)
	}
	return originalEnv
}

func restoreOriginalEnvValues(envVars map[string]string, originalEnv map[string]string) {
	for key, value := range originalEnv {
		SetEnv(key, value)
	}
}

func SetEnv(key, value string) {
	os.Setenv(key, value)
}
