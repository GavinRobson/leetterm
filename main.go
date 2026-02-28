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

	var flags []types.Flag
	flags = append(flags, types.Flag{
		Flag: "daily",
		Func: app.Daily,
	})
	flags = append(flags, types.Flag{
		Flag: "get",
		Func: app.Get,
	})
	flags = append(flags, types.Flag{
		Flag: "config",
		Func: app.Config,
	})
	flags = append(flags, types.Flag{
		Flag: "count",
		Func: app.Count,
	})
	flags = append(flags, types.Flag{
		Flag: "test",
		Func: app.Test,
	})
	flags = append(flags, types.Flag{
		Flag: "rand",
		Func: app.Rand,
	})

	app.HandleFlags(flags, args, ctx)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  leetterm daily [--lang <language>]")
	fmt.Println("  leetterm get <problemnumber|titleSlug> [--lang <language>]")
}
