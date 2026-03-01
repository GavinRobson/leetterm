package app

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"leet-term/api"
	"leet-term/appdata"
	"leet-term/log"
	"os"
)

func Config(ctx context.Context) error {
	getCmd := flag.NewFlagSet("config", flag.ExitOnError)
	_ = getCmd.Parse(os.Args[2:])

	args := getCmd.Args()

	appDir, err := appdata.AppDir()
	if err != nil {
		return err
	}

	if len(args) <= 1 {
		appDir, err := appdata.AppDir()
		if err != nil {
			return err
		}
		cfg, found, err := appdata.LoadConfig(appDir)
		if err != nil {
			return err
		}
		if !found {
			fmt.Println("Error: config file not found")
			return errors.New("config not found")
		}

		lang, err := api.GetLanguageByID(ctx, cfg.PreferredLang)
		if err != nil {
			return err
		}

		fmt.Print("Config:\n")
		fmt.Print(log.Struct(cfg) + "\n\n")
		fmt.Println("Default Language: " + lang.Name)
		return nil
	}

	switch args[0] {
	case "set-lang":
		appdata.SaveLang(appDir, args[1])
	default:
		fmt.Printf("unknown command: %s", args[0])
	}
	return nil
}
