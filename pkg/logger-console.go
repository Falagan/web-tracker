package pkg

import (
	"fmt"
	"sync"
)

type ConsoleObserver struct {
	mu sync.Mutex
}

func NewConsoleObserver() Observer {
	return &ConsoleObserver{}
}

func (o *ConsoleObserver) Log(level LogLevel, msg string) {
	o.mu.Lock()
	defer o.mu.Unlock()
	logLine := fmt.Sprintf("[%s] %s", level.name, msg)
	fmt.Println(logLine)
}
