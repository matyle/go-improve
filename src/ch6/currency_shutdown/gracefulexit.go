package main

import (
	"errors"
	"sync"
	"time"
)

type GracefullyShutdowner interface {
	Shutdown(t time.Duration) error
}

type ShutdownFunc func(time.Duration) error

func (f ShutdownFunc) Shutdown(t time.Duration) error {
	return f(t)
}

func CurrencyExit(t time.Duration, Shutdowners ...GracefullyShutdowner) error {
	c := make(chan struct{})

	go func() {
		var wg sync.WaitGroup
		for _, s := range Shutdowners {
			wg.Add(1)
			go func(shutdowners GracefullyShutdowner) {
				defer wg.Done()
				shutdowners.Shutdown(t)
			}(s)
		}
		wg.Wait()
		c <- struct{}{}
	}()

	timer := time.NewTimer(t)
	defer timer.Stop()
	select {
	case <-c:
		return nil
	case <-timer.C:
		return errors.New("timeout")
	}
}
