package internal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseISO6709Pair_Troll(t *testing.T) {
	// Antarctica/Troll
	const str = "-720041+0023206"
	ll, err := ParseISO6709Pair(str)
	require.NoError(t, err)
	require.NotZero(t, ll.Lat)
	require.NotZero(t, ll.Lon)
	require.Less(t, ll.Lat, 0.0)
}
func TestParseISO6709Pair_Jerusalem(t *testing.T) {
	// Asia/Jerusalem
	const str = "+314650+0351326"
	ll, err := ParseISO6709Pair(str)
	require.NoError(t, err)
	require.NotZero(t, ll.Lat)
	require.NotZero(t, ll.Lon)
}

func TestParseISO6709Pair_Shanghai(t *testing.T) {
	// Asia/Shanghai
	const str = "+3114+12128"
	ll, err := ParseISO6709Pair(str)
	require.NoError(t, err)
	require.NotZero(t, ll.Lat)
	require.NotZero(t, ll.Lon)
}

func TestParseISO6709Pair_Cambridge_Bay(t *testing.T) {
	// America/Cambridge_Bay
	const str = "+690650-1050310"
	ll, err := ParseISO6709Pair(str)
	require.NoError(t, err)
	require.NotZero(t, ll.Lat)
	require.NotZero(t, ll.Lon)
}
