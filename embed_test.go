package learn_go_embed

import (
	"embed"
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

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestEmbedMultipleFile(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(a)
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(b)
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(c)
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			contentFile, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(contentFile))
		}
	}
}
