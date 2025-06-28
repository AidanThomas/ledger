package tui

import (
	"log"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/AidanThomas/ledger/internal/ports"
	tea "github.com/charmbracelet/bubbletea"
)

var _ ports.UserInterface = (*TUI)(nil)

type TUI struct {
	ledger domain.App

	active     ViewName
	views      map[ViewName]View
	viewChange chan ViewName

	windowHeight int
	windowWidth  int
}

func New(l domain.App) *TUI {
	vc := make(chan ViewName)

	connView := NewConnStringView(l, vc)
	queryView := NewQueryInputView(l, vc)

	return &TUI{
		ledger: l,
		views: map[ViewName]View{
			ViewNameConn:  &connView,
			ViewNameQuery: &queryView,
		},
		active:     ViewNameConn,
		viewChange: vc,
	}
}

func (t *TUI) Run() error {
	go t.listenForViewChange()

	p := tea.NewProgram(t, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
}

func (t *TUI) Init() tea.Cmd {
	return t.resolveView().Activate()
}

func (t *TUI) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.windowHeight = msg.Height
		t.windowWidth = msg.Width
	}
	return t, t.resolveView().HandleMessage(msg)
}

func (t *TUI) View() string {
	return t.resolveView().GetView()
}

func (t *TUI) resolveView() View {
	view, ok := t.views[t.active]
	if !ok {
		log.Fatal("view not found")
	}
	return view
}

func (t *TUI) listenForViewChange() {
	vn := <-t.viewChange
	t.active = vn

	view := t.resolveView()
	view.Activate()
	view.HandleMessage(tea.WindowSizeMsg{
		Width:  t.windowWidth,
		Height: t.windowHeight,
	})
}
