package logseq

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"cloud.google.com/go/civil"
)

type MarkdownGraph struct {
	BaseDir string
}

func (g *MarkdownGraph) AppendJournal(d civil.Date, entries ...string) (string, error) {
	filepath := filepath.Join(g.BaseDir, "journals", fmt.Sprintf("%d_%02d_%02d.md",
		d.Year, d.Month, d.Day,
	))

	f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	for _, entry := range entries {
		content := strings.ReplaceAll(entry, "\n", " ")
		if _, err := fmt.Fprintf(f, "- %s\n", content); err != nil {
			return "", err
		}
	}

	return filepath, nil
}

func (g *MarkdownGraph) AppendJournalToday(entries ...string) (string, error) {
	return g.AppendJournal(civil.DateOf(time.Now()))
}
