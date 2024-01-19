package service

import "time"

type serviceOption func(*Service)

// WithKeys setup a basic notes for track.
func WithKeys(v []string) serviceOption {
	return func(s *Service) {
		s.keys = v
	}
}

// WithScales setup a basic notes for track.
func WithScales(v []string) serviceOption {
	return func(s *Service) {
		s.scales = v
	}
}

// WithInstruments setup service instrument list.
func WithInstruments(v []string) serviceOption {
	return func(s *Service) {
		s.instruments = v
	}
}

func WithUpdateInterval(d time.Duration) serviceOption {
	return func(s *Service) {
		s.updateInterval = d
	}
}

// WithInstruments setup service instrument list.
func WithGreetings(v []string) serviceOption {
	return func(s *Service) {
		s.greetings = v
	}
}
