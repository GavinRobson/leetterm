// Package api defines api services
package api

import (
	"context"
	"encoding/json"
	"fmt"
	"leet-term/supabase"
	"leet-term/types"
	"net/http"
)

const ALFA_URL = "https://alfa-leetcode-api.onrender.com"

var client = &http.Client{}

func GetDailyProblem() (*types.Question, error) {
	url := fmt.Sprintf(ALFA_URL+"/daily")

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var question types.Question

	err = json.NewDecoder(resp.Body).Decode(&question)
	if err != nil {
		return nil, err
	}

	return &question, nil
}

func GetProfileFull(username string) (*types.Profile, error) {
	url := fmt.Sprintf(ALFA_URL+"/%s", username)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var profile types.Profile

	err = json.NewDecoder(resp.Body).Decode(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func GetQuestionByID(id string, lang int, ctx context.Context) (*types.Question, error) {
	q, err := supabase.Supabase.Question.Find(
		ctx, 
		supabase.Eq("id", id),
		supabase.Eq("CodeSnippet.langId", lang),
		supabase.Select("*", "codeSnippets:CodeSnippet(*)"),
	)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func GetLanguages(ctx context.Context) ([]types.Language, error) {
	l, err := supabase.Supabase.Language.FindAll(ctx, supabase.Select("*"))
	if err != nil {
		return nil, err
	}

	return l, nil
}

func GetCount(ctx context.Context) (int, error) {
	count, err := supabase.Supabase.Question.Count(ctx)
	if err != nil {
		return 0, nil
	}
	return count, nil
}
