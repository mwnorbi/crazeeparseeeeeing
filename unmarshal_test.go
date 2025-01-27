package main

import (
	"testing"
)

var stringJson = "{\"messageType\":9,\"message\":{\"name\":\"Example\",\"potato\":42}}"
var bArray = []byte(stringJson)

var shortMessage = "9ExampleDeli42"
var bArrayShortMessage = []byte(shortMessage)

func BenchmarkUsingByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		usingByte(bArray)
	}
}

func BenchmarkUsingJsonUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalStuff(bArray)
	}
}

func BenchmarkUsingByteOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		onlyByte(bArray)
	}
}
