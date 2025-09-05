package domain

type EventBus interface {
	ProcessEvents(event []any) error
}
