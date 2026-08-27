package models

type SessionState uint

const (
	SessionStateTitleScreen SessionState = iota
	SessionStateDemo
)
