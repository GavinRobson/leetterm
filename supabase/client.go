package supabase

import (
	"net/http"
	"os"
)

type Client struct {
	URL string
	Key string
	HTTP *http.Client

	Question *QuestionTable
	Language *LanguageTable
}

var Supabase *Client

func LoadClient() {
	c := &Client{
		URL: os.Getenv("SUPABASE_URL"),
		Key: os.Getenv("SUPABASE_KEY"),
		HTTP: &http.Client{},
	}
	c.Question = &QuestionTable{c: c, name: "Question"}
	c.Language = &LanguageTable{c: c, name: "Lang"}
	Supabase = c
}
