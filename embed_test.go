package learn_go_embed

import (
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

//go:embed hehe1.png
var hehe1 []byte

func TestByteReader(t *testing.T) {
	err := ioutil.WriteFile("hehe1_new.png", hehe1, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
