package dictionary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Benchmark_match(b *testing.B) {
	dict := New()
	dict.Add("orange")
	dict.Add("range")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dict.match([]Term{{Value: "ran"}})
	}
}

func Test_match(t *testing.T) {
	dict := New()
	dict.Add("orange")
	dict.Add("range")

	t.Run("must return empty bitmap if nothin found", func(t *testing.T) {
		m := dict.match([]Term{{Value: "qwe"}})
		require.NotNil(t, m)
		require.True(t, m.IsEmpty())
	})

	t.Run("must be able to find word by terms", func(t *testing.T) {
		m := dict.match([]Term{{Value: "ora"}})
		require.False(t, m.IsEmpty())
		require.Equal(t, uint64(1), m.GetCardinality())
	})

	t.Run("must be able to find multiple words by terms", func(t *testing.T) {
		m := dict.match([]Term{{Value: "ran"}})
		require.False(t, m.IsEmpty())
		require.Equal(t, uint64(2), m.GetCardinality())
	})
}
