package models

type Wrapped[T any] struct {
	Data T
}

type WrappedWithMeta[T any] struct {
	Data T
	Meta Metadata
}
