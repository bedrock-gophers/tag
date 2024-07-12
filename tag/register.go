package tag

import (
	"errors"
	"fmt"
	"github.com/restartfu/gophig"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"os"
	"strings"
	"sync"
)

var (
	// tagMu is a mutex that protects the tags slice.
	tagMu sync.Mutex
	// tags is a slice of all tags.
	tags []Tag
	// tagsName is a map of all tags.
	tagsName = map[string]Tag{}
)

// register registers a tag.
func register(tgs ...Tag) {
	tagMu.Lock()
	for _, t := range tgs {
		tagsName[strings.ToLower(t.Name())] = t
		tags = append(tags, t)
	}
	tagMu.Unlock()
}

// Load loads all tag from a folder.
func Load(folder string) error {
	folder = strings.TrimSuffix(folder, "/")
	files, err := os.ReadDir(folder)
	if err != nil {
		return errors.New(fmt.Sprintf("error loading tag: %v", err))
	}

	var newTags []Tag
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		r, err := loadTag(folder + "/" + file.Name())
		if err != nil {
			return errors.New(fmt.Sprintf("error loading tag %s: %v", file.Name(), err))
		}
		newTags = append(newTags, r)
	}

	tagMu.Lock()
	tags = make([]Tag, 0)
	tagsName = map[string]Tag{}
	tagMu.Unlock()

	register(newTags...)
	return nil
}

// tagData is a struct that is used to decode tags from JSON.
type tagData struct {
	Name   string `json:"name"`
	Format string `json:"format"`
}

// loadTag loads a tag from a file.
func loadTag(filePath string) (Tag, error) {
	var data tagData
	err := gophig.GetConfComplex(filePath, gophig.JSONMarshaler{}, &data)
	if err != nil {
		return Tag{}, err
	}

	for _, r := range tags {
		if r.Name() == data.Name {
			return Tag{}, errors.New("tag with name " + data.Name + " already exists")
		}
	}

	return Tag{
		name:   data.Name,
		format: text.Colourf(data.Format),
	}, nil
}

// All returns all tag that are currently registered.
func All() []Tag {
	tagMu.Lock()
	t := make([]Tag, len(tags))
	copy(t, tags)
	tagMu.Unlock()
	return t
}

// ByName returns a tag by its name.
func ByName(name string) (Tag, bool) {
	tagMu.Lock()
	r, ok := tagsName[strings.ToLower(name)]
	tagMu.Unlock()
	return r, ok
}

// ByNameMust returns a tag by its name. If the tag does not exist, it panics.
func ByNameMust(name string) Tag {
	t, ok := ByName(name)
	if !ok {
		panic(fmt.Sprintf("tag %s does not exist", name))
	}
	return t
}
