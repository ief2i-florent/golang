package multithread

import (
	"github.com/ief2i-florent/golang/internal/fssync"
)

type Mission struct {
	source string
	target string
}

func (m *Mission) execute() int {
	fssync.Copy(m.source, m.target)

	return 1
}

func MakeMission(source string, target string) *Mission {
	return &Mission{
		source: source,
		target: target,
	}
}
