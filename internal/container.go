package internal

import (
	"os"
	"sync"
	"time"
)

type Container struct {
	id                   string
	stateDir             string
	initProcessStartTime uint64
	m                    sync.Mutex
	criuVersion          int
	created              time.Time
	fifo                 *os.File
}
