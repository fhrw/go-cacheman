package cacheman

import ()

type Store[K comparable, V any] interface {
	Get(key K) (val V, err error)
	Add(key K, val V) error
}

type extendable interface {
	[]any
}