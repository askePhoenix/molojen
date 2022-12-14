package main

import (
	"bytes"
	"fmt"
	"github.com/dave/jennifer/jen"
	"log"
	"os"
)

func main() {
	f := jen.NewFile("stream")

	f.Type().Id("IntArray").Index().Int()
	f.Type().Id("FuncIntArray").Func().Params(jen.Int(), jen.Int()).Int()
	f.Type().Id("FuncIntArrayConsumer").Func().Params(jen.Int(), jen.Int())

	f.
		Func().Params(jen.Id("i").
		Id("IntArray")).Id("Map").
		Params(jen.Id("f").Id("FuncIntArray")).
		Id("IntArray").Block(
		jen.Return(
			jen.Qual("github.com/samber/lo", "Map").
				Types(jen.Int(), jen.Int()).Call(jen.Id("i"), jen.Id("f")),
		),
	)

	f.
		Func().Params(jen.Id("i").
		Id("IntArray")).Id("ForEach").
		Params(jen.Id("f").Id("FuncIntArrayConsumer")).
		Block(
			jen.Qual("github.com/samber/lo", "ForEach").
				Types(jen.Int()).Call(jen.Id("i"), jen.Id("f")),
		)

	f.
		Func().Params(jen.Id("i").
		Id("IntArray")).Id("ToStringArray").
		Params().
		Id("StringArray").Block(
		jen.Return(
			jen.Qual("github.com/samber/lo", "Map").
				Types(jen.Int(), jen.String()).Call(jen.Id("i"),
				jen.Qual("github.com/askePhoenix/molojen/lo_test/fn", "IntToString")),
		),
	)

	f.Type().Id("StringArray").Index().String()
	f.Type().Id("FuncStringArray").Func().Params(jen.String(), jen.Int()).String()
	f.Type().Id("FuncStringArrayConsumer").Func().Params(jen.String(), jen.Int())

	f.
		Func().Params(jen.Id("s").
		Id("StringArray")).Id("Map").
		Params(jen.Id("f").Id("FuncStringArray")).
		Id("StringArray").Block(
		jen.Return(
			jen.Qual("github.com/samber/lo", "Map").
				Types(jen.String(), jen.String()).Call(jen.Id("s"), jen.Id("f")),
		),
	)

	f.
		Func().Params(jen.Id("s").
		Id("StringArray")).Id("ForEach").
		Params(jen.Id("f").Id("FuncStringArrayConsumer")).
		Block(
			jen.Qual("github.com/samber/lo", "ForEach").
				Types(jen.Int()).Call(jen.Id("s"), jen.Id("f")),
		)

	f.
		Func().Params(jen.Id("s").
		Id("StringArray")).Id("ToStringArray").
		Params().
		Id("IntArray").Block(
		jen.Return(
			jen.Qual("github.com/samber/lo", "Map").
				Types(jen.String(), jen.Int()).Call(jen.Id("s"),
				jen.Qual("github.com/askePhoenix/molojen/lo_test/fn", "StringToInt")),
		),
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
			log.Panicln(err)
		}
	}
}
