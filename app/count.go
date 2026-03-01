package app

import (
	"context"
	"fmt"
	"leet-term/api"
)
func Count(ctx context.Context) error {
	count, err := api.GetCount(ctx)
	if err != nil {
		return err
	}

	fmt.Println(count)
	return nil
}

