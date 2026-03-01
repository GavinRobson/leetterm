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
	err := t.c.selectJSON(ctx, t.name, &out, opts...)
	return out, err
}

func (t *LanguageTable) Find(ctx context.Context, opts ...QueryOpt) (*types.Language, error) {
	var out []types.Language
	err := t.c.selectJSON(ctx, t.name, &out, append(opts, Limit(1))...)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, ErrNotFound
	}
	return &out[0], nil
}


