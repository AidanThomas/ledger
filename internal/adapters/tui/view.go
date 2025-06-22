package tui

import tea "github.com/charmbracelet/bubbletea"

type ViewName int

const (
	ViewNameConn = iota
	ViewNameQuery
)

type View interface {
	Activate() tea.Cmd
	GetView() string
	HandleMessage(tea.Msg) tea.Cmd
}
