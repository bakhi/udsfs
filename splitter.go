package udsfs

import (
	"fmt"
	"strings"

	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
	"gopkg.in/sensorbee/sensorbee.v0/core"
	"gopkg.in/sensorbee/sensorbee.v0/data"
)

type WordSplitter struct {
	field string
}

func (w *WordSplitter) Process(ctx *core.Context, t *core.Tuple, writer core.Writer) error {
	var kwd []string
	if v, ok := t.Data[w.field]; !ok {
		return fmt.Errorf("the tuple doesn't have the required field: %v", w.field)
	} else if s, err := data.AsString(v); err != nil {
		return fmt.Errorf("'%v' field must be string: %v", w.field, err)
	} else {
		kwd = strings.Split(s, " ")
	}

	for _, k := range kwd {
		out := t.Copy()
		out.Data[w.field] = data.String(k)
		if err := writer.Write(ctx, out); err != nil {
			return err
		}
	}
	return nil
}

func (w *WordSplitter) Terminate(ctx *core.Context) error {
	return nil
}

func CreateWordSplitter(decl udf.UDSFDeclarer, inputStream, field string) (udf.UDSF, error) {
	if err := decl.Input(inputStream, nil); err != nil {
		return nil, err
	}
	return &WordSplitter{
		field: field,
	}, nil
}
