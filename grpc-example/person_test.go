package Person

import (
	"fmt"
	"log"
	"testing"

	"github.com/golang/protobuf/proto"
)

func TestPersonDemo1(t *testing.T) {
	// Create a new Person message
	person := &PersonSimple{
		Name:  "Alice",
		Age:   30,
		Email: []string{"alice@example.com", "alice@gmail.com"},
	}

	// Serialize the Person message to a byte slice
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// Deserialize the byte slice back into a new Person message
	newPerson := &Person{}
	err = proto.Unmarshal(data, newPerson)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	// Print the original and deserialized messages
	fmt.Println("Original Person: ", person)
	fmt.Println("New Person: ", newPerson)
}
