package tui

import (
	"fmt"

	"github.com/AidanThomas/ledger/internal/app/ledger"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type TUI struct {
	ledger *ledger.Ledger

	queryarea  textarea.Model
	resultarea textarea.Model
	err        error
}

func New(l *ledger.Ledger) *TUI {
	q := textarea.New()
	q.Placeholder = "Enter SQL query..."
	q.Focus()

	r := textarea.New()
	r.Placeholder = "Results will be here..."

	return &TUI{
		ledger:     l,
		queryarea:  q,
		resultarea: r,
		err:        nil,
	}
}

func (t *TUI) Start() error {
	p := tea.NewProgram(t, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (t *TUI) Init() tea.Cmd {
	return textarea.Blink
}

func (t *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if t.queryarea.Focused() {
				t.queryarea.Blur()
			}
		case tea.KeyEnter:
			result, err := t.ledger.Execute(t.queryarea.Value())
			if err != nil {
				fmt.Println(err)
			}
			t.resultarea.SetValue(result)
		case tea.KeyCtrlC:
			return t, tea.Quit
		default:
			if !t.queryarea.Focused() {
				cmd = t.queryarea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	case tea.WindowSizeMsg:
		t.queryarea.SetWidth(msg.Width)
		t.resultarea.SetWidth(msg.Width)
	case error:
		t.err = msg
		return t, nil
	}

	t.queryarea, cmd = t.queryarea.Update(msg)
	cmds = append(cmds, cmd)
	return t, tea.Batch(cmds...)
}

func (t TUI) View() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		t.queryarea.View(),
		t.resultarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}
