// Package algorithm ..
package algorithm

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// UserGlobalBucket ...
type UserGlobalBucket struct {
	Limit    uint
	Windowms time.Duration
	reqStack map[string]chan uint
	*sync.RWMutex
}

// Do that does
func (u *UserGlobalBucket) Do(r *http.Request) error {
	u.Lock()
	defer u.Unlock()
	ip := r.RemoteAddr
	if len(u.reqStack) >= int(u.Limit) {
		return fmt.Errorf("error")
	}
	if u.reqStack[ip] == nil {
		u.reqStack[ip] = make(chan uint, u.Limit)
	}
	u.reqStack[ip] <- 1
	return nil
}

// AfterFunc ...
func (u *UserGlobalBucket) AfterFunc(r *http.Request) {
	<-u.reqStack[r.RemoteAddr]
}

// SelfCleanUp ...
func (u *UserGlobalBucket) SelfCleanUp() {
	t := time.AfterFunc(u.Windowms, func() {
		u.Lock()
		defer u.Unlock()
		u.reqStack = make(map[string]chan uint)
	})
block:
	<-t.C
	goto block
}
