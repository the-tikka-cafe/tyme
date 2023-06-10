package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Conf struct {
	Host        string `json:"host"`
	Port        uint16 `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	DB          string `json:"database"`
	MaxActive   int32  `json:"max_active"`
	MaxIdle     int    `json:"max_idle"`
	IdleTimeout int    `json:"idle_timeout"`
	SSLMode     string `json:"ssl_mode"`
}

type Postgres struct {
	conf Conf
	pool *pgxpool.Pool
}

func New(c Conf) *Postgres {

	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", c.Username, c.Password, c.Host, c.Port, c.DB, c.SSLMode)

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		fmt.Println("Error creating connection string")
		return nil
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil
	}
	return &Postgres{
		conf: c,
		pool: pool,
	}
}

func (p *Postgres) GetConn() (*pgxpool.Conn, error) {
	c := context.Background()
	conn, err := p.pool.Acquire(c)
	if err != nil {
		return nil, err
	}

	return conn, nil
}