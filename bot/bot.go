package bot

import "net/http"

// Bot represents a LINE bot interface.
type Bot interface {
	http.Handler
}

// bot represents a LINE bot.
type bot struct {
}

func (b *bot) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

// New creates a LINE bot and returns it.
func New() Bot {
	return &bot{}
}
