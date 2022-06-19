package common

import (
	"github.com/devfeel/mapper"
	"github.com/pkg/errors"
)

func Mapper(from, to interface{}) error {
	m := mapper.NewMapper()
	m.SetEnabledJsonTag(false)
	return m.Mapper(from, to)
}

func MapperSlice[F, T any](sliceFrom []F, sliceTo []T) error {
	if sliceFrom == nil || sliceTo == nil || len(sliceFrom) != len(sliceTo) {
		return errors.New("can't convert slice")
	}
	for i := 0; i < len(sliceFrom); i++ {
		if err := mapper.Mapper(&(sliceFrom)[i], &sliceTo[i]); err != nil {
			return err
		}
	}
	return nil
}
