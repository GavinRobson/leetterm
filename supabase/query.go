package supabase

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var ErrNotFound = fmt.Errorf("not found")

type Query struct {
	params url.Values
}

type QueryOpt func(*Query)

func (c *Client) selectJSON(ctx context.Context, table string, dest any, opts ...QueryOpt) error {
	q := &Query{params: url.Values{}}
	for _, opt := range opts {
		opt(q)
	}

	u := fmt.Sprintf("%s/rest/v1/%s", c.URL, table)
	if enc := q.params.Encode(); enc != "" {
		u += "?" + enc
	}

	req, err := http.NewRequestWithContext(ctx, "GET", u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("apikey", c.Key)
	req.Header.Set("Authorization", "Bearer "+c.Key)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var b []byte
		b, _ = io.ReadAll(resp.Body)
		return fmt.Errorf("supabase %s: %s: %s", table, resp.Status, string(b))
	}

	return json.NewDecoder(resp.Body).Decode(dest)
}

