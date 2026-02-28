package supabase

import (
	"fmt"
	"strings"
)

func Eq(col string, val any) QueryOpt {
	return func(q *Query) {
		q.params.Add(col, fmt.Sprintf("eq.%v", val))
	}
}

func Limit(n int) QueryOpt {
	return func(q *Query) {
		q.params.Set("limit", fmt.Sprintf("%d", n))
	}
}

func Order(col string, asc bool) QueryOpt {
	return func(q *Query) {
		dir := "desc"
		if asc {
			dir = "asc"
		}
		q.params.Set("order", fmt.Sprintf("%s.%s", col, dir))
	}
}

func Select(cols ...string) QueryOpt {
	return func(q * Query) {
		q.params.Set("select", strings.Join(cols, ","))
	}
}

