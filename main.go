package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// os.Argsによる引数の取得
	fmt.Println("Hello Golang")
	fmt.Printf("args count %d \n", len(os.Args))
	fmt.Printf("args %#v \n", os.Args)

	for i, v := range os.Args {
		fmt.Printf("args[%d] -> %s\n", i, v)
	}

	// flagによる引数の扱い

	//// パターン１
	// var (
	// 	// パラメータ名、デフォルト値、パラメータの説明
	// 	msg   = flag.String("m", "default message", "Message")
	// 	num   = flag.Uint("n", 0, "Count>=0")
	// 	debug = flag.Bool("debug", false, "true: Debug enabled")
	// )

	// flag.Parse()
	// fmt.Printf("param -m -> %s\n", *msg)
	// fmt.Printf("param -n -> %d\n", *num)
	// fmt.Printf("param -debug -> %t\n", *debug)

	//// パターン２
	var (
		msg2   string
		num2   uint
		debug2 bool
	)

	flag.StringVar(&msg2, "m2", "default message 2", "Message2")
	flag.UintVar(&num2, "n2", 0, "Count >= 0")
	flag.BoolVar(&debug2, "debug2", false, "true: debug2 enabled")

	flag.Parse()

	fmt.Printf("param -m2 -> %s\n", msg2)
	fmt.Printf("param -n2 -> %d\n", num2)
	fmt.Printf("param -debug2 -> %t\n", debug2)
}
