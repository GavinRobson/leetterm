// Package app handles the user input
package app

import (
	"context"
	"leet-term/types"
)

func HandleFlags(flags []types.Flag, args []string, ctx context.Context) error {
	for _, flag := range flags {
		if args[1] == flag.Flag {
			err := flag.Func(ctx)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
