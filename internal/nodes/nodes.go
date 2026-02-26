package nodes

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type Type int

const (
	OptionNodeCard Type = iota
	OptionNodeCommand
)

type OptionNode struct {
	Title   string
	Options []string
	Next map[int]*OptionNode
	optType Type
	cursor  int
	cmd     tea.Cmd
}

func NewOptionNode(title string, options []string) *OptionNode {
	return &OptionNode{
		Title:   title,
		Options: options,
		Next:    make(map[int]*OptionNode),
		optType: OptionNodeCard,
		cursor:  0,
	}
}

// Explicit index; avoids "last element" bugs.
func (o *OptionNode) AddNext(index int, next *OptionNode) {
	if index < 0 || index >= len(o.Options) {
		return
	}
	o.Next[index] = next
}

func (o *OptionNode) SetCommand(cmd tea.Cmd) {
	o.cmd = cmd
	o.optType = OptionNodeCommand
}

func (o *OptionNode) Init() tea.Cmd { return nil }

func (o *OptionNode) View() tea.View {
	s := fmt.Sprintf("%s\n\n", o.Title)
	for i, choice := range o.Options {
		cursor := " "
		if o.cursor == i {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %d. %s\n", cursor, i, choice)
	}
	s += "\n↑/↓ (k/j)  enter  q\n"
	return tea.NewView(s)
}

func (o *OptionNode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return o, tea.Quit

		case "up", "k":
			if o.cursor > 0 {
				o.cursor--
			}
			return o, nil

		case "down", "j":
			if o.cursor < len(o.Options)-1 {
				o.cursor++
			}
			return o, nil

		case "enter":
			// If this node is a command node, execute.
			if o.optType == OptionNodeCommand && o.cmd != nil {
				return o, o.cmd
			}

			// Otherwise, navigate if there is a next node.
			if next, ok := o.Next[o.cursor]; ok && next != nil {
				return next, nil
			}
			return o, nil

		case "esc":
			return o, tea.Quit
		}
	}
	return o, nil
}
