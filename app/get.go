package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"leet-term/api"
	"leet-term/appdata"
	"os"
)

func Get(ctx context.Context) error {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	appDir, err := appdata.AppDir()
	if err != nil {
		return err
	}

	cfg, found, err := appdata.LoadConfig(appDir)
	if err != nil {
		return err
	}
	if !found {
		return errors.New("not found")
	}	

	prefLang := cfg.PreferredLang
	if prefLang <= 0 {
		prefLang = 6
	}

	fmt.Println(prefLang)
	lang := getCmd.Int("lang", prefLang, "Preferred language")
	_ = getCmd.Parse(os.Args[2:])

	fmt.Println(prefLang)
	args := getCmd.Args()
	if len(args) < 1 {
		fmt.Println("Error: missing problem id or titleSlug")
		fmt.Println("Usage: leetterm get <problem> (--lang <language>)")
		os.Exit(1)
	}

	q, err := api.GetQuestionByID(args[0], *lang, ctx)
	if err != nil {
		fmt.Println("error getting question by id")
		fmt.Println(err)
	}

	appdata.SaveQuestion(ctx, cfg.DefaultWorkspace, q, *lang)
	appdata.SaveDirection(ctx, cfg.DefaultWorkspace, q)

	fmt.Printf("Success!\n%s saved to %s", q.Title, cfg.DefaultWorkspace)
	return nil
}
