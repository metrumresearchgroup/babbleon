package repository

type NotificationRepository interface {
	SendNotification(notification string, destination string) error
}
