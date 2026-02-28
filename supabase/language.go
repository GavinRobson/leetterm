package supabase

import (
	"context"
	"leet-term/types"
)

type LanguageTable struct {
	c *Client
	name string
}

func (t *LanguageTable) FindAll(ctx context.Context, opts ...QueryOpt) ([]types.Language, error) {
	var out []types.Language
	err := t.c.selectJSON(ctx, t.name, &out, append(opts)...)
	return out, err
}


