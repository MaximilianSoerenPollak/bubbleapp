package main

import (
	"os"
	"time"

	"github.com/alexanderbh/bubbleapp/app"
	"github.com/alexanderbh/bubbleapp/component/button"
	"github.com/alexanderbh/bubbleapp/component/stack"
	"github.com/alexanderbh/bubbleapp/component/text"
	"github.com/alexanderbh/bubbleapp/shader"
	"github.com/alexanderbh/bubbleapp/style"

	zone "github.com/alexanderbh/bubblezone/v2"
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss/v2"
)

func NewRoot() model[struct{}] {
	ctx := &app.Context[struct{}]{
		Styles: style.DefaultStyles(),
		Zone:   zone.New(),
	}

	stack := stack.New(ctx, &stack.Options[struct{}]{
		Children: []*app.Base[struct{}]{
			text.New(ctx, "Shader examples:", nil),
			text.New(ctx, "Small Caps Shader", &text.Options{
				Foreground: ctx.Styles.Colors.Primary,
				Shader:     shader.NewSmallCapsShader(),
			}),
			button.New(ctx, " Blink ", &button.Options{
				Variant: button.Danger,
				Shader:  shader.NewBlinkShader(time.Second/3, lipgloss.NewStyle().Foreground(ctx.Styles.Colors.Success).BorderForeground(ctx.Styles.Colors.Success)),
			}),
		}},
	)

	base := app.New(ctx, app.AsRoot())
	base.AddChild(stack)

	return model[struct{}]{
		base: base,
	}
}

type model[T any] struct {
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
	p := tea.NewProgram(NewRoot(), tea.WithAltScreen(), tea.WithMouseAllMotion())
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
