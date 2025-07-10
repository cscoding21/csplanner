package activity

import (
	"csserver/internal/services/iam/appuser"
	"csserver/internal/services/project"
	"csserver/internal/utils/quilljs"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func getAddCommentActivityDelta(user appuser.Appuser, project project.Project) string {
	mentionEmbed := quilljs.NewMentionEmbed(user.ID, user.FirstName)
	delta := quilljs.New(nil).
		InsertEmbed(mentionEmbed,
			map[string]interface{}{
				"bold": true,
			}).
		Insert(" commented on the project ", nil).
		Insert(project.ProjectBasics.Name, map[string]interface{}{"bold": true})

	out, err := json.Marshal(delta)
	if err != nil {
		log.Errorf("getAddCommentActivityDelta: %s", err)
		return ""
	}

	return string(out)
}
