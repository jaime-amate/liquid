package tags

import (
	"fmt"
	"io"

	"github.com/osteele/liquid/expressions"
	"github.com/osteele/liquid/render"
)

func incrementTag(source string) (func(io.Writer, render.Context) error, error) {
	_, err := expressions.Parse(source)
	if err != nil {
		return nil, err
	}

	return func(w io.Writer, ctx render.Context) error {
		value := 0
		if v := ctx.Get(source); v != nil {
			if intValue, ok := v.(int); ok {
				value = intValue
			}
		}

		ctx.Set(source, value+1)

		_, err = fmt.Fprint(w, value)
		return err
	}, nil
}