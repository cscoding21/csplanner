package notification

import (
	"context"
	"csserver/internal/common"
)

// FindUserNotifications return a paged list of notifications specific to the logged in user
func (s *NotificationService) FindUserNotifications(ctx context.Context, paging common.Pagination) (common.PagedResults[Notification], error) {
	filters := common.Filters{Filters: []common.Filter{}}
	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)

	filters.AddFilter(common.Filter{
		Key:       "user_email",
		Value:     userEmail,
		Operation: common.FilterOperationEqual,
	})

	return s.FindNotifications(ctx, paging, filters)
}

// SetNotificationsRead mark the list of notification ids as read as long as the belong to the calling user
func (s *NotificationService) SetNotificationsRead(ctx context.Context, ids []string) error {
	sql := `
	update notification
	set is_read = true
	where true
		and user_id = $user
		and id in $ids
	`

	userEmail := s.ContextHelper.GetUserEmailFromContext(ctx)
	return s.DBClient.Execute(sql, map[string]interface{}{
		"user": userEmail,
		"ids":  ids,
	})
}
