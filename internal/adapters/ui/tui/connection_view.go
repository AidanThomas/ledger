package tui

import (
	"fmt"
	"log"

	"github.com/AidanThomas/ledger/internal/domain"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var _ View = (*ConnectionView)(nil)

type ConnectionView struct {
	ledger   domain.App
	connList list.Model
}

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type connection struct{ name, conn string }

func (c connection) Title() string       { return c.name }
func (c connection) Description() string { return c.conn }
func (c connection) FilterValue() string { return c.name }

func NewConnStringView(l domain.App) *ConnectionView {
	savedConns, err := l.GetConnections()
	if err != nil {
		log.Fatalf("cannot get list of saved connections: %s", err)
	}
	connections := make([]list.Item, len(savedConns))
	for i, c := range savedConns {
		connections[i] = connection{
			name: c.Name,
			conn: c.Conn,
		}
	}

	li := list.New(connections, list.NewDefaultDelegate(), 0, 0)
	li.Title = "Select a connection"

	return &ConnectionView{
		ledger:   l,
		connList: li,
	}
}

func (v *ConnectionView) Name() ViewName { return ViewNameConn }

func (v *ConnectionView) Activate() tea.Cmd {
	return nil
}

func (v *ConnectionView) GetView() string {
	return docStyle.Render(v.connList.View())
}

func (v *ConnectionView) HandleMessage(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			c := v.connList.SelectedItem().(connection)
			if err := v.ledger.Connect(c.conn); err != nil {
				fmt.Println(err)
			} else {
				cmds = append(cmds, ChangeView(ViewNameQuery))
			}
		case tea.KeyCtrlC:
			return tea.Quit
		}
	case tea.WindowSizeMsg:
		hor, ver := docStyle.GetFrameSize()
		v.connList.SetSize(msg.Width-hor, msg.Height-ver)
	}
	v.connList, cmd = v.connList.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}
