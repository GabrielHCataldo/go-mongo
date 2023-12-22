package option

import "time"

type Exists struct {
	// Specifies a collation to use for string comparisons during the operation. This option is only valid for MongoDB
	// versions >= 3.4. For previous server versions, the driver will return an error if this option is used. The
	// default value is nil, which means the default collation of the collection will be used.
	Collation *Collation
	// A string or document that will be included in server logs, profiling logs, and currentOp queries to help trace
	// the operation.  The default is nil, which means that no comment will be included in the logs.
	Comment *string
	// The index to use for the aggregation. This should either be the index name as a string or the index specification
	// as a document. The driver will return an error if the hint parameter is a multi-key map. The default value is nil,
	// which means that no hint will be sent.
	Hint any
	// The maximum amount of time that the query can run on the server. The default value is nil, meaning that there is
	// no time limit for query execution.
	//
	// NOTE: MaxTime will be deprecated in a future release. The more general Timeout option may be used in
	// its place to control the amount of time that a single operation can run before returning an error. MaxTime is
	// ignored if Timeout is set on the client.
	MaxTime *time.Duration
}

func NewExists() Exists {
	return Exists{}
}

func (e Exists) SetCollation(collation *Collation) Exists {
	e.Collation = collation
	return e
}

func (e Exists) SetComment(comment string) Exists {
	e.Comment = &comment
	return e
}

func (e Exists) SetHint(a any) Exists {
	e.Hint = a
	return e
}

func (e Exists) SetMaxTime(d time.Duration) Exists {
	e.MaxTime = &d
	return e
}

func GetExistsOptionByParams(opts []Exists) Exists {
	result := Exists{}
	for _, opt := range opts {
		if opt.Collation != nil {
			result.Collation = opt.Collation
		}
		if opt.Comment != nil {
			result.Comment = opt.Comment
		}
		if opt.Hint != nil {
			result.Hint = opt.Hint
		}
		if opt.MaxTime != nil {
			result.MaxTime = opt.MaxTime
		}
	}
	return result
}