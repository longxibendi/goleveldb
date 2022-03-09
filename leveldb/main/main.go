package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/comparer"
	"github.com/syndtr/goleveldb/leveldb/memdb"
)

func main() {
	db := memdb.New(comparer.DefaultComparer, 1)
	fmt.Println("data=", string(db.GetkvData()))
	fmt.Println("len(data)=", len(db.GetkvData()))
	fmt.Println("nodeData", db.GetnodeData())
	fmt.Println("len(nodeData)", len(db.GetnodeData()))

	db.Put([]byte("c"), []byte("C"))
	fmt.Println(db)

	fmt.Println("data=", string(db.GetkvData()))
	fmt.Println("len(data)=", len(db.GetkvData()))
	fmt.Println("nodeData", db.GetnodeData())
	fmt.Println("len(nodeData)", len(db.GetnodeData()))

	db.Put([]byte("a"), []byte("A"))
	fmt.Println(db)
	fmt.Println("data=", string(db.GetkvData()))
	fmt.Println("len(data)=", len(db.GetkvData()))
	fmt.Println("nodeData", db.GetnodeData())
	fmt.Println("len(nodeData)", len(db.GetnodeData()))

	db.Put([]byte("b"), []byte("b"))
	fmt.Println(db)
	fmt.Println("data=", string(db.GetkvData()))
	fmt.Println("len(data)=", len(db.GetkvData()))
	fmt.Println("nodeData", db.GetnodeData())
	fmt.Println("len(nodeData)", len(db.GetnodeData()))

}
