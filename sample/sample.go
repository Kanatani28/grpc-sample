package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	// urfave/cli を使用
	app := cli.NewApp()

	// アプリケーションの基本情報
	app.Name = "sample"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	// Action 実行する処理の内容
	app.Action = func(context *cli.Context) error {

		message := context.Args().Get(0)

		if context.Bool("c") {
			message = message + " meow"
		}
		if context.String("lang") == "spanish" {
			fmt.Println("Hola", message)
		} else {
			fmt.Println("Hello", message)
		}
		return nil
	}

	// Flags でオプションを定義
	// Flagsに[]cli.Flagを登録することで設定できる
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "cat, c, aa",
			Usage: "Echo with cat",
		},
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "Run with debug mode",
		},
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for greeting",
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
