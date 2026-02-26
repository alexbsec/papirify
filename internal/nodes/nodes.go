package nodes

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type OptionNode struct {
	Title    string
	Options  []string
	Selected map[int]OptionNode
	cursor   int
}

func NewOptionNode(title string, options []string) OptionNode {
	return OptionNode{
		Title:    title,
		Options:  options,
		Selected: make(map[int]OptionNode),
		cursor:   0,
	}
}

func (o OptionNode) AddOption(option string, next OptionNode) OptionNode {
	o.Options = append(o.Options, option)
	o.Selected[len(o.Options)-1] = next
	return o
}

func (o OptionNode) Init() tea.Cmd {
	return nil
}

func (o OptionNode) View() tea.View {
	s := o.Title + "\n\n"
	for i, option := range o.Options {
		cursor := " "
		if _, ok := o.Selected[i]; ok {
			cursor = ">"
		}
		s += fmt.Sprintf("%s %s\n", cursor, option)
	}
	return tea.NewView(s)
}

func (o OptionNode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return o, tea.Quit
		case "up", "k":
			if o.cursor > 0 {
				o.cursor--
			}
		case "down", "j":
			if o.cursor < len(o.Options)-1 {
				o.cursor++
			}

		case "enter":
			if nextOpt, ok := o.Selected[o.cursor]; ok {
				delete(o.Selected, o.cursor)
				return nextOpt, nil
			}
		}
	}

	return o, nil
}
