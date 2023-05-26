package printer

import (
	"fmt"
	"strings"
)

// ASCIITablePrinter to STDOUT
type ASCIITablePrinter struct {
	headers []string
	rows    [][]string
}

// NewASCIITablePrinter simple table printer
func NewASCIITablePrinter() *ASCIITablePrinter {
	return &ASCIITablePrinter{
		headers: make([]string, 0),
		rows:    make([][]string, 0),
	}
}

// AddHeader to table
func (t *ASCIITablePrinter) AddHeader(headers []string) {
	t.headers = headers
}

// AddRow to table
func (t *ASCIITablePrinter) AddRow(values []string) {
	t.rows = append(t.rows, values)
}

func (t *ASCIITablePrinter) String() string {
	var builder strings.Builder

	// Print headers
	for _, header := range t.headers {
		builder.WriteString(fmt.Sprintf("| %-10s ", header))
	}
	builder.WriteString("|\n")

	// Print separator
	separator := strings.Repeat("-", (11*len(t.headers))+1)
	builder.WriteString(separator + "\n")

	// Print rows
	for _, row := range t.rows {
		for _, value := range row {
			builder.WriteString(fmt.Sprintf("| %-10s ", value))
		}
		builder.WriteString("|\n")
	}

	// Print bottom separator
	builder.WriteString(separator + "\n")

	return builder.String()
}
