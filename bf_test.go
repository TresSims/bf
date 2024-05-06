package main

import (
	"strings"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	instructions := "++++++++[>++++[>++>+++>+++>+<<<<-]>+>+>->>+[<]<-]>>.>---.+++++++..+++.>>.<-.<.+++.------.--------.>>+.>++."
	got, err := ParseString(instructions)
	if strings.Compare(got, "Hello World!\n") != 0 || err != nil {
		t.Errorf("TestHelloWorld, = '''\n%s\n'''\n want: '''\nHello World!\n'''", got)
	}
}

func TestHelloWorldFile(t *testing.T) {
	instructions := "./hello-world.bf"
	got, err := ParseFile(instructions)
	if strings.Compare(got, "Hello World!\n") != 0 || err != nil {
		t.Errorf("TestHelloWorldFile, = '''\n%s\n'''\n want: '''\nHello World!\n'''\n", got)
	}
}
