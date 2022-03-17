package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

// NewLoggingService logs each method call to s.
func NewLoggingService(s Service, logger log.Logger) Service {
	return &loggingService{
		s:      s,
		logger: logger,
	}
}

type loggingService struct {
	s      Service
	logger log.Logger
}

func (l *loggingService) PhysiciansExecute(ctx context.Context, request PhysiciansRequest) (response PhysiciansResponse, err error) {
	defer func(start time.Time) {
		l.logger.Log("event", "PhsyiciansExecute")
	}(time.Now())
	return l.s.PhysiciansExecute(ctx, request)
}

func (l *loggingService) ScheduleExecute(ctx context.Context, request ScheduleRequest) (response ScheduleResponse, err error) {
	defer func(start time.Time) {
		l.logger.Log("event", "ScheduleExecute")
	}(time.Now())
	return l.s.ScheduleExecute(ctx, request)
}

