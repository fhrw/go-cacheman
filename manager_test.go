package cacheman

import (
	"testing"
)

func Test_Get(t *testing.T) {
	tests := []struct{
		c ManagedCache[int, []any]
		key int
		want []any
	}{
		{
			c: ManagedCache[int, []any]{
				Near: Cache[int, []any]{},		
				Far: Cache[int, []any]{1: []any{"foo"}},
			},
			key: 1,
			want: []any{"foo"},
		},
		{
			c: ManagedCache[int, []any]{
				Near: Cache[int, []any]{1: []any{"foo"}},
				Far: Cache[int, []any]{},		
			},
			key: 1,
			want: []any{"foo"},
		},
	}

	for _, tt := range tests {
		got, err := tt.c.Get(tt.key)
		if err != nil {
			t.Errorf("ManagedCache.Get(%v) returned err: %v", tt.key, err)
		}
		if !Compare(got, tt.want){
			t.Errorf("ManagedCache.Get(%v) = %v, want %v", tt.key, got, tt.want)
		}
	}
}

func Compare[T comparable] (a, b []T) bool {
	if a == nil && b != nil { return false }
	if b == nil && a != nil { return false }
	
	if len(a) != len(b) { return false }
	
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}