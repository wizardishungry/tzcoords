package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type LatLon struct {
	Lat, Lon float64
}

type Angle struct {
	sign, degrees, minutes, seconds int
}

func (a *Angle) ToFloat64() float64 {
	const (
		minutesFactor = (1.0 / 60.0)
		secondsFactor = (1.0 / 3600.0)
	)
	f := float64(a.degrees) +
		float64(a.minutes)*minutesFactor +
		float64(a.seconds)*secondsFactor
	return f * float64(a.sign)
}

func ParseISO6709Pair(str string) (LatLon, error) {
	s := func(r rune) bool {
		return r == '-' || r == '+'
	}

	latStart := strings.IndexFunc(str, s)
	lonStart := strings.IndexFunc(str[latStart+1:], s) + latStart + 1

	lat := str[latStart:lonStart]
	lon := str[lonStart:]

	latF, err := ParseISO6709(lat, 2, 90)
	if err != nil {
		return LatLon{}, err
	}
	lonF, err := ParseISO6709(lon, 3, 180)
	if err != nil {
		return LatLon{}, err
	}

	return LatLon{
		Lat: latF.ToFloat64(),
		Lon: lonF.ToFloat64(),
	}, nil
}

func ParseISO6709(str string, maxDigits, maxDegrees int) (Angle, error) {
	a := Angle{}

	a.sign = 1
	if str[0] == '-' {
		a.sign = -1
	}
	str = str[1:]

	l := len(str)

	if (l-maxDigits)%2 != 0 || (l-maxDigits) > 4 || (l-maxDigits) < 0 {
		return a, fmt.Errorf("bad length %d", l)
	}

	d, err := strconv.Atoi(str[:maxDigits])
	if err != nil {
		return a, err
	}
	if d >= maxDegrees {
		return a, fmt.Errorf("absolute value of angle outside of range (%d >= %d)", d, maxDegrees)
	}
	a.degrees = d

	str = str[maxDigits:]

	if len(str) > 0 {
		mStr := str[:2]

		m, err := strconv.Atoi(mStr)
		if err != nil {
			return a, err
		}
		a.minutes = m
		sStr := str[2:]
		if len(sStr) > 0 {

			s, err := strconv.Atoi(sStr)
			if err != nil {
				return a, err
			}
			a.seconds = s
		}
	}

	return a, nil
}
