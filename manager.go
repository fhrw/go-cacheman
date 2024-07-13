package cacheman

type ManagedCache[K comparable, V extendable] struct {
	Near Store[K, V]
	Far Store[K, V]
}

func (m ManagedCache[K, V]) Get(k K) (V, error) {
	if nearVal, err := m.Near.Get(k); err == nil {
		return nearVal, nil
	}
	if farVal, err := m.Far.Get(k); err == nil {
		m.Near.Add(k, farVal)
		return farVal, nil
	}
	return nil, ErrNotAvailable
}

func (m ManagedCache[K, V]) Add(k K, v V) error {
	if err := m.Far.Add(k, v); err != nil {
		return ErrFarWriteFail
	}	
	m.Near.Add(k, v)
	return nil
}