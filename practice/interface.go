package practice

type (
	Iterable interface {
		Iterator() Iterator
	}
	Iterator interface {
		Next() (interface{}, error)
	}
	Storage interface {
		Collection() (interface{}, error)
	}
	Collection interface {
		Iterable

		ID() int64
		Name() string
		Type() string

		Save() error
		Load() error

		BucketPartition() (string, error)
	}
	Dataset interface {
		Iterable

		ID() int64
		Name() string
		Type() string

		Save() error
		Load() error
		IsProcessed() bool

		BucketPartition() (string, error)
	}
	DataCollection interface {
		Iterable

		ID() int64
		Name() string
		Type() string

		Save() error
		Load() error
		IsProcessed() bool

		BucketPartition() (string, error)
	}
)
