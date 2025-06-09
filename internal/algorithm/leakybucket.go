// Package algorithm ...
package algorithm

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// LeakyBucket to init algorithm
type LeakyBucket struct {
	Limit     uint
	Windowms  time.Duration
	requestCh chan uint
	*sync.RWMutex
}

// Do does what it supposed to do...
func (l *LeakyBucket) Do(_ *http.Request) error {
	if len(l.requestCh) >= int(l.Limit) {
		return fmt.Errorf("error")
	}
	l.requestCh <- 1
	return nil
}

// SelfCleanUp ...
func (l *LeakyBucket) SelfCleanUp() {
	t := time.AfterFunc(l.Windowms, func() {
		l.Lock()
		defer l.Unlock()
		l.requestCh = make(chan uint)
	})
block:
	<-t.C
	goto block
}

// AfterFunc ...
func (l *LeakyBucket) AfterFunc(_ *http.Request) {
	<-l.requestCh
}
