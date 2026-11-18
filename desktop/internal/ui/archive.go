package ui

import "github.com/pointylab/pointy/desktop/internal/config"

// archiveSession ends a session Pointy started, so closing the card that
// asked in it leaves nothing behind in DSH's list. A session the user
// started is the plugin's to keep: it ignores the request. The tray turns
// the whole thing off.
func (u *UI) archiveSession(sessionID string) {
	if sessionID == "" || u.deps.Agent == nil || !autoArchives(u.deps.Settings) {
		return
	}
	if !u.deps.Agent.Archive(sessionID) {
		u.log("archive", "the plugin is not connected")
	}
}

// autoArchives reports whether closing a card archives its session; the
// setting is on until the user turns it off in the tray.
func autoArchives(settings *config.Store) bool {
	return settings == nil || settings.Get().AutoArchive
}
