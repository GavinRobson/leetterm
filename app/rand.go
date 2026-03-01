package app

import (
	"context"
	"fmt"
	"leet-term/api"
	"leet-term/appdata"
	"leet-term/log"
	"math/rand"
	"strconv"
)

func Rand(ctx context.Context) error {
	prefLang, err := appdata.LoadLang()
	if err != nil {
		return err
	}
	count, err := api.GetCount(ctx)
	if err != nil {
		return nil
	}

	randID := rand.Intn(count - 1) + 1
	q, err := api.GetQuestionByID(strconv.Itoa(randID), prefLang, ctx)
	if err != nil {
		return err
	}

	if q == nil {
		err := Rand(ctx)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println(log.Struct(q))
	return nil
}
