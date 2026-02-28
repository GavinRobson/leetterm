package appdata

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"leet-term/api"
	"leet-term/log"
	"leet-term/types"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Version int `json:"version"`
	Username string `json:"username"`
	PreferredLang int `json:"preferred_lang"`
	DefaultWorkspace string `json:"default_workspace"`
}

type State struct {
	Version int `json:"version"`
	Completed map[string]bool `json:"completed"`
	Meta map[string]string `json:"meta,omitempty"`
}

const AppName = "leet-term"

func AppDir() (string, error) {
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(base, AppName), nil
}

func configPath(appDir string) string { return filepath.Join(appDir, "config.json") }
func statePath(appDir string) string { return filepath.Join(appDir, "state.json") }

func EnsureAppDir() (string, error) {
	appDir, err := AppDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(appDir, 0o755); err != nil {
		return "", err
	}
	return appDir, nil
}

func LoadConfig(appDir string) (*Config, bool, error) {
	p := configPath(appDir)
	b, err := os.ReadFile(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, false, nil
		}
		return nil, false, err
	}
	var c Config
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, true, fmt.Errorf("config corrupt: %w", err)
	}
	return &c, true, nil
}

func LoadLang() (int, error) {
	appDir, err := AppDir()
	if err != nil {
		return -1, err
	}

	cfg, found, err := LoadConfig(appDir)
	if err != nil {
		return -1, err
	}
	if !found {
		return -1, errors.New("config not found")
	}

	return cfg.PreferredLang, nil
}

func SaveConfig(appDir string, c *Config) error {
	c.Version = 1
	p := configPath(appDir)

	tmp := p + ".tmp"
	b, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}

func SaveLang(appDir string, lang string) error {
	ctx := context.Background()
	p := configPath(appDir)
	langs, err := api.GetLanguages(ctx)
	if err != nil {
		return err
	}

	var prefLang int

	for _, l := range langs {
		fmt.Println(log.Struct(l))
		if lang == l.Name || lang == l.Slug {
			prefLang = l.ID
			break
		} 
	}

	cfg, found, err := LoadConfig(appDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			SaveConfig(appDir, &Config{PreferredLang: prefLang})
		}
		return err
	}

	if !found {
		SaveConfig(appDir, &Config{PreferredLang: prefLang})
	}

	cfg.PreferredLang = prefLang

	fmt.Println(cfg)

	tmp := p + ".tmp"
	b, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}

func ValidateLang(c *Config) error {
	if c.PreferredLang <= 0 {
		return errors.New("missing preferred_lang")
	}
	return nil
}

func ValidateConfig(c *Config) error {
	if c.Username == "" {
		return errors.New("missing username")
	}
	if c.PreferredLang <= 0 {
		return errors.New("missing preferred_lang")
	}
	if c.DefaultWorkspace == "" {
		return errors.New("missing default_workspace")
	}
	return nil
}

func LoadOrCreateState(appDir string) (*State, error) {
	p := statePath(appDir)
	b, err := os.ReadFile(p)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			s := &State{Version: 1, Completed: map[string]bool{}}
			return s, SaveState(appDir, s)
		}
	}
	var s State
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("state corrupt: %w", err)
	}
	if s.Completed == nil {
		s.Completed = map[string]bool{}
	}
	return &s, nil
}

func SaveState(appDir string, s *State) error {
	s.Version = 1
	p := statePath(appDir)

	tmp := p + ".tmp"
	b, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(tmp, b, 0o600); err != nil {
		return err
	}
	return os.Rename(tmp, p)
}

func SaveProblem(saveDir string, p *types.Problem, lang string) error {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	path := filepath.Join(saveDir, p.Question.TitleSlug, lang)

	if err := os.MkdirAll(path, 0o755); err != nil {
		return err
	}

	file := "problem." + lang
	filename := filepath.Join(path, file)

	b, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filename, b, 0644); err != nil {
		return err
	}

	return nil
}
