package main

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	utcLoc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(fmt.Sprintf("failed to load location 'UTC': %s", err))
	}

	t.Run("2022-10-05t19:59:25.644225z", func(t *testing.T) {
		result, err := Parse("2022-10-05t19:59:25.644225z")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 10, 5, 19, 59, 25, 644225000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("2022-09-21T09:52:19Z", func(t *testing.T) {
		result, err := Parse("2022-09-21T09:52:19Z")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 9, 21, 9, 52, 19, 0, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("2022-10-05T16:27:08.419-04:00", func(t *testing.T) {
		result, err := Parse("2022-10-05T16:27:08.419-04:00")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 10, 5, 20, 27, 8, 419000000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("2022-10-05T09:10:19 EDT", func(t *testing.T) {
		t.Skip("location-dependent: https://github.com/itlightning/dateparse?tab=readme-ov-file#timezone-considerations")

		result, err := Parse("2022-10-05T09:10:19 EDT")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 10, 5, 13, 10, 19, 0, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (seconds)", func(t *testing.T) {
		result, err := Parse("1665001628")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 10, 5, 20, 27, 8, 0, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (milliseconds)", func(t *testing.T) {
		result, err := Parse("1665001628419")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 10, 5, 20, 27, 8, 419000000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (nanoseconds)", func(t *testing.T) {
		result, err := Parse("1665001628419000123")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 10, 5, 20, 27, 8, 419000123, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("unix time (seconds) with decimal", func(t *testing.T) {
		result, err := Parse("1667421543.4626768")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2022, 11, 2, 20, 39, 3, 462676800, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})

	t.Run("ULID", func(t *testing.T) {
		result, err := Parse("01D78XZ44G0000000000000000")
		if err != nil {
			t.Error(err)
		}
		expected := time.Date(2019, 03, 31, 03, 51, 23, 536000000, utcLoc)
		if !result.Equal(expected) {
			t.Errorf("expected %s; got %s", expected, result)
		}
	})
}
