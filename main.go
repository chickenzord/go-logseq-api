package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/chickenzord/go-logseq-api/internal/logseq"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Config struct {
	LogseqBaseDir string `envconfig:"logseq_base_dir" default:"."`
	BindHost      string `envconfig:"bind_host" default:"0.0.0.0"`
	BindPort      string `envconfig:"bind_port" default:"8080"`
}

func (c *Config) BindAddress() string {
	return fmt.Sprintf("%s:%s", c.BindHost, c.BindPort)
}

func main() {
	_ = godotenv.Overload()

	var cfg Config

	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}

	g := logseq.MarkdownGraph{BaseDir: cfg.LogseqBaseDir}

	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
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

	if err := e.Start(cfg.BindAddress()); err != nil {
		panic(err)
	}
}
