package types

type Color string

const (
	Red     Color = "\033[31m"
	Green   Color = "\033[32m"
	Yellow  Color = "\033[33m"
	Blue    Color = "\033[34m"
	Purple  Color = "\033[35m"
	Cyan    Color = "\033[36m"
	NoColor Color = "\033[0m"
)

type FileStatus uint8

const (
	Added     FileStatus = iota // IF IN INDEX BUT NOT IN OBJECT
	Removed                     // IF NOT IN INDEX BUT IN OBJECT
	Modified                    // IF IN INDEX AND IN OBJECT WITH SAME NAME
	Renamed                     // IF IN INDEX AND IN OBJECT BUT WITH DIFFERENT NAME
	Untracked                   // IF NOT IN INDEX AND NOT IN OBJECT
	Ignored                     // IF NOT IN INDEX AND NOT IN OBJECT AND IGNORED
	Unknown                     // ERROR CASE
)

type FileStatusStruct struct {
	Path   string     `json:"path"`
	Status FileStatus `json:"status"`
}
