package main

import (
	"bytes"
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
)

func main() {
	f := jen.NewFile("main")
	f.Func().Id("main").Params().Block(
		jen.Qual("fmt", "Println").Call(jen.Lit("Hello, world")),
	)
	buf := &bytes.Buffer{}

	save(buf, f)
}

func save(buf *bytes.Buffer, f *jen.File) {
	err := f.Render(buf)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(buf.String())
		if err := os.WriteFile("./jen_test/test/generated.go", buf.Bytes(), 0644); err != nil {
			panic(err)
		}
	}
}
