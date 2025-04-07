package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	InsertUser(ctx context.Context, telegramID, channelID int64) error
}

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) InsertUser(ctx context.Context, telegramID, channelID int64) error {
	query := `INSERT INTO users (telegram_id, channel_id) VALUES ($1, $2)
              ON CONFLICT (telegram_id) DO NOTHING;`
	_, err := r.db.Exec(ctx, query, telegramID, channelID)
	return err
}
