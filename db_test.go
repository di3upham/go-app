package main

import (
	"testing"
	"time"

	pb "git.local/go-app/model"
	"google.golang.org/protobuf/proto"
)

func TestDB(t *testing.T) {
	if true {
		return // disable
	}
	db := NewDB()
	var err error
	inOrder := &pb.Order{Id: "order1", Status: "open", CreatedAt: time.Now().UnixNano() / 1e6, ProductUrl: "https://sampleapp.local/1.png"}
	err = db.UpsertOrder(inOrder)
	if err != nil {
		t.Error(err)
		return
	}

	exoOrder := inOrder
	acoOrder, err := db.ReadOrder("order1")
	if err != nil {
		t.Error(err)
		return
	}
	if !proto.Equal(acoOrder, exoOrder) {
		t.Errorf("want %#v, actual %#v", exoOrder, acoOrder)
	}

	err = db.DeleteOrder("order1")
	if err != nil {
		t.Error(err)
		return
	}
}
