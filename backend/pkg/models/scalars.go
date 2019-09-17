package models

import (
	"errors"
	"io"
	"strconv"
)

type MoneyAmount float32

func (ma *MoneyAmount) UnmarshalGQL(v interface{}) error {
	i, ok := v.(float32)
	if !ok {
		return errors.New("MoneyAmount must be Int")
	}
	*ma = MoneyAmount(i)
	return nil
}

func (ma MoneyAmount) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.FormatFloat(float64(ma), 'f', 2, 32)))
}
