package connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ChangeRow(conn *pgx.Conn, ctx context.Context) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ('Хочу пиццу', 'Ну приколХАХАХАХАХ', FALSE, '2025-11-10 18:00:05');
	`

	_, err := conn.Exec(ctx, sqlQuery)

	if err != nil {
		return err
	}

	return nil
}
