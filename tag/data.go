package tag

import (
	"github.com/df-mc/atomic"
	"github.com/restartfu/gophig"
)

// tagsData is a struct that is used to encode tags to BSON or any other format that requires encoding.
type tagsData struct {
	Tags   []string `json:"tags"`
	Active string   `json:"active"`
}

func marshalTags(t *Tags, marshaler gophig.Marshaler) ([]byte, error) {
	var d tagsData
	t.tagMu.Lock()
	defer t.tagMu.Unlock()

	for _, tag := range t.tags {
		d.Tags = append(d.Tags, tag.Name())
	}

	if tg, active := t.Active(); active {
		d.Active = tg.Name()
	}
	return marshaler.Marshal(d)
}

func unmarshalTags(t *Tags, b []byte, marshaler gophig.Marshaler) error {
	var d tagsData
	if err := marshaler.Unmarshal(b, &d); err != nil {
		return err
	}

	for _, name := range d.Tags {
		if tag, ok := ByName(name); ok {
			t.Add(tag)
		}
	}

	if tag, ok := ByName(d.Active); ok && t.Contains(tag) {
		t.active = *atomic.NewValue(tag)
	}
	return nil
}
