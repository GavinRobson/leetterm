package app

import (
	"context"
	"fmt"
	"leet-term/appdata"
)

type TestInfo struct {
	Name string
	Func func(context.Context)(string, bool)
}

func Test(ctx context.Context) error {
	tests := []TestInfo{
		{
			Name: "Loading Config",
			Func: TestLoadConfig,
		},
	}

	for _, test := range tests {
		result, passed := test.Func(ctx)
		if !passed {
			fmt.Printf("Test %s failed with: %s", test.Name, result)
		}
		fmt.Printf("Test %s passed! ", test.Name)
	}

	return nil
}

func TestLoadConfig(ctx context.Context) (string, bool) {
	appDir, err := appdata.AppDir()
	if err != nil {
		return "", false
	}
	_, found, err := appdata.LoadConfig(appDir)
	if !found || err != nil {
		return "", false
	}

	return "", true
}
