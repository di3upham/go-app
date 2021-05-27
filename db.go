package main

import (
	"context"
	"database/sql"
	"time"

	pb "git.local/go-app/model"
	"github.com/dgraph-io/ristretto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jackc/pgx/v4"
)

type PostgreSQLDB struct {
	conn *pgx.Conn
}

func NewPostgreSQLDB() *PostgreSQLDB {
	conn, err := pgx.Connect(context.Background(), "postgres://username:password@localhost:5432/dbname")
	if err != nil {
		panic(err)
	}
	return &PostgreSQLDB{conn: conn}
}

func (db *PostgreSQLDB) ReadOrder(id string) (*pb.Order, error) {
	var status, productUrl string
	var createdAt int64
	err := db.conn.QueryRow(context.Background(), `SELECT status, created_at, product_url FROM orders WHERE id=$1`, id).Scan(&status, &createdAt, &productUrl)
	if err != nil && err != pgx.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pb.Order{Id: id, Status: status, CreatedAt: createdAt, ProductUrl: productUrl}, nil
}

type DB struct {
	cache *ristretto.Cache
	sql   *sql.DB
}

func NewDB() *DB {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e3, // number of keys to track frequency of (1k).
		MaxCost:     1e8, // maximum cost of cache (100MB).
		BufferItems: 64,  // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	sqldb, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	sqldb.SetConnMaxLifetime(time.Minute * 3)
	sqldb.SetMaxOpenConns(10)
	sqldb.SetMaxIdleConns(10)

	return &DB{cache: cache, sql: sqldb}
}

func (db *DB) UpsertOrder(order *pb.Order) error {
	_, err := db.sql.Exec(`
			INSERT INTO orders(id, status, created_at, product_url)
			VALUES(?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE status=VALUES(status), created_at=VALUES(created_at), product_url=VALUES(product_url)
		`, order.GetId(), order.GetStatus(), order.GetCreatedAt(), order.GetProductUrl())

	if err != nil {
		return err
	}
	if _, found := db.cache.Get(order.GetId()); found {
		db.cache.Set(order.GetId(), order, 1024)
	}
	return nil
}

func (db *DB) ReadOrder(id string) (*pb.Order, error) {
	if value, found := db.cache.Get(id); found {
		return value.(*pb.Order), nil
	}
	var status, productUrl string
	var createdAt int64
	err := db.sql.QueryRow(`SELECT status, created_at, product_url FROM orders WHERE id=? LIMIT 1`, id).Scan(&status, &createdAt, &productUrl)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	order := &pb.Order{Id: id, Status: status, CreatedAt: createdAt, ProductUrl: productUrl}
	db.cache.Set(order.GetId(), order, 1024)
	return order, nil
}

func (db *DB) DeleteOrder(id string) error {
	_, err := db.sql.Exec(`DELETE FROM orders WHERE id=?`, id)
	return err
}
