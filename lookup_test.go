package tzcoords

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestLookupByString(t *testing.T) {
	for s := range toCoords {
		lat, lon, err := ByString(s)
		require.NoError(t, err)
		require.NotZero(t, lat)
		require.NotZero(t, lon)
	}
}

func TestLookupByLocation(t *testing.T) {
	for s := range toCoords {
		loc, err := time.LoadLocation(s)
		require.NoError(t, err)
		lat, lon, err := ByLocation(loc)
		require.NoError(t, err)
		require.NotZero(t, lat)
		require.NotZero(t, lon)
	}
}

func TestLookupByLocation_Nil(t *testing.T) {
	lat, lon, err := ByLocation(nil)
	require.Error(t, err)
	require.ErrorAs(t, err, &ErrNotFound)
	require.Zero(t, lat)
	require.Zero(t, lon)
}
func TestLookupByString_Empty(t *testing.T) {
	lat, lon, err := ByString("")
	require.Error(t, err)
	require.ErrorAs(t, err, &ErrNotFound)
	require.Zero(t, lat)
	require.Zero(t, lon)
}

func TestLookupByString_UTC(t *testing.T) {
	lat, lon, err := ByString("UTC")
	require.Error(t, err)
	require.ErrorAs(t, err, &ErrNotFound)
	require.Zero(t, lat)
	require.Zero(t, lon)
}

func TestLookupByLocation_UTC(t *testing.T) {
	loc, err := time.LoadLocation("UTC")
	require.NoError(t, err)
	require.NotNil(t, loc)
	lat, lon, err := ByLocation(loc)
	require.Error(t, err)
	require.ErrorAs(t, err, &ErrNotFound)
	require.Zero(t, lat)
	require.Zero(t, lon)
}
