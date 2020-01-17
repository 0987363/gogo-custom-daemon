package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/0987363/mgo"
	"github.com/0987363/mgo/bson"

	//	"github.com/0987363/timestamp"
	"github.com/gogo/protobuf/proto"
	//	"github.com/golang/protobuf/ptypes"

	cu "github.com/0987363/gogo-custom-daemon"
	gt "github.com/0987363/gogotype"
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
	fmt.Println("init id:", id)
	now := time.Now()
	d := cu.Device{
		ID:        cu.NewObjectId(id),
		Timestamp: &now,
		DD:        &gt.NullInt64{sql.NullInt64{12121, true}},
	}
	t := &Test{Timestamp: &now}
	fmt.Printf("tt: %+v\n", t)
	fmt.Printf("dd: %+v\n", &d)

	fmt.Printf("tt2: %+v\n", *t)
	fmt.Printf("dd2: %+v\n", d)

	fmt.Printf("tt timestamp: %+v\n", t.Timestamp)
	fmt.Printf("dd timestamp: %+v\n", d.Timestamp)

	data, err := proto.Marshal(&d)
	if err != nil {
		log.Fatal("marshal failed: ", err)
	}
	fmt.Printf("proto data: %x\n", data)

	d2 := cu.Device{}
	if err := proto.Unmarshal(data, &d2); err != nil {
		log.Fatal("unmarshal failed: ", err)
	}
	d2.This()
	fmt.Printf("unmarshal time: %+v\n", d2.Timestamp)
	fmt.Printf("unmarshal data: %+v\n", d2)
	fmt.Printf("unmarshal point: %+v\n", &d2)
	//	fmt.Printf("unmarshal data: %+v, userid: %+v\n", d2, *d2.UserID)

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
