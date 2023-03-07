package scheduler

import "xyy/learngo/crawler/engine"

type QueuedScheduler struct {
}

func (q QueuedScheduler) Submit(request engine.Request) {
	panic("implement me")
}

func (q QueuedScheduler) ConfigureMasterWorkerChan(requests chan engine.Request) {
	panic("implement me")
}
