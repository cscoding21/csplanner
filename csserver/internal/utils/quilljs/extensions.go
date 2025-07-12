package quilljs

import (
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"
)

//https://github.com/fmpwizard/go-quilljs-delta/blob/main/delta/delta.go

// NewMentionEmbed creates a new embedding for a delta
func NewMentionEmbed(userID string, name string) Embed {
	embed := Embed{
		Key: "mention",
		Value: map[string]any{
			"index":          "0",
			"denotationChar": "@",
			"id":             userID,
			"value":          name,
		},
	}

	return embed
}

// DeltaToString converts a delta to a raw text string
func DeltaToString(delta Delta) string {
	var builder strings.Builder

	for _, o := range delta.Ops {
		if o.InsertEmbed != nil {
			if o.InsertEmbed.Key == "mention" {
				builder.WriteString(o.InsertEmbed.Value.(map[string]any)["value"].(string))
			}
		} else {
			builder.WriteString(string(o.Insert))
		}
	}

	return builder.String()
}

// StringToDelta takes a simple string and converts it into a single-insert delta
func StringToDelta(text string) *Delta {
	var d Delta
	json.Unmarshal([]byte(text), &d)

	return &d
}

// StringToDeltaString takes a simple string and converts it into a single-insert delta and returns a JSON representation
func StringToDeltaString(text string) string {
	delta := StringToDelta(text)

	out, err := json.Marshal(delta)
	if err != nil {
		log.Errorf("StringToDeltaString: %s", err)
	}

	return string(out)
}
