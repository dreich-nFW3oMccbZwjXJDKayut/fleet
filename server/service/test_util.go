package service

import (
	"github.com/WatchBeam/clock"
	kitlog "github.com/go-kit/kit/log"
	"github.com/kolide/kolide-ose/server/config"
	"github.com/kolide/kolide-ose/server/kolide"
)

func newTestService(ds kolide.Datastore) (kolide.Service, error) {
	return NewService(ds, kitlog.NewNopLogger(), config.TestConfig(), nil, clock.C)
}

func newTestServiceWithClock(ds kolide.Datastore, c clock.Clock) (kolide.Service, error) {
	return NewService(ds, kitlog.NewNopLogger(), config.TestConfig(), nil, c)
}