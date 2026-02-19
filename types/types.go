// Package types defines type for the program
package types

import "time"

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
	ID string `json:"questionId"`
	Title string `json:"title"`
	TitleSlug string `json:"titleSlug"`
	Difficulty string `json:"difficulty"`
	IsPaidOnly bool `json:"isPaidOnly"`
	ContentHTML string `json:"content"`
	CodeSnippets []CodeSnippet `json:"codeSnippets"`
}

type CodeSnippet struct {
	Lang string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code string `json:"code"`
}

type DailyEnvelope struct {
	Active Problem `json:"activeDailyCodingChallengeQuestion"`
}
