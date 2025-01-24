package addstructs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBitArray(t *testing.T) {
	t.Parallel()

	t.Run("Set/Get bit", func(t *testing.T) {
		t.Parallel()

		array := NewBitArray(1000)

		err := array.Set(100)
		require.NoError(t, err)

		get, err := array.Get(100)
		require.NoError(t, err)
		require.Equal(t, 1, get)
	})

	t.Run("Set/Get bit by invalid index", func(t *testing.T) {
		t.Parallel()

		array := NewBitArray(1000)
		expectedErr := new(OutOfRange)

		err := array.Set(100000)
		errorAs(t, err, expectedErr)

		_, err = array.Get(100000)
		errorAs(t, err, expectedErr)

		err = array.Toggle(100000)
		errorAs(t, err, expectedErr)
	})

	t.Run("Toggle bit", func(t *testing.T) {
		t.Parallel()

		array := NewBitArray(1000)

		err := array.Toggle(10)
		require.NoError(t, err)

		get, err := array.Get(10)
		require.NoError(t, err)
		require.Equal(t, 1, get)

		err = array.Toggle(10)
		require.NoError(t, err)

		get, err = array.Get(10)
		require.NoError(t, err)
		require.Equal(t, 0, get)
	})

	t.Run("Clear array", func(t *testing.T) {
		t.Parallel()

		array := NewBitArray(1000)

		err := array.Set(10)
		require.NoError(t, err)

		array.Clear()

		get, err := array.Get(10)
		require.NoError(t, err)
		require.Equal(t, 0, get)
	})
}

// error.As require target error -> &&error
func errorAs(t *testing.T, err error, target error) {
	require.ErrorAs(t, err, &target)
}
