package generators

import (
	"math/rand"

	g "github.com/fogleman/gg"

	common "samvasta.com/procgen/common"
)

type TestGenerator struct {
}

func (generator TestGenerator) Generate(g g.Context, rand *rand.Rand) {

	g.DrawCircle(float64(g.Width())/2, float64(g.Height())/2, 0.4*float64(common.Min(g.Width(), g.Height())))
	g.SetRGB(0, 0, 0)
	g.Fill()
}
