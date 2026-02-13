package view

import (
	"fmt"
	"new-er/keyboard-cheatsheet/data/ast/nodes"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	activatedKeyStyle         = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF"))
	keepSequenceActivatedTime = time.Second
)

type nodeModel interface {
	Update(msg tea.Msg) (nodeModel, tea.Cmd)
	View() string
	Init() tea.Cmd
	GetActiveKeys() int
	IsActivated() bool
}

func InitialNodeModel(node nodes.Node) nodeModel {
	switch n := node.(type) {
	case nodes.AliasNode:
		return InitialAliasNodeModel(n)
	case nodes.AnyOfNode:
		return InitialAnyOfNodeModel(n)
	case nodes.ComboNode:
		return InitialComboNodeModel(n)
	case nodes.KeyNode:
		return InitialKeyNodeModel(n)
	case nodes.OptionalNode:
		return InitialOptionalNodeModel(n)
	case nodes.SequenceNode:
		return InitialSequenceNodeModel(n)
	default:
		panic("Unknown node type " + fmt.Sprintf("%T", n))
	}
}

type aliasNodeModel struct {
	node nodes.AliasNode
	key  nodeModel
}

func (m aliasNodeModel) Update(msg tea.Msg) (nodeModel, tea.Cmd) {
	key, _ := m.key.Update(msg)
	return aliasNodeModel{
		node: m.node,
		key:  key,
	}, nil
}
func (m aliasNodeModel) View() string {
	return m.node.Alias
}
func (m aliasNodeModel) Init() tea.Cmd {
	return nil
}
func InitialAliasNodeModel(node nodes.AliasNode) aliasNodeModel {
	return aliasNodeModel{node: node, key: InitialNodeModel(node.Key)}
}
func (m aliasNodeModel) GetActiveKeys() int {
	return m.key.GetActiveKeys()
}
func (m aliasNodeModel) IsActivated() bool {
	return m.key.IsActivated()
}

type anyOfNodeModel struct {
	node  nodes.AnyOfNode
	left  nodeModel
	right nodeModel
}

func (m anyOfNodeModel) Update(msg tea.Msg) (nodeModel, tea.Cmd) {
	left, _ := m.left.Update(msg)
	right, _ := m.right.Update(msg)
	return anyOfNodeModel{
		node:  m.node,
		left:  left,
		right: right,
	}, nil
}
func (m anyOfNodeModel) View() string {
	return fmt.Sprintf("%s|%s", m.left.View(), m.right.View())
}
func (m anyOfNodeModel) Init() tea.Cmd {
	return nil
}
func InitialAnyOfNodeModel(node nodes.AnyOfNode) anyOfNodeModel {
	return anyOfNodeModel{node: node, left: InitialNodeModel(node.Left), right: InitialNodeModel(node.Right)}
}
func (m anyOfNodeModel) GetActiveKeys() int {
	return m.left.GetActiveKeys() + m.right.GetActiveKeys()
}
func (m anyOfNodeModel) IsActivated() bool {
	return m.left.IsActivated() || m.right.IsActivated()
}

type comboNodeModel struct {
	node  nodes.ComboNode
	left  nodeModel
	right nodeModel
}

func (m comboNodeModel) Update(msg tea.Msg) (nodeModel, tea.Cmd) {
	left, _ := m.left.Update(msg)
	right, _ := m.right.Update(msg)
	return comboNodeModel{
		node:  m.node,
		left:  left,
		right: right,
	}, nil
}
func (m comboNodeModel) View() string {
	return fmt.Sprintf("%s+%s", m.left.View(), m.right.View())
}
func (m comboNodeModel) Init() tea.Cmd {
	return nil
}
func InitialComboNodeModel(node nodes.ComboNode) comboNodeModel {
	return comboNodeModel{node: node, left: InitialNodeModel(node.Left), right: InitialNodeModel(node.Right)}
}
func (m comboNodeModel) GetActiveKeys() int {
	return m.left.GetActiveKeys() + m.right.GetActiveKeys()
}
func (m comboNodeModel) IsActivated() bool {
	return m.left.IsActivated() && m.right.IsActivated()
}

type keyNodeModel struct {
	node         nodes.KeyNode
	isKeyPressed bool
}

func (m keyNodeModel) Update(msg tea.Msg) (nodeModel, tea.Cmd) {
	switch msg := msg.(type) {
	case PressedKeysChangedMsg:
		{
			for _, key := range msg.Keys {
				for _, name := range key.Names {
					if strings.EqualFold(name, m.node.Name) {
						return keyNodeModel{node: m.node, isKeyPressed: true}, nil
					}
				}
			}
		}
	}
	return keyNodeModel{node: m.node, isKeyPressed: false}, nil
}
func (m keyNodeModel) View() string {
	if m.isKeyPressed {
		return activatedKeyStyle.Render(m.node.Name)
	}
	return m.node.Name
}
func (m keyNodeModel) Init() tea.Cmd {
	return nil
}
func InitialKeyNodeModel(node nodes.KeyNode) keyNodeModel {
	return keyNodeModel{node: node, isKeyPressed: false}
}
func (m keyNodeModel) GetActiveKeys() int {
	if m.isKeyPressed {
		return 1
	}
	return 0
}
func (m keyNodeModel) IsActivated() bool {
	return m.isKeyPressed
}

type optionalNodeModel struct {
	node nodes.OptionalNode
	key  nodeModel
}

func (m optionalNodeModel) Update(msg tea.Msg) (nodeModel, tea.Cmd) {
	key, _ := m.key.Update(msg)
	return optionalNodeModel{
		node: m.node,
		key:  key,
	}, nil
}
func (m optionalNodeModel) View() string {
	return fmt.Sprintf("%s?", m.key.View())
}
func (m optionalNodeModel) Init() tea.Cmd {
	return nil
}
func InitialOptionalNodeModel(node nodes.OptionalNode) optionalNodeModel {
	return optionalNodeModel{node: node, key: InitialNodeModel(node.Key)}
}
func (m optionalNodeModel) GetActiveKeys() int {
	return m.key.GetActiveKeys()
}
func (m optionalNodeModel) IsActivated() bool {
	return m.key.IsActivated()
}

type sequenceNodeModel struct {
	node           nodes.SequenceNode
	left           nodeModel
	right          nodeModel
	activationTime time.Time
}

func (m sequenceNodeModel) Update(msg tea.Msg) (nodeModel, tea.Cmd) {
	left, _ := m.left.Update(msg)
	right, _ := m.right.Update(msg)

	return sequenceNodeModel{
		node:           m.node,
		left:           left,
		right:          right,
	}, nil
}
func (m sequenceNodeModel) View() string {
	return fmt.Sprintf("%s %s", m.left.View(), m.right.View())
}
func (m sequenceNodeModel) Init() tea.Cmd {
	return nil
}
func InitialSequenceNodeModel(node nodes.SequenceNode) sequenceNodeModel {
	return sequenceNodeModel{node: node, left: InitialNodeModel(node.Left), right: InitialNodeModel(node.Right)}
}
func (m sequenceNodeModel) GetActiveKeys() int {
	return m.left.GetActiveKeys() + m.right.GetActiveKeys()
}
func (m sequenceNodeModel) IsActivated() bool {
	return m.left.IsActivated() && m.right.IsActivated()
}
