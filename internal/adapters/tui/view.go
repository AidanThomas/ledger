package tui

import tea "github.com/charmbracelet/bubbletea"

type ViewName string

const (
	ViewNameConn  ViewName = "connection_view"
	ViewNameQuery ViewName = "query_view"
)

type View interface {
	Activate() tea.Cmd
	GetView() string
	HandleMessage(tea.Msg) tea.Cmd
}
