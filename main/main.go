package main

import (
	"fmt"
	"log"
	"time"

	"github.com/0987363/mgo"
	"github.com/0987363/mgo/bson"
//	"github.com/0987363/timestamp"
	"github.com/gogo/protobuf/proto"
//	"github.com/golang/protobuf/ptypes"

	cu ".."
)

type Test struct {
	ID        bson.ObjectId `json:"id,omitempty" bson:"id,omitempty"`
	B         int
	Timestamp *time.Time `bson:"timestamp,omitempty"`
}

func main() {
	db, err := mgo.Dial("192.168.88.16")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	c := db.DB("test").C("test2")

	id := bson.NewObjectId()
	fmt.Println("init:", id)
	now := time.Now()
	d := cu.Device{
		ID:        &id,
		Timestamp: &now,
	}

	data, err := proto.Marshal(&d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("proto data: %x\n", data)

	d2 := cu.Device{}
	if err := proto.Unmarshal(data, &d2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("unmarshal data: %+v\n", d2)

	if err := c.Insert(&d); err != nil {
		log.Fatal("err:", err)
	}
//	fmt.Printf("insert db data:%+v\n", dd)

	d3 := Test{}
	if err := c.Find(bson.M{"id": id}).One(&d3); err != nil {
		fmt.Printf("err:%+v\n", err)
	}
	fmt.Printf("from db data:%+v\n", d3)
}
