package scheduler

import "xyy/learngo/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workChan <- w
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workChan = make(chan chan engine.Request)
	go func() {
		/* define two queues to contain these channel elements
		   and then consider how to send r to a specific worker & how to send a specific (next_)request to worker*/
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
