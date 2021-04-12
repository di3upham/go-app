package main

import (
	"strconv"
	"sync/atomic"
	"time"

	pb "git.local/go-app/model"
)

type OrderDB interface {
	UpsertOrder(order *pb.Order) error
	ReadOrder(id string) (*pb.Order, error)
}

type Ordermgr struct {
	orderIndex int64
	db         OrderDB
}

func NewOrdermgr(db OrderDB) *Ordermgr {
	return &Ordermgr{db: db}
}

func (mgr *Ordermgr) CreateOrder(order *pb.Order) (*pb.Order, error) {
	if order == nil {
		return nil, nil
	}
	if order.GetProductUrl() != "" {
		if err := ValidUrl(order.GetProductUrl()); err != nil {
			return nil, err
		}
	}
	order.Id = strconv.Itoa(int(atomic.AddInt64(&mgr.orderIndex, 1)))
	order.CreatedAt = time.Now().UnixNano() / 1e6
	err := mgr.db.UpsertOrder(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (mgr *Ordermgr) DeleteOrder(id string) error {
	// TODO
	return nil
}

func (mgr *Ordermgr) ListOrders(id string) (*pb.Orders, error) {
	// TODO
	return nil, nil
}

func (mgr *Ordermgr) UpdateOrder(*pb.Order) (*pb.Order, error) {
	// TODO
	return nil, nil
}

func (mgr *Ordermgr) ReadOrder(id string) (*pb.Order, error) {
	// TODO
	if id == "" {
		return nil, nil
	}
	return mgr.db.ReadOrder(id)
}
