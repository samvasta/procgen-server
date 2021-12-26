package generators

import (
	"errors"
	"image"
	"math/rand"
	"strings"

	g "github.com/fogleman/gg"
)

type Generator interface {
	Generate(g g.Context, rand *rand.Rand)
}

func Draw(generatorType string, width, height int, seed int64) (image.Image, error) {
	generator, err := getGenerator(generatorType)
	if err != nil {
		return nil, err
	}

	drawContext := g.NewContext(width, height)

	rand := rand.New(rand.NewSource(seed))

	generator.Generate(*drawContext, rand)

	return drawContext.Image(), nil
}

func getGenerator(generatorType string) (Generator, error) {
	switch strings.ToLower(generatorType) {
	case "test":
		return TestGenerator{}, nil
	}

	return nil, errors.New("Generator does not exist")
}
