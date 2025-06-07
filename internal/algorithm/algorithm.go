// Package algorithm ...
package algorithm

import (
	"fmt"
	"net"
	"reflect"
	"time"

	"github.com/praveenmahasena/go-limiter/internal/config"
)

// Algorithm ...
type Algorithm struct {
	a map[string]Algo
}

// Algo ...
type Algo interface{}

// SelfCleanUp does what it's meant to do
func (a *Algorithm) SelfCleanUp() {
	for _, elem := range a.a {
		e := reflect.TypeOf(elem)
	}
}

// New to init algorithm
func New(algo []config.Rule) (*Algorithm, error) {
	mp := map[string]Algo{}

	for _, elem := range algo {
		if mp[elem.ID] != nil {
			return nil, fmt.Errorf("id %v already exists", elem.ID)
		}
		switch elem.Algorithm {
		case "leaky-bucket":
			mp[elem.ID] = LeakyBucket{elem.Limit, time.Duration(elem.Windowms), make(map[*net.IP]uint)}
		case "token-bucket":
			mp[elem.ID] = TokenBucket{elem.Limit, time.Duration(elem.Windowms), make(map[*net.IP]uint)}
		case "fixed-window-counter":
			mp[elem.ID] = FixedWindowCounter{elem.Limit, time.Duration(elem.Windowms), make(map[*net.IP]uint)}
		case "sliding-window-log":
			mp[elem.ID] = SlidingWindowLog{elem.Limit, time.Duration(elem.Windowms), make(map[*net.IP]uint)}
		case "sliding-window-counter":
			mp[elem.ID] = SlidingWindowCounter{elem.Limit, time.Duration(elem.Windowms), make(map[*net.IP]uint)}
		}
	}
	return &Algorithm{mp}, nil
}

// LeakyBucket to init algorithm
type LeakyBucket struct {
	Limit    uint
	Windowms time.Duration
	track    map[*net.IP]uint
}

// TokenBucket to init algorithm
type TokenBucket struct {
	Limit    uint
	Windowms time.Duration
	track    map[*net.IP]uint
}

// FixedWindowCounter to init algorithm
type FixedWindowCounter struct {
	Limit    uint
	Windowms time.Duration
	track    map[*net.IP]uint
}

// SlidingWindowLog to init algorithm
type SlidingWindowLog struct {
	Limit    uint
	Windowms time.Duration
	track    map[*net.IP]uint
}

// SlidingWindowCounter to init algorithm
type SlidingWindowCounter struct {
	Limit    uint
	Windowms time.Duration
	track    map[*net.IP]uint
}
