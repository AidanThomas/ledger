package main

import (
	"fmt"
	"log"
	"strings"

	psql "github.com/AidanThomas/ledger/internal/adapaters"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	queryarea  textarea.Model
	resultarea textarea.Model
	err        error
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func initialModel() model {
	q := textarea.New()
	q.Placeholder = "Enter SQL query..."
	q.Focus()

	r := textarea.New()
	r.Placeholder = "Results will be here..."

	return model{
		queryarea:  q,
		resultarea: r,
		err:        nil,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if m.queryarea.Focused() {
				m.queryarea.Blur()
			}
		case tea.KeyEnter:
			psql := psql.NewPSQL()
			columns, result, err := psql.DoQuery(m.queryarea.Value())
			if err != nil {
				fmt.Println(err)
			}
			m.resultarea.SetValue(buildTable(columns, result))
		case tea.KeyCtrlC:
			return m, tea.Quit
		default:
			if !m.queryarea.Focused() {
				cmd = m.queryarea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	case tea.WindowSizeMsg:
		m.queryarea.SetWidth(msg.Width)
		m.resultarea.SetWidth(msg.Width)
	case error:
		m.err = msg
		return m, nil
	}

	m.queryarea, cmd = m.queryarea.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		m.queryarea.View(),
		m.resultarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func buildTable(columns []string, result [][]string) string {
	lengths := make([]int, len(columns))
	modifiedColumns := make([]string, len(columns))
	for i, c := range columns {
		longest := len(c)
		for _, row := range result {
			l := len(row[i])
			if l > longest {
				longest = l
			}
		}
		lengths[i] = longest
		modifiedColumns[i] = c + strings.Repeat(" ", longest-len(c))
	}

	out := strings.Join(modifiedColumns, " | ")
	out += "\n" + strings.Repeat("-", len(out)) + "\n"

	var rows []string
	for _, row := range result {
		modifiedRow := make([]string, len(row))
		for i, v := range row {
			modifiedRow[i] = v + strings.Repeat(" ", lengths[i]-len(v))
		}
		rows = append(rows, strings.Join(modifiedRow, " | "))
	}

	out += strings.Join(rows, "\n")
	return out
}
