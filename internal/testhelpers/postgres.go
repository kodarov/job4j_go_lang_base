package testhelpers

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewTestPool(ctx context.Context) (*pgxpool.Pool, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "6543")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "tracker")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.New: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pool.Ping: %w", err)
	}

	return pool, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Migrate(ctx context.Context, pool *pgxpool.Pool, migrationsDir string) error {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		return fmt.Errorf("os.ReadDir: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".sql") {
			content, err := os.ReadFile(filepath.Join(migrationsDir, file.Name()))
			if err != nil {
				return fmt.Errorf("os.ReadFile: %w", err)
			}

			// Goose-aware splitting: only take the "Up" part
			sqlContent := string(content)
			if idx := strings.Index(sqlContent, "-- +goose Down"); idx != -1 {
				sqlContent = sqlContent[:idx]
			}

			// Very basic SQL comment stripping to avoid pgx issues with some characters
			lines := strings.Split(sqlContent, "\n")
			var filtered []string
			for _, line := range lines {
				trimmed := strings.TrimSpace(line)
				if trimmed == "" || strings.HasPrefix(trimmed, "--") {
					continue
				}
				filtered = append(filtered, line)
			}
			sql := strings.Join(filtered, "\n")

			if strings.TrimSpace(sql) == "" {
				continue
			}

			_, err = pool.Exec(ctx, sql)
			if err != nil {
				if strings.Contains(err.Error(), "already exists") {
					continue
				}
				return fmt.Errorf("pool.Exec (%s): %w", file.Name(), err)
			}
		}
	}
	return nil
}

func CleanDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "TRUNCATE TABLE items RESTART IDENTITY CASCADE")
	return err
}
