package tui

import (
	"fmt"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var _ View = (*ConnectionView)(nil)

type ConnectionView struct {
	ledger     domain.App
	viewChange chan ViewName

	connInput textinput.Model
}

func NewConnStringView(l domain.App, vc chan ViewName) ConnectionView {
	c := textinput.New()
	c.Placeholder = "BALLS"

	return ConnectionView{
		ledger:     l,
		viewChange: vc,

		connInput: c,
	}
}

func (v *ConnectionView) Activate() tea.Cmd {
	v.connInput.Focus()
	return textinput.Blink
}

func (v *ConnectionView) GetView() string {
	return fmt.Sprintf(
		"%s\n%s\n",
		"Input connection string",
		v.connInput.View(),
	)
}

func (v *ConnectionView) HandleMessage(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEsc:
			if v.connInput.Focused() {
				v.connInput.Blur()
			}
		case tea.KeyEnter:
			if err := v.ledger.Connect(v.connInput.Value()); err != nil {
				fmt.Println(err)
			} else {
				v.viewChange <- ViewNameQuery
			}
		case tea.KeyCtrlC:
			return tea.Quit
		default:
			if !v.connInput.Focused() {
				cmd = v.connInput.Focus()
				cmds = append(cmds, cmd)
			}
		}
	}
	v.connInput, cmd = v.connInput.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}
