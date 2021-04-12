package main

import (
	pb "git.local/go-app/model"
	"github.com/dgraph-io/ristretto"
)

type DB struct {
	cache *ristretto.Cache
}

func NewDB() *DB {
	db := &DB{}

	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e3, // number of keys to track frequency of (1k).
		MaxCost:     1e8, // maximum cost of cache (100MB).
		BufferItems: 64,  // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}
	db.cache = cache

	// TODO new persistence db

	return db
}

func (db *DB) UpsertOrder(order *pb.Order) error {
	// TODO replace by persistence db
	db.cache.Set(order.GetId(), order, 1024)

	if _, found := db.cache.Get(order.GetId()); found {
		db.cache.Set(order.GetId(), order, 1024)
	}
	return nil
}

func (db *DB) ReadOrder(id string) (*pb.Order, error) {
	if value, found := db.cache.Get(id); found {
		return value.(*pb.Order), nil
	}
	// TODO read from persistence db
	return nil, nil
}
