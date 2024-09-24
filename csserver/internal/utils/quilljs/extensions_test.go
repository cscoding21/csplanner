package quilljs

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestDeltaToString(t *testing.T) {
	delta := New(nil)
	want := "hello world"

	delta.Insert("hello ", map[string]interface{}{"bold": true})
	delta.Insert("world", nil)

	out := DeltaToString(*delta)

	if out != "hello world" {
		t.Errorf("DeltaToString: wanted %s - got %s", want, out)
	}
}

func TestDeltaToStringWithEmbed(t *testing.T) {
	userName := "Jeph"
	resourceName := "Test Resource"
	want := "Jeph created the resource Test Resource"
	mentionEmbed := NewMentionEmbed("user:jeph", userName)
	delta := New(nil).
		InsertEmbed(mentionEmbed,
			map[string]interface{}{
				"bold": true,
			}).
		Insert(" created the resource ", nil).
		Insert(resourceName, map[string]interface{}{"bold": true})

	out := DeltaToString(*delta)

	if out != want {
		t.Errorf("DeltaToString: wanted %s - got %s", want, out)
	}

	log.Info(out)
}
