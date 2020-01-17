package custombytesnonstruct

import fmt "fmt"

/*
type Device struct {
	ID        bson.ObjectId `protobuf:"bytes,1,opt,name=ID,proto3,customtype=github.com/0987363/mgo/bson.ObjectId" json:"id,omitempty" bson:"_id,omitempty"`
	Timestamp *time.Time    `protobuf:"bytes,10,opt,name=UpdatedAt,proto3,stdtime" json:"updated_at,omitempty" bson:"updated_at,omitempty"`

	UserID        *ObjectId `protobuf:"bytes,13,opt,name=ObjectId,proto3,customtype=ObjectId" json:"user_id,omitempty" bson:"user_id,omitempty"`
}
*/

func (c *Device) This() {
	fmt.Printf("This : %+v\n", *c)
}

/*
func (c *Device) Size() int {
	return 9
}
*/

/*
func (c *Device) Unmarshal(data []byte) error {

//		id := Device(bson.DeviceHex(string(data)))
//		*c = id
//		fmt.Println("un", c)
//		return nil

	network := bytes.NewBuffer(data)
	dec := gob.NewDecoder(network)
	err := dec.Decode(c)
	if err != nil {
		return err
	}

	return nil
}

func (c Device) Marshal() ([]byte, error) {
	fmt.Println("m", c)
	//	return []byte(bson.Device(c).Hex()), nil

	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	err := enc.Encode(c)
	if err != nil {
		return nil, err
	}
	return network.Bytes(), nil
}
*/

/*
func (c *ObjectId) Size() int {
	return len(*c)
}
*/
