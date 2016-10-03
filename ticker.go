package udsfs

import (
	"sync/atomic"
	"time"

	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type Ticker struct {
	interval time.Duration
	stopped  int32
}

func (t *Ticker) Process(ctx *core.Context, tuple *core.Tuple, w core.Writer) error {
	var i int64
	for ; atomic.LoadInt32(&t.stopped) == 0; i++ {
		newTuple := core.NewTuple(data.Map{"tick": data.Int(i)})
		if err := w.Write(ctx, newTuple); err != nil {
			return err
		}
		time.Sleep(t.interval)
	}
	return nil
}

func (t *Ticker) Terminate(ctx *core.Context) error {
	atomic.StoreInt32(&t.stopped, 1)
	return nil
}

func CreateTicker(decl udf.UDSFDeclarer, i data.Value) (udf.UDSF, error) {
	interval, err := data.ToDuration(i)
	if err != nil {
		return nil, err
	}
	return &Ticker{
		interval: interval,
	}, nil
}
