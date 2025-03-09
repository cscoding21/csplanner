package notification

/*
// FindUserNotifications return a paged list of notifications specific to the logged in user
func (s *NotificationService) FindUserNotifications(ctx context.Context, paging common.Pagination) (common.PagedResults[Notification], error) {
	filters := common.Filters{Filters: []common.Filter{}}
	userEmail := config.GetUserEmailFromContext(ctx)

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
		and user_id = $1
		and id in $2
	`

	userEmail := config.GetUserEmailFromContext(ctx)
	_, err := s.db.Exec(ctx, sql, userEmail, ids)

	return err
}
*/
