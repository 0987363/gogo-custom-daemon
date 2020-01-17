package custombytesnonstruct

import (
	"bytes"
	"encoding/gob"
	fmt "fmt"
)

type UserID int

func (c *UserID) Size() int {
	return 8
}

func (c *UserID) Unmarshal(data []byte) error {
	network := bytes.NewBuffer(data)
	dec := gob.NewDecoder(network)
	err := dec.Decode(c)
	if err != nil {
		return err
	}

	return nil
}

func (c UserID) Marshal() ([]byte, error) {
	fmt.Println("m", c)
	//	return []byte(bson.UserID(c).Hex()), nil

	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	err := enc.Encode(c)
	if err != nil {
		return nil, err
	}
	return network.Bytes(), nil
}

/*
func (c *ObjectId) Size() int {
	return len(*c)
}
*/
