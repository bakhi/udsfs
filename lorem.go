package udsfs

import (
	"math/rand"
	"strings"
	"time"

	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

/*
var (
	Lorem = strings.Split(strings.Replace('lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor
	incididunt ut labore et dolore magna aliqua Ut enim ad minim veniam quis nostrud exercitation ullamco laboris
	nisi ut aliquip ex ea commodo consequat Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore
	eu fugiat nulla pariatur Excepteur sint occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim
	id est laborum', "\n", " ", -1), " ")
)
*/

var Lorem = strings.Split(strings.Replace("lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Ut enim ad minim veniam quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur Excepteur sint occaecat cupidatat non proident sunt in culpa qui officia deserunt mollit anim id est laborumlorem ipsum dolor sit amet consectetur adipiscing elit", "\n", " ", -1), " ")

type LoremSource struct {
	interval time.Duration
}

func (l *LoremSource) GenerateStream(ctx *core.Context, w core.Writer) error {
	for {
		var text []string
		for l := rand.Intn(5) + 5; l > 0; l-- {
			text = append(text, Lorem[rand.Intn(len(Lorem))])
		}

		t := core.NewTuple(data.Map{
			"text": data.String(strings.Join(text, " ")),
		})
		if err := w.Write(ctx, t); err != nil {
			return err
		}

		time.Sleep(l.interval)
	}
}

func (l *LoremSource) Stop(ctx *core.Context) error {
	return nil
}

func CreateLoremSource(ctx *core.Context, ioParams *bql.IOParams, params data.Map) (core.Source, error) {
	interval := 1 * time.Microsecond
	if v, ok := params["interval"]; ok {
		i, err := data.ToDuration(v)
		if err != nil {
			return nil, err
		}
		interval = i
	}
	return core.ImplementSourceStop(&LoremSource{
		interval: interval,
	}), nil
}
