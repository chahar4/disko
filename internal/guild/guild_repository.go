package guild

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

func (r *repository) AddGuild(ctx context.Context, guild *Guild) (*Guild, error) {
	query := "INSERT INTO guilds(name , owner_id) VALUES ($1 ,$2 ) returning id"

	var lastinsertguildId int
	err := r.db.QueryRowContext(ctx, query , guild.Name , guild.OwenerID).Scan(&lastinsertguildId)
	if err != nil {
		return nil, err
	}
	guild.ID = lastinsertguildId
	return guild, nil
}

func (r *repository) GetAllGuildByUserID(ctx context.Context, userID int) (*[]Guild, error) {
	query := `SELECT g.* FROM guilds g JOIN user_guild ug ON g.id = ug.guild_id WHERE ug.user_id = $1;`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	var guilds []Guild
	for rows.Next() {
		var guild Guild
		if err := rows.Scan(&guild.ID, &guild.Name, &guild.OwenerID); err != nil {
			return nil, err
		}
		guilds = append(guilds, guild)
	}
	return &guilds, nil
}

func (r *repository) AddUserToGuild(ctx context.Context, guildID, userID int) error {
	query := "INSERT INTO user_guild (user_id , guild_id) VALUES ($1 ,$2)"
	_, err := r.db.ExecContext(ctx, query, userID, guildID)
	if err != nil {
		return err
	}
	return nil
}
