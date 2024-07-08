package tui

import (
	"fmt"
	"github.com/belphemur/mangal/key"
	"github.com/belphemur/mangal/provider"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/viper"
)

func (b *statefulBubble) Init() tea.Cmd {
	if names := viper.GetStringSlice(key.DownloaderDefaultSources); b.state != historyState && len(names) != 0 {
		var providers []*provider.Provider

		for _, name := range names {
			p, ok := provider.Get(name)
			if !ok {
				b.raiseError(fmt.Errorf("provider %s not found", name))
				return nil
			}

			providers = append(providers, p)
		}

		b.setState(loadingState)
		return tea.Batch(b.startLoading(), b.loadSources(providers), b.waitForSourcesLoaded())
	}

	return tea.Batch(textinput.Blink, b.loadProviders())
}
