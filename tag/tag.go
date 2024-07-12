package tag

import (
	"github.com/df-mc/atomic"
	"slices"
	"sync"
)

// Tag represents a tag that can be applied to a player. It is used to give players a certain 'tag' in chat.
type Tag struct {
	// name is the name of the tag. It is used to identify the tag and to compare it to other tags.
	name string
	// format is the format of the tag. It is used to display the tag in chat.
	format string
}

// Name returns the name of the tag. The name is used to identify the tag and to compare it to other tags.
func (t Tag) Name() string {
	return t.name
}

// Format returns the format of the tag. The format is used to display the tag in chat.
func (t Tag) Format() string {
	return t.format
}

// Tags represents a list of tags that can be applied to a player.
type Tags struct {
	tagMu sync.Mutex
	tags  []Tag

	active atomic.Value[Tag]
}

// Active returns the active tag of the list of tags.
func (t *Tags) Active() (Tag, bool) {
	tag := t.active.Load()
	return tag, tag != Tag{}
}

// UpdateActive updates the active tag of the list of tags.
func (t *Tags) UpdateActive(tag Tag) {
	t.active.Store(tag)
}

// Add adds a tag to the list of tags.
func (t *Tags) Add(tag Tag) {
	t.tagMu.Lock()
	defer t.tagMu.Unlock()
	t.tags = append(t.tags, tag)
}

// Remove removes a tag from the list of tags.
func (t *Tags) Remove(tag Tag) {
	t.tagMu.Lock()
	defer t.tagMu.Unlock()
	i := slices.IndexFunc(t.tags, func(other Tag) bool {
		return tag == other
	})
	t.tags = slices.Delete(t.tags, i, i+1)
}

// Contains returns true if the list of tags contains the tag provided.
func (t *Tags) Contains(tag Tag) bool {
	t.tagMu.Lock()
	defer t.tagMu.Unlock()
	return slices.Contains(t.tags, tag)
}

// All returns all tags that are currently applied to the list of tags.
func (t *Tags) All() []Tag {
	t.tagMu.Lock()
	defer t.tagMu.Unlock()
	return t.tags
}
