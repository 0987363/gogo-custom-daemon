package main

import (
	"fmt"
	"log"

	//	"github.com/0987363/timestamp"
	"github.com/gogo/protobuf/proto"
	//	"github.com/golang/protobuf/ptypes"

	cu "github.com/0987363/gogotype"
)

type Test struct {
	B cu.NullInt64
}

func main() {
	d := Test{}
	fmt.Printf("dd2: %+v\n", d)

	data, err := proto.Marshal(&d)
	if err != nil {
		log.Fatal("marshal failed: ", err)
	}
	fmt.Printf("proto data: %x\n", data)

	d2 := Test{}
	if err := proto.Unmarshal(data, &d2); err != nil {
		log.Fatal("unmarshal failed: ", err)
	}
	fmt.Printf("unmarshal data: %+v\n", d2)
}
