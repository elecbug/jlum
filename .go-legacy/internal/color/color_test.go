package color_test

import (
	"testing"

	"github.com/elecbug/jlum/internal/color"
)

func TestColor(t *testing.T) {
	c := color.Color{
		Red:   255,
		Green: 125,
		Blue:  1,
	}

	t.Logf("%+v", c)
	t.Logf("%s", c.ToHex())
	t.Logf("%+v", color.FromHex(c.ToHex()))
}
