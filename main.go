package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"leet-term/api"
	"leet-term/appdata"
)

// appDir, err := appdata.EnsureAppDir()
// if err != nil {
// 	fmt.Print("error: EnsureAppDir")
// 	return
// }
// cfg, found, err := appdata.LoadConfig(appDir)
// if err != nil {
// 	fmt.Print("error: LoadConfig")
// 	return
// }
//
// if !found || appdata.ValidateConfig(cfg) != nil {
// 	cfg, err = initflow.RunInit(appDir)
// 	if err != nil {
// 		fmt.Print("error: RunInit")
// 		return
// 	}
// }
//
// state, err := appdata.LoadOrCreateState(appDir)
// if err != nil {
// 	fmt.Print("error: LoadOrCreateState")
// 	return
// }
//
// username := cfg.Username
// profile, err := api.GetProfileFull(username)
// if err != nil {
// 	log.Fatal(err)
// }
//

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "daily":
		dailyCmd := flag.NewFlagSet("daily", flag.ExitOnError)
		lang := dailyCmd.String("lang", "cpp", "Preferred language")
		_ = dailyCmd.Parse(os.Args[2:])

		handleDaily(*lang)

	case "get":
		getCmd := flag.NewFlagSet("get", flag.ExitOnError)
		lang := getCmd.String("lang", "cpp", "Preferred language")
		_ = getCmd.Parse(os.Args[2:])

		args := getCmd.Args()
		if len(args) < 1 {
			fmt.Println("Error: missing problem id or titleSlug")
			fmt.Println("Usage: leetterm get <problem> --lang <language>")
			os.Exit(1)
		}

		handleGet(args[0], *lang)

	default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
	}
}

func handleDaily(lang string) {
	p, err := api.GetDailyProblem(lang)
	if err != nil {
		fmt.Println("error getting random problem: %w", err)
		return
	}

	in := bufio.NewReader(os.Stdin)
	fmt.Print("Save to? [.] ")
	input, _ := in.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "" {
		input = "."
	}

	saveDir := input
	if input == "." {
		saveDir, err = os.Getwd()
		if err != nil {
			fmt.Println("could not get current directory:", err)
			return
		}
	} 

	if err := appdata.SaveProblem(saveDir, p, lang); err != nil {
		fmt.Println("save failed:", err)
		return
	}

	fmt.Printf("Saved to [%s]", saveDir)
}

func handleGet(problem string, lang string) {
	fmt.Println("Fetching problem", problem)
	fmt.Println("Language:", lang)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  leetterm daily [--lang <language>]")
	fmt.Println("  leetterm get <problemnumber|titleSlug> [--lang <language>]")
}
