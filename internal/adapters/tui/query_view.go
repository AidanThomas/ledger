package tui

import (
	"fmt"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

var _ View = (*QueryView)(nil)

type QueryView struct {
	ledger domain.Ledger

	queryarea  textarea.Model
	resultarea textarea.Model
}

func NewQueryInputView(l domain.Ledger) QueryView {
	q := textarea.New()
	q.Placeholder = "Enter SQL query..."

	r := textarea.New()
	r.Placeholder = "Results will be here..."

	return QueryView{
		ledger: l,

		queryarea:  q,
		resultarea: r,
	}
}

func (v *QueryView) Activate() tea.Cmd {
	v.queryarea.Focus()
	return textarea.Blink
}

func (v *QueryView) GetView() string {
	return fmt.Sprintf(
		"%s\n%s\n%s",
		v.queryarea.View(),
		v.resultarea.View(),
		"(ctrl+c to quit)",
	) + "\n\n"
}

func (v *QueryView) HandleMessage(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if v.queryarea.Focused() {
				v.queryarea.Blur()
			}
		case tea.KeyEnter:
			result, err := v.ledger.Execute(v.queryarea.Value())
			if err != nil {
				fmt.Println(err)
			}
			v.resultarea.SetValue(result)
		case tea.KeyCtrlC:
			return tea.Quit
		default:
			if !v.queryarea.Focused() {
				cmd = v.queryarea.Focus()
				cmds = append(cmds, cmd)
			}
		}
	case tea.WindowSizeMsg:
		v.queryarea.SetWidth(msg.Width)
		v.resultarea.SetWidth(msg.Width)
	}

	v.queryarea, cmd = v.queryarea.Update(msg)
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
