package custombytesnonstruct

import (
	fmt "fmt"

	"github.com/0987363/mgo/bson"
)

type ObjectId bson.ObjectId

func NewObjectId(id bson.ObjectId) *ObjectId {
	i := ObjectId(id)
	return &i
}

func (c *ObjectId) Size() int {
	return 24
}

func (c ObjectId) String() string {
	return fmt.Sprintf(`ObjectIdHex("%x")`, string(c))
}

func (c *ObjectId) Unmarshal(data []byte) error {
	id := ObjectId(bson.ObjectIdHex(string(data)))
	*c = id
	fmt.Println("un: ", *c)
	return nil
}

func (c ObjectId) Marshal() ([]byte, error) {
	fmt.Println("hex: ", bson.ObjectId(c).Hex())
	return []byte(bson.ObjectId(c).Hex()), nil
}

/*
func (c *ObjectId) Size() int {
	return len(*c)
}
*/
