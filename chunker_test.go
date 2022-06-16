package PDL

import "testing"

func TestSplit(t *testing.T) {
	const size = 7
	const chunkSize = 3
	chunks, err := Split(size, chunkSize)
	require.NoError(t, err)
	expected := []Chunk{
		{0, 2},
		{3, 5},
		{6, 6},
	}

}
