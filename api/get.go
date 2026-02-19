// Package api defines api services
package api

import (
	"encoding/json"
	"fmt"
	"leet-term/types"
	"net/http"
)

const URL = "https://alfa-leetcode-api.onrender.com"

var client = &http.Client{}

func GetProfileFull(username string) (*types.Profile, error) {
	url := fmt.Sprintf(URL + "/%s", username)	

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

func GetGetProblem() {
	
}

func GetProblem(titleSlug string) (*types.Problem, error) {
	url := fmt.Sprintf(URL + "/select/raw?titleSlug=%s", titleSlug)

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var problem types.Problem

	err = json.NewDecoder(resp.Body).Decode(&problem)
	if err != nil {
		return nil, err
	}

	return &problem, nil
}

func GetDailyProblem(lang string) (*types.Problem, error) {
	url := fmt.Sprintf(URL + "/daily/raw")
	
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	var env types.DailyEnvelope

	err = json.NewDecoder(resp.Body).Decode(&env)
	if err != nil {
		return nil, err
	}

	return &env.Active, nil
}
