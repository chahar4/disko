package channel

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...any) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...any) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) AddChannel(ctx context.Context, channel *Channel) (*Channel, error) {
	query := "INSERT INTO channels(name , guild_id) VALUES ($1 ,$2) returning id"

	var lastInsertChannelId int
	err := r.db.QueryRowContext(ctx, query, channel.Name, channel.GuildID).Scan(&lastInsertChannelId)
	if err != nil {
		return nil, err
	}
	channel.ID = lastInsertChannelId
	return channel, nil
}

func (r *repository) AddMessage(ctx context.Context, message *Message) (*Message, error) {
	query := "INSERT INTO message (author_id,channel_id,content) VALUES ($1,$2,$3)"
	var lastMessageInserted int
	err := r.db.QueryRowContext(ctx, query).Scan(&lastMessageInserted)
	if err != nil {
		return nil, err
	}
	message.ID = lastMessageInserted
	return message, nil
}
