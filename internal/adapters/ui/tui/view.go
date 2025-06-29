package tui

import tea "github.com/charmbracelet/bubbletea"

type ViewName string

const (
	ViewNameConn  ViewName = "connection_view"
	ViewNameQuery ViewName = "query_view"
)

type View interface {
	Name() ViewName
	Activate() tea.Cmd
	GetView() string
	HandleMessage(tea.Msg) tea.Cmd
}

func ChangeView(name ViewName) tea.Cmd {
	return func() tea.Msg {
		return ChangeViewMsg{name}
	}
}

type ChangeViewMsg struct {
	ChangeTo ViewName
}
