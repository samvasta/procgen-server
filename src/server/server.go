package server

import (
	"bytes"
	"image/png"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	generators "samvasta.com/procgen/generators"
)

func Serve() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/list", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"generators": []string{
				"test",
			},
		})
	})

	r.GET("/generate/:width/:height/:generator", func(c *gin.Context) {

		widthStr := c.Param("width")
		heightStr := c.Param("height")

		generatorType := c.Param("generator")

		seedStr := c.Query("seed")

		seed, err := strconv.ParseInt(seedStr, 64, 64)
		if err != nil {
			seed = GetRandomSeed()
		}

		width, err := strconv.Atoi(widthStr)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Could not parse width parameter",
			})
			return
		}

		height, err := strconv.Atoi(heightStr)

		if err != nil {
			c.JSON(400, gin.H{
				"error": "Could not parse height parameter",
			})
			return
		}

		image, err := generators.Draw(generatorType, width, height, seed)

		if err != nil {
			c.JSON(500, "Internal Server Error")
			return
		}

		buf := new(bytes.Buffer)
		err = png.Encode(buf, image)

		if err != nil {
			c.JSON(500, "Internal Server Error")
			return
		}

		c.Data(200, "image/png", buf.Bytes())

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func GetRandomSeed() int64 {
	return time.Now().UnixNano()
}
