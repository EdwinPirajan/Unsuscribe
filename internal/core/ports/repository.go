package ports

type UnsubscribeRepository interface {
	SaveUnsubscribe(email string) error
}

type UnsubscribeService interface {
	Unsubscribe(email string) error
}
