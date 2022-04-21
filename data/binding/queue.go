package binding

import (
	"sync"

	"github.com/Jaeel/fyne/internal/async"
)

var (
	once  sync.Once
	queue *async.UnboundedFuncChan
)

func queueItem(f func()) {
	once.Do(func() {
		queue = async.NewUnboundedFuncChan()
		go func() {
			for f := range queue.Out() {
				f()
			}
		}()
	})
	queue.In() <- f
}

func waitForItems() {
	done := make(chan struct{})
	queue.In() <- func() { close(done) }
	<-done
}
