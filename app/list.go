package app

import (
	"context"
	"flag"
	"fmt"
	"leet-term/api"
	"leet-term/supabase"
	"os"
)

func List(ctx context.Context) error {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	if err := listCmd.Parse(os.Args[2:]); err != nil {
		return err
	}

	args := listCmd.Args()

	if len(args) == 0 {
		fmt.Println("Error: not enough arguments")
		fmt.Println("Usage: leetterm list <lang, questions, cmd>")
		os.Exit(1)
	}

	switch args[0] {
	case "lang":
		langs, err := api.GetLanguages(ctx)
		if err != nil {
			return err
		}
		fmt.Println("Available Languages:")
		for _, l := range langs {
			fmt.Println(l.Name)
		}
	case "question":
		questions, err := supabase.Supabase.Question.FindMany(ctx)
		if err != nil {
			return err
		}

		for _, q := range questions {
			fmt.Printf("ID: %d Name: %s\n", q.ID, q.TitleSlug)
		}
	default:
		fmt.Println("unknown command")
	}

	return nil
}
