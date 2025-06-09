// Package algorithm ...
package algorithm

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/praveenmahasena/go-limiter/internal/config"
)

// Algorithm ...
type Algorithm map[string]Algo

// Algo ...
type Algo interface {
	SelfCleanUp()
	Do(*http.Request) error
	AfterFunc(*http.Request)
}

// New to init algorithm
func New(algo []config.Rule) (Algorithm, error) {
	algorithm := make(Algorithm)
	for _, elem := range algo {
		path := fmt.Sprintf("%v %v", strings.ToUpper(elem.HTTPMethod), elem.Path)
		if algorithm[path] != nil {
			return nil, fmt.Errorf("path value with method %v already exists", path)
		}
		switch elem.Algorithm {
		case "leaky-bucket":
			algorithm[path] = &LeakyBucket{elem.Limit, time.Duration(elem.Windowms) * time.Millisecond, make(chan uint, elem.Limit), &sync.RWMutex{}}
		case "user-global-bucket":
			algorithm[path] = &UserGlobalBucket{elem.Limit, time.Duration(elem.Windowms) * time.Millisecond, make(map[string]chan uint, elem.Limit), &sync.RWMutex{}}
		default:
			return nil, fmt.Errorf("error: method algoritm does not exists")
		}
	}
	return algorithm, nil
}
