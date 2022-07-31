package EnumFlag

import "sync/atomic"

type EnumFlag[T ~uint64] interface {
	Store(value uint64)
	Load() uint64

	HasFlag(flag T) bool
	Remove(flag T)
	Combine(flag T)
}

type enumFlag[T ~uint64] struct {
	flagVal *uint64
}

func NewEnumFlag[T ~uint64]() EnumFlag[T] {
	var defaultVal uint64 = 0
	return &enumFlag[T]{
		flagVal: &defaultVal,
	}
}

func (a *enumFlag[T]) HasFlag(flag T) bool {
	var left uint64 = atomic.LoadUint64(a.flagVal)
	var right uint64 = uint64(flag)

	return (left & right) != 0
}

func (a *enumFlag[T]) Remove(flag T) {
	var left uint64 = atomic.LoadUint64(a.flagVal)
	var right uint64 = uint64(flag)

	//a.flagVal = T(left & ^right)
	atomic.StoreUint64(a.flagVal, left & ^right)
}

func (a *enumFlag[T]) Combine(flag T) {
	var left uint64 = atomic.LoadUint64(a.flagVal)
	var right uint64 = uint64(flag)

	//a.flagVal = T(left | right)
	atomic.StoreUint64(a.flagVal, left|right)
}

func (a *enumFlag[T]) Store(value uint64) {
	//a.flagVal = value
	atomic.StoreUint64(a.flagVal, value)
}

func (a *enumFlag[T]) Load() uint64 {
	//return uint64(a.flagVal)
	return atomic.LoadUint64(a.flagVal)
}
