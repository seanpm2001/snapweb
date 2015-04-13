package webprogress

import (
	"sync"
)

type Status struct {
	status map[string]*WebProgress
	l      sync.Mutex
}

func New() *Status {
	return &Status{status: make(map[string]*WebProgress)}
}

// Add add pkg to the list of progresses to track, it is idempotent
func (i *Status) Add(pkg string) *WebProgress {
	i.l.Lock()
	defer i.l.Unlock()

	if _, ok := i.status[pkg]; !ok {
		i.status[pkg] = NewWebProgress()
	}

	return i.status[pkg]
}

// Remove removes pkg to the list of progresses to track, it is a no op
// to remove multiple times.
func (i *Status) Remove(pkg string) {
	i.l.Lock()
	defer i.l.Unlock()

	if _, ok := i.status[pkg]; ok {
		delete(i.status, pkg)
	}
}

// Get returns a *WebProgress corresponding to pkg or nil if not tracked.
func (i *Status) Get(pkg string) (*WebProgress, bool) {
	i.l.Lock()
	defer i.l.Unlock()

	w, ok := i.status[pkg]
	return w, ok
}
