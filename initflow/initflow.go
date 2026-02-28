// Package initflow defines the initialization of the project
package initflow

import (
	"bufio"
	"context"
	"fmt"
	"leet-term/appdata"
	"leet-term/supabase"
	"leet-term/types"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/term"
)

func readKey() (byte, error) {
	buf := make([]byte, 1)
	_, err := os.Stdin.Read(buf)
	return buf[0], err
}

func pickPreferredLang(options []types.Language) (string, error) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	selected := 0

	fmt.Print("\033[s")

	for {
		fmt.Print("\033[u")
		fmt.Println("Preferred language:")

		fmt.Print("\033[2K\r")
		for i, opt := range options {
			if i == selected {
				fmt.Printf("> %s\n", opt.Name)
			} else {
				fmt.Printf("  %s\n", opt.Name)
			}
			fmt.Print("\033[2K\r")
		}

		key, _ := readKey()

		switch key {
		case 'j':
			if selected < len(options)-1 {
				selected++
			}
		case 'k':
			if selected > 0 {
				selected--
			}
		case '\r', '\n':
			fmt.Println()
			return options[selected].Name, nil
		case 3:
			return "", fmt.Errorf("init cancelled")
		}
	}
}

func RunInit(appDir string) (*appdata.Config, error) {
	ctx := context.Background()
	in := bufio.NewReader(os.Stdin)

	fmt.Print("LeetCode username: ")
	username, _ := in.ReadString('\n')
	username = strings.TrimSpace(username)

	options, err := supabase.Supabase.Language.FindAll(ctx, supabase.Order("id", true))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	lang, err := pickPreferredLang(options)

	if err != nil {
		return nil, err
	}

	cwd, _ := os.Getwd()
	def := filepath.Join(cwd, "leetcode")

	fmt.Printf("Default workspace [%s]: ", def)
	ws, _ := in.ReadString('\n')
	ws = strings.TrimSpace(ws)

	if ws == "" {
		ws = def
	}

	_ = os.MkdirAll(ws, 0o755)

	cfg := &appdata.Config{
		Username:         username,
		PreferredLang:    lang,
		DefaultWorkspace: ws,
	}

	if err := appdata.ValidateConfig(cfg); err != nil {
		return nil, err
	}
	if err := appdata.SaveConfig(appDir, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
