package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Options struct {
	Continue bool
	Install  bool
}

func Run(options *Options) (tea.Model, error) {

	bubble := newBubble()

	if options.Install {
		bubble.newState(scrapersInstallState)
	} else if options.Continue {
		_, err := bubble.loadHistory()
		if err != nil {
			return nil, err
		}

		bubble.newState(historyState)
	} else {
		bubble.newState(sourcesState)
	}

	return tea.NewProgram(bubble, tea.WithAltScreen()).Run()
}
