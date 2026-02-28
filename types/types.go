// Package types defines type for the program
package types

import (
	"context"
	"time"
)

type Flag struct {
	Flag string
	Func func(ctx context.Context)(error)
}

type Language struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	File string `json:"file"`
}

type Profile struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	Ranking int `json:"ranking"`
}

type Problem struct {
	Question Question `json:"question"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Question struct {
	ID int `json:"id"`
	Title string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
	IsPaidOnly bool `json:"isPaidOnly"`
	Content string `json:"content"`
	CodeSnippet []CodeSnippet `json:"codeSnippets"`
}

type CodeSnippet struct {
	ID int `json:"id"`
	Code string `json:"code"`
	QuestionID int `json:"questionId"`
	LangID int `json:"langId"`
}

type DailyEnvelope struct {
	Active Problem `json:"activeDailyCodingChallengeQuestion"`
}

type TotalQuestions struct {
	Count int `json:"totalQuestions"`
}
