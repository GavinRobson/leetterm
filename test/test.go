package test

import (
	"context"
	"fmt"
	"leet-term/config"
	"leet-term/supabase"
	"os"
)

func test() {
	config.LoadEnv()
	ctx := context.Background()
	sb := supabase.New(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	q, err := sb.Question.Find(ctx, supabase.Eq("titleSlug", "add-two-numbers"))
	if err != nil {
		panic(err)
	}

	fmt.Println(q)
}
