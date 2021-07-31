package main

import (
	"testing"
)

func TestHello(t *testing.T)  {
	
	//test for valid argument
	result := hello("Hejsan") //should return "Hello Hejsan!"

	if result != "Hello Hejsan!" {
		t.Errorf("hello(\"Hejsan\") failed, expected %v, got %v", "Hello duuuude!", result)
	} else {
		t.Logf("hello(\"Hejsan\") success, expected %v, got %v", "Hello duuuude!", result)
	}
}
