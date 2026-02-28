package supabase

import (
	"context"
	"fmt"
	"leet-term/types"
	"net/http"
	"strconv"
	"strings"
)

type QuestionTable struct {
	c *Client
	name string
}

func (t *QuestionTable) Find(ctx context.Context, opts ...QueryOpt) (*types.Question, error) {
	var out []types.Question
	err := t.c.selectJSON(ctx, t.name, &out, append(opts, Limit(1))...)
	if err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, ErrNotFound
	}
	return &out[0], nil
}

func (t *QuestionTable) FindMany(ctx context.Context, opts ...QueryOpt) ([]types.Question, error) {
	var out []types.Question
	err := t.c.selectJSON(ctx, t.name, &out, opts...)
	return out, err
}

func (t *QuestionTable) Count(ctx context.Context) (int, error) {
	url := fmt.Sprintf("%s/rest/v1/%s?select=*", t.c.URL, t.name)

	req, err := http.NewRequestWithContext(ctx, "HEAD", url, nil)
	if err != nil {
		return 0, nil
	}

	req.Header.Set("apikey", t.c.Key)
	req.Header.Set("Authorization", "Bearer "+t.c.Key)
	req.Header.Set("Prefer", "count=exact")

	resp, err := t.c.HTTP.Do(req)
	if err != nil {
		return 0, nil
	}
	defer resp.Body.Close()

	cr := resp.Header.Get("Content-Range")
	if cr == "" {
		return 0, fmt.Errorf("missing Content-Range header")
	}

	parts := strings.Split(cr, "/")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid Content-Range: %s", cr)
	}

	total, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return total, nil
}
