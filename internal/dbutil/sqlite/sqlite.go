// Package sqlite implementations operations related to the DevOps and
// the writing and reading of data to and from SQLite
package sqlite

import (
	"errors"
)

var (
	ErrSqliteDevOp = errors.New("sqlite devops error")
)
