package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func usingByte(input []byte) *qq {
	switch input[15] {
	case 56:
		fmt.Println("not the one i am looking for")
	case 57:
		var message qqMessage
		err := json.Unmarshal(input, &message)
		if err != nil {
			fmt.Println("having problem at usingbyte", err)
		}
		return &message.Message
	}

	return nil
}

func normalStuff(input []byte) *qq {
	var mtype baseMessage
	err := json.Unmarshal(input, &mtype)
	if err != nil {
		fmt.Println("having problem at normalstuff", err)
	}

	switch mtype.MessageType {
	case 1:
		fmt.Println("also not the one i am looking for")
	case 9:
		var message qqMessage
		json.Unmarshal(input, &message)
		return &message.Message
	}

	return nil
}

func onlyByte(input []byte) *qq {
	deli := "Deli"
	switch input[0] {
	case 56:
		fmt.Println("not likey")
	case 57:
		deliIndex := getDelimiterIndex(input, deli)
		a := string(input[*deliIndex+len(deli):])
		num, _ := strconv.Atoi(a)

		return &qq{
			Name:   string(input[1:*deliIndex]),
			Potato: int32(num),
		}
	}

	return nil
}
