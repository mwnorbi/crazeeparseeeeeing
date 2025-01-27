package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {

	var shortMessage = "9ExampleDeli42"
	var bArrayShortMessage = []byte(shortMessage)

	qqStruct := qq{
		Name:   "Example",
		Potato: 42,
	}

	fmt.Println(bArrayShortMessage)
	asd := onlyByte(bArrayShortMessage)

	if *asd != qqStruct {
		panic("nope")
	}

	fmt.Println(bArrayShortMessage[1:])
	fmt.Println(bArrayShortMessage[1:4])

	fmt.Println(string([]byte{56, 57, 58, 59}))
	a := bArrayShortMessage[1:4]

	fmt.Println(bytes.Equal([]byte{69, 120, 97}, a))

	mymessage := "{\"type\":\"123\", \"something\":\"aaaaaa\"}"
	b, _ := json.Marshal(mymessage)
	fmt.Println(b)
	arr := []byte{34}
	var str string
	json.Unmarshal(arr, &str)

	fmt.Println(string(arr[:]))

	qqe := qq{
		Name:   "qqstruct",
		Potato: 23,
	}

	b, _ = json.Marshal(qqe)
	fmt.Println(b)

	var qe qq

	json.Unmarshal(b, &qe)
	fmt.Println(qe)
	stringJson := "{\"messageType\":9,\"message\":{\"name\":\"Example\",\"potato\":42}}"
	fmt.Println([]byte(stringJson))
	bs := []byte(stringJson)
	fmt.Println(normalStuff(bs))
	fmt.Println(usingByte(bs))
}
