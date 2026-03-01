package main

import (
	"context"
	"fmt"
	"leet-term/app"
	"leet-term/config"
	"leet-term/supabase"
	"leet-term/types"
	"os"
)

func main() {
	config.LoadEnv()
	ctx := context.Background()

	supabase.LoadClient()

	args := os.Args
	if len(args) < 2 {
		printUsage()
		return
	}

	flags := []types.Flag{
		{
			Flag: "daily",
			Func: app.Daily,
		},
		{
			Flag: "get",
			Func: app.Get,
		},
		{
			Flag: "list",
			Func: app.List,
		},
		{
			Flag: "config",
			Func: app.Config,
		},
		{
			Flag: "count",
			Func: app.Count,
		},
		{
			Flag: "test",
			Func: app.Test,
		},
		{
			Flag: "rand",
			Func: app.Rand,
		},
	}
	app.HandleFlags(flags, args, ctx)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  leetterm daily [--lang <language>]")
	fmt.Println("  leetterm get <problemnumber|titleSlug> [--lang <language>]")
}
