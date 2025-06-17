package activity

import "context"

// LogActivity logs a given activity
func (s *ActivityService) LogActivity(
	ctx context.Context,
	def ActivityDef,
	activityContext string,
	summary string) error {

	act := Activity{
		Type:    def.Type,
		Context: activityContext,
	}

	_, err := s.CreateActivity(ctx, act)
	return err
}
