package logs

var (
	level = map[string]int{
		"info":    0,
		"warning": 1,
		"error":   2,
	}
)

type Logs struct {
	Level string
	log   string
}

func New(log, lev string) *Logs {
	return &Logs{lev, log}
}

func (l *Logs) Set(log string, lev string) {
	if level[l.Level] < level[lev] {
		l.log = log
		l.Level = lev
	}
}

func (l *Logs) SetError(log string) {
	l.log = log
	l.Level = "error"
}

func (l *Logs) SetWarning(log string) {
	if level[l.Level] < level["warning"] {
		l.log = log
		l.Level = "warning"
	}
}

func (l *Logs) SetInfo(log string) {
	if level[l.Level] == level["info"] {
		l.log = log
		l.Level = "info"
	}
}

func (l *Logs) Error() string {
	if 1 <= level[l.Level] {
		return l.log
	}
	return ""
}

func (l *Logs) Log() string {
	return l.log
}
