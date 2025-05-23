package main

import (
	"os"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/box"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
)

type CustomData struct{}

func NewRoot() model[CustomData] {
	ctx := &app.Context[CustomData]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
	}

	stack := stack.New(ctx, &stack.Options[CustomData]{
		Children: []*app.Base[CustomData]{
			box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Danger}),
			box.New(ctx, &box.Options[CustomData]{
				Child: stack.New(ctx, &stack.Options[CustomData]{
					Horizontal: true,
					Children: []*app.Base[CustomData]{
						box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Primary}),
						box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Secondary}),
						box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Tertiary}),
					}},
				),
			}),
			box.New(ctx, &box.Options[CustomData]{Bg: ctx.Styles.Colors.Warning}),
		}},
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stack)

	return model[CustomData]{
		base: base,
	}
}

type model[T CustomData] struct {
	base *app.Base[T]
}

func (m model[T]) Init() tea.Cmd {
	return m.base.Init()
}

func (m model[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	cmd := m.base.Update(msg)

	return m, cmd

}

func (m model[T]) View() string {
	return m.base.Render()
}

func main() {
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
