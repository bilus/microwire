package templates

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func Raw(text string) (t templ.Component) {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, text)
		return err
	})
}
