// Package initflow defines the initialization of the project
package initflow

import (
	"bufio"
	"fmt"
	"leet-term/appdata"
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

func pickPreferredLang(options []string) (string, error) {
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

		 for i, opt := range options {
			if i == selected {
				fmt.Printf("> %s\n", opt)
			} else {
				fmt.Printf("  %s\n", opt)
			}
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
			return options[selected], nil
		case 3:
			return "", fmt.Errorf("init cancelled")
		}
	}
}

func RunInit(appDir string) (*appdata.Config, error) {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("LeetCode username: ")
	username, _ := in.ReadString('\n')
	username = strings.TrimSpace(username)

	options := []string{"Javascript", "Go", "Rust"}
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
		Username: username,
		PreferredLang: lang,
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
