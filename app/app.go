package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"leet-term/api"
	"leet-term/appdata"
	"leet-term/log"
	"leet-term/types"
	"math/rand"
	"os"
	"strconv"
)

func HandleFlags(flags []types.Flag, args []string, ctx context.Context) error {
	for _, flag := range flags {
		if args[1] == flag.Flag {
			err := flag.Func(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Daily(ctx context.Context) error {
	fmt.Println("Getting daily problem...")
	return nil
}

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

	lang := getCmd.Int("lang", prefLang, "Preferred language")
	_ = getCmd.Parse(os.Args[2:])

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
	fmt.Println(log.Struct(q))
	return nil
}


func Config(ctx context.Context) error {
	getCmd := flag.NewFlagSet("config", flag.ExitOnError)
	_ = getCmd.Parse(os.Args[2:])

	args := getCmd.Args()

	appDir, err := appdata.AppDir()
	if err != nil {
		return err
	}

	if len(args) <= 1 {
		fmt.Println("not enough arguments")
		return fmt.Errorf("not enough arguments")
	}

	switch args[0] {
	case "set-lang":
		appdata.SaveLang(appDir, args[1])
	default:
		fmt.Printf("unknown command: %s", args[0])
	}
	return nil
}

func Count(ctx context.Context) error {
	count, err := api.GetCount(ctx)
	if err != nil {
		return err
	}

	fmt.Println(count)
	return nil
}

func Rand(ctx context.Context) error {
	prefLang, err := appdata.LoadLang()
	if err != nil {
		return err
	}
	count, err := api.GetCount(ctx)
	if err != nil {
		return nil
	}

	randID := rand.Intn(count - 1) + 1
	q, err := api.GetQuestionByID(strconv.Itoa(randID), prefLang, ctx)
	if err != nil {
		return err
	}

	if q == nil {
		err := Rand(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println(log.Struct(q))
	return nil
}
