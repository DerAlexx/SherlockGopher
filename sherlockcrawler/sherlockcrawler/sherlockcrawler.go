package sherlockcrawler

import (
	"context"
	"fmt"
	"sync"
	"time"

	proto "github.com/ob-algdatii-20ss/SherlockGopher/sherlockcrawler/proto/crawlertoanalyser"
	"github.com/pkg/errors"
)

//Time to wait in milliseconds after checking for the tasks again.
const delaytime = 10

/*
Sherlockcrawler will be the Crawlerservice.
*/
type Sherlockcrawler struct {
	Queue            CrawlerQueue //Queue with all tasks
	Dependencies     *Sherlockdependencies
	SherlockStreamer *SherlockStreamingServer
}

/*
NewSherlockCrawlerService will return a new sherlockcrawler instance.
*/
func NewSherlockCrawlerService() Sherlockcrawler {
	return Sherlockcrawler{
		Queue: NewCrawlerQueue(),
	}
}

/*
InjectDependency will inject the dependencies for the a sherlockcrawler instance
into the actual instance.
*/
func (sherlock *Sherlockcrawler) InjectDependency(deps *Sherlockdependencies) {
	sherlock.Dependencies = deps
}

/*
getQueue will return a pointer to the queue.
*/
func (sherlock *Sherlockcrawler) getQueue() *CrawlerQueue {
	return &sherlock.Queue
}

/*
getSherlockStreamer will return the sherlockstreamer service of the current sherlockcrawler.
*/
func (sherlock *Sherlockcrawler) getSherlockStreamer() *SherlockStreamingServer {
	return sherlock.SherlockStreamer
}

/*
SetSherlockStreamer will set the sherlockstreamer service of the current sherlockcrawler.
*/
func (sherlock *Sherlockcrawler) SetSherlockStreamer(server *SherlockStreamingServer) {
	sherlock.SherlockStreamer = server
}

/*
CreateTask will append the current queue with a task.
*/
func (sherlock Sherlockcrawler) CreateTask(ctx context.Context, in *proto.CrawlTaskCreateRequest, out *proto.CrawlTaskCreateResponse) error {
	task := NewTask()
	task.setAddr(in.GetUrl())
	if id := (*sherlock.getQueue()).AppendQueue(&task); id > 0 {
		out.Statuscode = proto.URL_STATUS_ok
		return nil
	}
	out.Statuscode = proto.URL_STATUS_failure
	return errors.New("cannot create a task for this queue")
}

/*
manageUndoneTasks will manage all undone tasks and start them.
*/
func (sherlock *Sherlockcrawler) manageUndoneTasks(waitgroup *sync.WaitGroup) {
	var localwaitgroup sync.WaitGroup
	for _, v := range *sherlock.getQueue().getThisQueue() {
		if v.getTaskState() == UNDONE {
			go v.MakeRequestAndStoreResponse(&localwaitgroup)
			localwaitgroup.Add(1)
		}
	}
	waitgroup.Done()
}

/*
manageFinishedTasks will managed all finished tasks.
*/
func (sherlock *Sherlockcrawler) manageFinishedTasks(waitgroup *sync.WaitGroup) {
	for k, v := range *sherlock.getQueue().getThisQueue() {
		if v.getTaskState() == FINISHED {
			if append := sherlock.getSherlockStreamer().getQueue().AppendQueue(v); append > 0 {
				sherlock.getQueue().RemoveFromQueue(k)
			}
		}
	}
	waitgroup.Done()
}

/*
manageFailedTasks manage all tasks which failed.
*/
func (sherlock *Sherlockcrawler) manageFailedTasks(waitgroup *sync.WaitGroup) {
	var localwaitgroup sync.WaitGroup
	for k, v := range *sherlock.getQueue().getThisQueue() {
		if v.getTaskState() == FINISHED {
			if v.getTrysError() < 3 {
				go v.MakeRequestAndStoreResponse(&localwaitgroup)
				localwaitgroup.Add(1)
			} else {
				if append := sherlock.getSherlockStreamer().getQueue().AppendQueue(v); append > 0 {
					sherlock.getQueue().RemoveFromQueue(k)
				}
			}
		}
	}
	waitgroup.Done()
}

/*
ManageTasks will be a function to check all tasks in a period of time.
Should run as a go routine. Will work like a cronJob.
*/
func (sherlock *Sherlockcrawler) ManageTasks() {
	for {
		var waitgroup sync.WaitGroup
		go sherlock.manageUndoneTasks(&waitgroup)
		go sherlock.manageFinishedTasks(&waitgroup)
		go sherlock.manageFailedTasks(&waitgroup)
		waitgroup.Wait()
		fmt.Println(sherlock.getQueue().GetStatusOfQueue())
		time.Tick(delaytime * time.Millisecond)
	}
}
