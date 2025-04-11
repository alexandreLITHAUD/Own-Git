package utils

const (
	added    uint8 = iota // IF IN INDEX BUT NOT IN OBJECT
	removed               // IF NOT IN INDEX BUT IN OBJECT
	modified              // IF IN INDEX AND IN OBJECT WITH SAME NAME
	renamed               // IF IN INDEX AND IN OBJECT BUT WITH DIFFERENT NAME
)
