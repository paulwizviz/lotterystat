package tball

import "errors"

var (
	ErrDrawDate = errors.New("invalid draw date")
	ErrBall1    = errors.New("invalid ball 1")
	ErrBall2    = errors.New("invalid ball 2")
	ErrBall3    = errors.New("invalid ball 3")
	ErrBall4    = errors.New("invalid ball 4")
	ErrBall5    = errors.New("invalid ball 5")
	ErrTBall    = errors.New("invalid thunder ball")
	ErrSeq      = errors.New("invalid seq")
	ErrRec      = errors.New("invalid record")
)
