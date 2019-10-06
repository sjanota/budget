package models_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/sjanota/budget/backend/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDate_MarshalGQL(t *testing.T) {
	cases := []struct {
		date     models.Date
		expected string
	}{
		{models.Date{2020, time.January, 9}, "\"2020-01-09\""},
		{models.Date{2020, time.December, 20}, "\"2020-12-20\""},
	}
	for _, test := range cases {
		t.Run(test.expected, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			test.date.MarshalGQL(buffer)
			assert.Equal(t, test.expected, buffer.String())
		})
	}
}

func TestDate_UnmarshalGQL(t *testing.T) {
	cases := []struct {
		date     string
		expected *models.Date
	}{
		{"\"2020-01-09\"", &models.Date{2020, time.January, 9}},
		{"\"2020-12-20\"", &models.Date{2020, time.December, 20}},
		{"\"2020-12-20", nil},
		{"\"2020-1220\"", nil},
		{"\"aaaa\"", nil},
		{"\"202o-12-20\"", nil},
		{"\"2020-12-32\"", nil},
		{"\"2020-13-20\"", nil},
		{"\"202012-20\"", nil},
		{"\"2006-01-02T15:04:05Z07:00\"", nil},
	}

	for _, test := range cases {
		t.Run(test.date, func(t *testing.T) {
			date := &models.Date{}
			err := date.UnmarshalGQL(test.date)
			if test.expected == nil {
				require.EqualError(t, err, models.ErrMalformedDate.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, date)
			}
		})
	}
}

func TestMonth_MarshalGQL(t *testing.T) {
	cases := []struct {
		month    models.Month
		expected string
	}{
		{models.Month{2020, time.January}, "\"2020-01\""},
		{models.Month{2020, time.December}, "\"2020-12\""},
	}
	for _, test := range cases {
		t.Run(test.expected, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			test.month.MarshalGQL(buffer)
			assert.Equal(t, test.expected, buffer.String())
		})
	}
}

func TestMonth_UnmarshalGQL(t *testing.T) {
	cases := []struct {
		month    string
		expected *models.Month
	}{
		{"\"2020-01\"", &models.Month{2020, time.January}},
		{"\"2020-12\"", &models.Month{2020, time.December}},
		{"\"2020-12", nil},
		{"\"aaaa\"", nil},
		{"\"202o-12\"", nil},
		{"\"2020-13\"", nil},
		{"\"202012\"", nil},
		{"\"2006-01-02T15:04:05Z07:00\"", nil},
	}

	for _, test := range cases {
		t.Run(test.month, func(t *testing.T) {
			date := &models.Month{}
			err := date.UnmarshalGQL(test.month)
			if test.expected == nil {
				require.EqualError(t, err, models.ErrMalformedMonth.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expected, date)
			}
		})
	}
}

func TestMonth_Contains(t *testing.T) {
	cases := []struct {
		title    string
		month    models.Month
		date     models.Date
		expected bool
	}{
		{"Success", models.Month{2020, time.January}, models.Date{2020, time.January, 12}, true},
		{"Wrong month", models.Month{2020, time.January}, models.Date{2020, time.February, 12}, false},
		{"Wrong year", models.Month{2020, time.January}, models.Date{2030, time.January, 12}, false},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.month.Contains(test.date), test.expected)
		})
	}
}

func TestMonth_Next(t *testing.T) {
	cases := []struct {
		title    string
		month    models.Month
		expected models.Month
	}{
		{"December", models.Month{2020, time.December}, models.Month{2021, time.January}},
		{"Other month", models.Month{2020, time.June}, models.Month{2020, time.July}},
	}

	for _, test := range cases {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.month.Next(), test.expected)
		})
	}
}

func TestAmount_Add(t *testing.T) {
	term1 := models.Amount{12, 62}
	term2 := models.Amount{23, 83}
	expected := models.Amount{36, 45}
	assert.Equal(t, expected, term1.Add(term2))
}
