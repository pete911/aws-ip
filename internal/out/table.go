package out

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
	"text/tabwriter"
)

type Table struct {
	logger *slog.Logger
	writer *tabwriter.Writer
}

func NewTable(logger *slog.Logger, output io.Writer) Table {
	return Table{
		logger: logger,
		writer: tabwriter.NewWriter(output, 1, 1, 2, ' ', 0),
	}
}

func (t Table) AddRow(columns ...string) {
	if _, err := fmt.Fprintln(t.writer, strings.Join(columns, "\t")); err != nil {
		t.logger.Error(fmt.Sprintf("table: add row: %v", err))
	}
}

func (t Table) Print() {
	if err := t.writer.Flush(); err != nil {
		t.logger.Error(fmt.Sprintf("table: print: %v", err))
	}
}

func FromInt(i int) string {
	if i == 0 {
		return "-"
	}
	return fmt.Sprintf("%d", i)
}

func TrimTo(in string, max int) string {
	if len(in) < max {
		return in
	}
	return fmt.Sprintf("%s...", in[:max-3])
}
