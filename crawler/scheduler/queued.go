package scheduler

import "learnCrawler/crawler/engine"

//实现接口Scheduler
type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request //每一个worker有自己的chan（createWorker）
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	s.workerChan = make(chan chan engine.Request)

	go func() {
		//队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				//如果没有匹配到则activeRequest、activeWorker为nil
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				//收到一个就让他在队列排队
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
				//如果没有匹配到则activeRequest、activeWorker为nil;这里绝对不会select到
				//activeRequest送入activeWorker后要把他们从队列拿掉
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}


