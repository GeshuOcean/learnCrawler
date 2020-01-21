package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifiler
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifiler interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	//开启一个goroutine生成requestChan,workerChan chan（这里scheduler其实就是一个goroutine）
	e.Scheduler.Run()

	//生成worker等待
	for i := 0; i < e.WorkerCount; i++ {
		//这里传 e.Scheduler Scheduler继集成了ReadyNotifiler
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//requestChan添加Request
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() {
				e.ItemChan <- item
			}()
		}

		for _, request := range result.Request {
			if isDuplicate(request.Url){
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)
func isDuplicate(url string) bool{
	if visitedUrl[url]{
		return true
	}
	visitedUrl[url]=true
	return false
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifiler) {
	go func() {
		for {
			//请求request chan添加到s.workerChan
			//每一个Scheduler都有自己的Request Chan
			ready.WorkerReady(in)

			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
