package tui

import (
	"log"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/AidanThomas/ledger/internal/ports"
	tea "github.com/charmbracelet/bubbletea"
)

var _ ports.UserInterface = (*TUI)(nil)

type TUI struct {
	ledger domain.Ledger

	active ViewName
	views  map[ViewName]View
}

func New(l domain.Ledger) *TUI {
	connView := NewConnStringView(l)
	queryView := NewQueryInputView(l)

	return &TUI{
		ledger: l,
		views: map[ViewName]View{
			ViewNameConn:  &connView,
			ViewNameQuery: &queryView,
		},
		active: ViewNameQuery,
	}
}

func (t *TUI) Run() error {
	t.ledger.Connect("postgres://postgres:password@localhost:5432/ledger_test?sslmode=disable")
	p := tea.NewProgram(t, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func (t *TUI) Init() tea.Cmd {
	view, ok := t.views[t.active]
	if !ok {
		log.Fatal("view not found")
	}
	return view.Activate()
}

func (t *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	view, ok := t.views[t.active]
	if !ok {
		log.Fatal("view not found")
	}
	return t, view.HandleMessage(msg)
}

func (t TUI) View() string {
	view, ok := t.views[t.active]
	if !ok {
		log.Fatal("view not found")
	}
	return view.GetView()
}
