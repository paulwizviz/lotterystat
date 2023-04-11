// Package dbutil implements a series of interfaces to facilitate
// interactions with databases
package dbutil

import "github.com/paulwizviz/lotterystat/internal/csvdata"

type CSVDataType interface {
	csvdata.EuroDraw | csvdata.Set4LifeDraw
}

type StoreCSVData[T CSVDataType] interface {
	Store(T) error
}

type StoreCSVDataFunc[T CSVDataType] func(T) error

func (s StoreCSVDataFunc[T]) Store(data T) error {
	return s(data)
}
