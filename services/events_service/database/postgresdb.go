package database

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/piyushsharma67/events_booking/services/events_service/models"
	"github.com/piyushsharma67/events_booking/services/events_service/postgresdb"
	"github.com/piyushsharma67/events_booking/services/events_service/utils"
)

type Sqldb struct {
	querries *postgresdb.Queries
}

func NewSqldb(querries *postgresdb.Queries) Database {
	return &Sqldb{
		querries: querries,
	}
}

/* Initialising the DB */

func InitPostgres() (*pgxpool.Pool, *postgresdb.Queries) {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=" + sslmode
	slog.Info("Connecting to Postgres", "dsn", dsn)

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Invalid DSN: %v", err)
	}
	// Wait until Postgres is ready
	// THIS is where you must retry
	for i := 0; i < 15; i++ {
		fmt.Printf("Ping attempt %d...\n", i+1)

		// Use a strict timeout for each ping attempt
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		err = pool.Ping(ctx)
		cancel()

		if err == nil {
			break
		}

		fmt.Printf("Database not ready: %v. Retrying in 2s...\n", err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("❌ Database never became ready")
	}

	// Run schema.sql
	content, err := os.ReadFile("postgresdb/schema.sql")
	if err != nil {
		panic(err)
	}

	// Split statements by semicolon and execute individually
	statements := splitSQLStatements(string(content))
	for _, stmt := range statements {
		if strings.TrimSpace(stmt) == "" {
			continue
		}
		if _, err := pool.Exec(context.Background(), stmt); err != nil {
			slog.Error("failed to execute statement:", "err", err, "stmt", stmt)
			panic(err)
		}
	}

	queries := postgresdb.New(pool)
	return pool, queries
}

// splitSQLStatements splits SQL by semicolon and ignores semicolons inside dollar-quoted blocks
func splitSQLStatements(sql string) []string {
	var stmts []string
	scanner := bufio.NewScanner(strings.NewReader(sql))
	scanner.Split(bufio.ScanLines)
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), "--") {
			continue // skip comments
		}
		sb.WriteString(line + "\n")
		if strings.HasSuffix(strings.TrimSpace(line), ";") {
			stmts = append(stmts, sb.String())
			sb.Reset()
		}
	}
	if sb.Len() > 0 {
		stmts = append(stmts, sb.String())
	}
	return stmts
}

func (s *Sqldb) GenerateEvent(ctx context.Context, event *models.Event) (*models.Event, error) {
	startTime, err := utils.GetPgTime(event.StartTime)

	if err != nil {
		return nil, err
	}
	endTime, err := utils.GetPgTime(event.EndTime)

	if err != nil {
		return nil, err
	}
	dbArg := postgresdb.CreateEventParams{
		OrganizerID: 1,
		Title:       event.Title,
		Description: utils.ToText(event.Description),
		Location:    event.Location,
		ImageUrl:    utils.ToText(event.ImageURL),
		StartTime:   startTime,
		EndTime:     endTime,
		Status:      utils.OPEN,
	}
	genRow, err := s.querries.CreateEvent(ctx, dbArg)

	if err != nil {
		return nil, err
	}

	return &models.Event{
		Description: genRow.Description.String,
		Title:       event.Title,
		ImageURL:    event.ImageURL,
		Location:    event.Location,
		StartTime:   event.StartTime,
		ID:          event.ID,
		EndTime:     event.EndTime,
		Rows:        event.Rows,
		Timestamps:  event.Timestamps,
	}, nil

}

func (s *Sqldb) DeleteEvent(ctx context.Context, eventId any) error {
	return nil
}
func (s *Sqldb) UpdateEvent(ctx context.Context, event *models.Event) (models.Event, error) {
	return models.Event{}, nil
}
func (s *Sqldb) GetEvent(ctx context.Context, eventId any) (models.Event, error) {
	return models.Event{}, nil
}
