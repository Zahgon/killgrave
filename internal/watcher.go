package killgrave

import (
	"github.com/radovskyb/watcher"
)

// InitializeWatcher initialize a watcher to check for modification on all files
// in the given path to watch
func InitializeWatcher(pathToWatch string) (*watcher.Watcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AttachWatcher start the watcher, if any error was produced while the starting process the application would crash
// you need to pass a function, this function is the function that will be executed when the watcher
// receive any event the type of defined on the InitializeWatcher function
func AttachWatcher(w *watcher.Watcher, fn func()) { _ = "STUB: not implemented"; return }

func CloseWatcher(w *watcher.Watcher) { _ = "STUB: not implemented"; return }

func readEventsFromWatcher(w *watcher.Watcher, fn func()) { _ = "STUB: not implemented"; return }
