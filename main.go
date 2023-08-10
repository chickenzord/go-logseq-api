package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/chickenzord/go-logseq-api/internal/logseq"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	_ = godotenv.Overload()

	g := logseq.MarkdownGraph{BaseDir: "."}

	e := echo.New()
	e.POST("/journals/today", func(c echo.Context) error {
		bytes, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		filepath, err := g.AppendJournalToday(string(bytes))
		if err != nil {
			return err
		}

		return c.String(http.StatusOK, fmt.Sprintf("written: %s", filepath))
	})

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}
}
