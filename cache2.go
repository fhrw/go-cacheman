package cacheman

type Cache2[K comparable, V extendable] map[K]V

func (c Cache2[K, V]) Get(k K) (V, error) {
	if cached, ok := c[k]; ok {
		return cached, nil
	}
	return nil, ErrNotAvailable
}

func (c Cache2[K, V]) Add(k K, v V) error {
	if cached, err := c.Get(k); err == nil {
		cached = append(cached, v)
		c[k] = cached
		return nil
	}
	c[k] = append([]any{}, v)
	return nil
}