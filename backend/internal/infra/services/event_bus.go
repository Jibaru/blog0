package services

import "blog0/internal/domain"

type TriggerDevEventBus struct {
	triggerDev *TriggerDev
}

func NewTriggerDevEventBus(triggerDev *TriggerDev) *TriggerDevEventBus {
	return &TriggerDevEventBus{
		triggerDev: triggerDev,
	}
}

func (b *TriggerDevEventBus) ProcessEvents(event []any) error {
	for _, e := range event {
		switch evt := e.(type) {
		case *domain.PostCreated:
		case *domain.PostUpdated:
			_, err := b.triggerDev.GeneratePostAudio(evt.PostID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
