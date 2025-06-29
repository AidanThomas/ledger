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

type RedrawMsg struct{}

func New(l domain.App) *TUI {
	t := &TUI{
		ledger: l,
		views:  make(map[ViewName]View),
		active: ViewNameConn,
	}
	t.registerView(NewConnStringView(l))
	t.registerView(NewQueryInputView(l))
	return t
}

func (t *TUI) Run() error {
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
	case ChangeViewMsg:
		t.active = msg.ChangeTo
		t.resolveView().HandleMessage(tea.WindowSizeMsg{
			Width:  t.windowWidth,
			Height: t.windowHeight,
		})
		t.resolveView().Activate()
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

func (t *TUI) registerView(v View) {
	n := v.Name()
	if _, ok := t.views[n]; !ok {
		t.views[n] = v
		return
	}
	log.Fatalf("trying to register already existing view: %s\n", n)
}
