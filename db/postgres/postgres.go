package postgres

import (
	"L0/config"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type PsqlRepo struct {
	pool *pgxpool.Pool
}

func InitPsqlDB(c *config.Config) PsqlRepo {
	ctx := context.Background()
	connectionUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		*c.Postgres.Host,
		*c.Postgres.Port,
		*c.Postgres.User,
		*c.Postgres.Password,
		*c.Postgres.DbName)
	fmt.Println(connectionUrl)
	pool, err := pgxpool.New(ctx, connectionUrl)
	if err != nil {
		log.Fatal(err)
	}
	return PsqlRepo{pool: pool}
}

func (r *PsqlRepo) Create(uid string, body []byte) (err error) {
	ctx := context.Background()
	_, err = r.pool.Exec(ctx, "INSERT INTO notes VALUES ($1, $2) ON CONFLICT DO NOTHING", uid, body)
	if err != nil {
		return err
	}
	return nil
}

func (r *PsqlRepo) CheckIfExist() (pgx.Rows, error) {
	ctx := context.Background()
	rows, err := r.pool.Query(ctx, "SELECT * FROM notes")
	if rows != nil {
		return rows, err
	}
	return rows, errors.New("don't exist")
}

func (r *PsqlRepo) InitTable() error {
	ctx := context.Background()
	_, err := r.pool.Exec(ctx, `CREATE TABLE notes (
    	orderUid TEXT PRIMARY KEY,
    	infoJson json
		);`)
	if err != nil {
		return err
	}
	return nil
}
