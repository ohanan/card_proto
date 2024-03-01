package prompt

import (
	"os"
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ohanan/card_proto/pkg/protoservice"
	"github.com/ohanan/card_proto/pkg/protoservice/proto"
)

func Run(p protoservice.PluginXServer) {
	m := &Model{
		ch:        make(chan *proto.ActionResult),
		thisIDMap: map[uint64]struct{}{},
	}
	ps := protoservice.NewPluginServerFromXServer(p)
	protoservice.TyrInitHelper(protoservice.NewHostServerFromXServer(m), ps)
	program := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion())
	m.p = program
	client := protoservice.NewPluginXClient(protoservice.NewPluginClientFromServer(ps))
	m.PluginXClient = client
	client.StartMode("a", 1)
	_, _ = program.Run()
}

var _ protoservice.HostXServer = (*Model)(nil)
var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type Model struct {
	*protoservice.PluginXClient
	req           *proto.AskAction_Req
	thisIDMap     map[uint64]struct{}
	thisIDList    []int32
	p             *tea.Program
	ch            chan *proto.ActionResult
	width, height int
}

func (m *Model) GetPlayerInfo(req *proto.GetPlayerInfo_Req, resp *proto.GetPlayerInfo_Resp) {
}

func (m *Model) RegisterNotify(req *proto.RegisterNotify_Req, resp *proto.RegisterNotify_Resp) {
}

func (m *Model) AskAction(req *proto.AskAction_Req, resp *proto.AskAction_Resp) {
	m.p.Send(req)
	resp.Result = <-m.ch
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case *proto.AskAction_Req:
		m.req = msg
	case tea.MouseMsg:
		if msg.Action != tea.MouseActionPress {
			return m, nil
		}
		idx := msg.Y - 2
		if idx >= 0 && idx < len(m.thisIDList) {
			id := m.thisIDList[idx]
			m.thisIDList = m.thisIDList[:0]
			m.thisIDMap = map[uint64]struct{}{}
			m.ch <- &proto.ActionResult{
				Id:   id,
				Type: proto.UserActionType_PRIMARY,
			}
		}
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			os.Exit(0)
		}
		if len(msg.Runes) == 1 {
			id, err := strconv.ParseUint(string(msg.Runes), 36, 32)
			if err == nil {
				if _, ok := m.thisIDMap[id]; ok {
					m.thisIDList = m.thisIDList[:0]
					m.thisIDMap = map[uint64]struct{}{}
					m.ch <- &proto.ActionResult{
						Id:   int32(id),
						Type: proto.UserActionType_PRIMARY,
					}
					return m, nil
				}
			}

		}
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	}
	return m, nil
}
func getFlag(option *proto.Action_ActionOption) string {
	if option.Disabled {
		return "x"
	}
	if option.Selected {
		return "v"
	}
	return " "
}
func (m *Model) View() string {
	if m.req == nil {
		return "waiting"
	}
	var rows []table.Row
	maxIDLength := 2
	maxNameLength := 4
	maxTypeLength := 4
	for _, action := range m.req.Actions {
		switch a := action.Option.(type) {
		case *proto.Action_Card:
			for _, option := range a.Card.Options {
				id := strconv.FormatUint(uint64(option.Option.Id), 36)
				m.thisIDList = append(m.thisIDList, option.Option.Id)
				maxIDLength = max(maxIDLength, len(id))
				name := option.Option.Name
				maxNameLength = max(maxNameLength, len(name))
				maxTypeLength = max(maxTypeLength, 4)
				r := []string{
					getFlag(option.Option),
					id,
					"CARD",
					name,
					option.Card.String(),
				}

				rows = append(rows, r)
			}

		}
		action.GetCard()
	}
	t := table.New(
		table.WithColumns([]table.Column{
			{
				Title: "",
				Width: 1,
			},
			{
				Title: "ID",
				Width: maxIDLength,
			},
			{
				Title: "Type",
				Width: maxTypeLength,
			},
			{
				Title: "Name",
				Width: maxNameLength + 2,
			},
			{
				Title: "Data",
				Width: m.width - maxIDLength - 1 - maxTypeLength - maxNameLength - 16,
			},
		}),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(m.height-4),
	)
	return baseStyle.Render(t.View()) + "\n"
}
