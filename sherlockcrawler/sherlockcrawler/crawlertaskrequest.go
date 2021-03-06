package sherlockcrawler

import (
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

/*
taskState will be a type representing the current taskState of the task.
*/
type taskState int

const (
	//UNDONE will be an untouched task .
	UNDONE taskState = 0
	//PROCESSING will be a task currently working on.
	PROCESSING taskState = 1
	//FINISHED will be a task which is successfully completed.
	FINISHED taskState = 2
	//FAILED is a task which was in the state of PROCESSING but failed to complet.
	FAILED taskState = 3
)

/*
CrawlerTaskRequest will be a request made by the analyser.
*/
type CrawlerTaskRequest struct {
	taskID            uint64 //taskID, send every time.
	addr              string //addr, once
	taskState         taskState
	taskError         error //error, send as string in case there is an error then dont send a body
	taskErrorTry      int   //never
	response          *http.Response
	responseHeader    *http.Header //header, once (typ map)
	responseBody      string
	responseBodyBytes []byte        //body, split
	statusCode        int           //statusCode, once
	responseTime      time.Duration //response time, once
}

func NewCrawlerTaskRequest(taskID uint64, addr string, taskState taskState, taskError error, taskErrorTry int, response *http.Response, responseHeader *http.Header, responseBody string, responseBodyBytes []byte, statusCode int, responseTime time.Duration) *CrawlerTaskRequest {
	return &CrawlerTaskRequest{taskID: taskID, addr: addr, taskState: taskState, taskError: taskError, taskErrorTry: taskErrorTry, response: response, responseHeader: responseHeader, responseBody: responseBody, responseBodyBytes: responseBodyBytes, statusCode: statusCode, responseTime: responseTime}
}

/*
NewTask will return an empty CrawlerTaskRequest.
*/
func NewTask() CrawlerTaskRequest {
	return CrawlerTaskRequest{
		taskState:    UNDONE,
		responseTime: time.Duration(0),
	}
}

/*
GetTaskID will return the id of a given task.
*/
func (sherlock *CrawlerTaskRequest) GetTaskID() uint64 {
	return sherlock.taskID
}

/*
GetAddr getter for the address.
*/
func (sherlock *CrawlerTaskRequest) GetAddr() string {
	return sherlock.addr
}

/*
GetTaskState will return the taskState of the task.
*/
//nolint: golint
func (sherlock *CrawlerTaskRequest) GetTaskState() taskState {
	return sherlock.taskState
}

/*
GetResponse will return the response of a crawlerTask.
*/
func (sherlock *CrawlerTaskRequest) GetResponse() http.Response {
	return *(sherlock.response)
}

/*
GetResponseHeader will return the Header of the Response.
*/
func (sherlock *CrawlerTaskRequest) GetResponseHeader() http.Header {
	return *(sherlock.responseHeader)
}

/*
GetResponseBody will return the Header of the Response.
*/
func (sherlock *CrawlerTaskRequest) GetResponseBody() string {
	return sherlock.responseBody
}

/*
GetResponseByReference will return the response of a crawlerTask.
*/
func (sherlock *CrawlerTaskRequest) GetResponseByReference() *http.Response {
	return sherlock.response
}

/*
GetResponseBodyInBytes will return the responseBody as a bytearray.
*/
func (sherlock *CrawlerTaskRequest) GetResponseBodyInBytes() []byte {
	return sherlock.responseBodyBytes
}

/*
getTask will return an error which was caused by the http package.
*/
func (sherlock *CrawlerTaskRequest) GetTaskError() error {
	return sherlock.taskError
}

/*
GetStatusCode will return the statusCode.
*/
func (sherlock *CrawlerTaskRequest) GetStatusCode() int {
	return sherlock.statusCode
}

/*
GetTryError will return the number of errors occurred.
*/
func (sherlock *CrawlerTaskRequest) GetTryError() int {
	return sherlock.taskErrorTry
}

/*
GetResponseTime will return the time it took to make the response and get an answer.
*/
func (sherlock *CrawlerTaskRequest) GetResponseTime() time.Duration {
	return sherlock.responseTime
}

/*
setTaskID will set the task id of a given task.
*/
func (sherlock *CrawlerTaskRequest) setTaskID(lid uint64) {
	sherlock.taskID = lid
}

/*
setAddr will set the addr to a given CrawlerTaskRequest.
*/
func (sherlock *CrawlerTaskRequest) setAddr(addr string) {
	sherlock.addr = addr
}

/*
setResponse will set the response of the Request to a given CrawlerTaskRequest.
*/
func (sherlock *CrawlerTaskRequest) setTaskState(taskState taskState) {
	sherlock.taskState = taskState
}

/*
setResponse will set the response of the Request to a given CrawlerTaskRequest.
*/
func (sherlock *CrawlerTaskRequest) setResponse(response *http.Response) {
	sherlock.response = response
}

/*
setResponseBody will set the responseBody of the Response to a given CrawlerTaskRequest.
*/
func (sherlock *CrawlerTaskRequest) setResponseBody(body string) {
	sherlock.responseBody = body
}

/*
setResponseHeader will set the responseHeader of the Response to a given CrawlerTaskRequest.
*/
func (sherlock *CrawlerTaskRequest) setResponseHeader(header *http.Header) {
	sherlock.responseHeader = header
}

/*
setResponseBodyInBytes will set the responseBody in bytes of the Response to a given CrawlerTaskRequest.
*/
func (sherlock *CrawlerTaskRequest) setResponseBodyInBytes(body []byte) {
	sherlock.responseBodyBytes = body
}

/*
setTaskError will set the error raised by the http package.
*/
func (sherlock *CrawlerTaskRequest) setTaskError(error error) {
	sherlock.taskError = error
}

/*
setStatusCode set the statusCode of the task.
*/
func (sherlock *CrawlerTaskRequest) setStatusCode(statusCode int) {
	sherlock.statusCode = statusCode
}

/*
setTryIfError will be the setter for the number of missed tries.
*/
func (sherlock *CrawlerTaskRequest) setTryIfError(numErrors int) {
	sherlock.taskErrorTry = numErrors
}

/*
setResponseTime will set the time it took to make the request and receive an answer.
*/
func (sherlock *CrawlerTaskRequest) setResponseTime(time time.Duration) {
	sherlock.responseTime = time
}

/*
onError will set all things needed if an error occurred while dealing with a task or ŕequest/response.
*/
func (sherlock *CrawlerTaskRequest) onError(err error) {
	sherlock.setTaskError(err)
	sherlock.setTryIfError(sherlock.GetTryError() + 1)
	sherlock.setTaskState(FAILED)
}

/*
MakeRequestAndStoreResponse will make a request and store the result in the field response of the task.
*/
//nolint: gosimple
func (sherlock *CrawlerTaskRequest) MakeRequestAndStoreResponse(waitGroup *sync.WaitGroup) bool {
	sherlock.setTaskState(PROCESSING)
	log.WithFields(log.Fields{
		"addr": sherlock.addr,
	}).Info("MakeRequestAndStoreResponse Start")
	if sherlock.addr == "" {
		errMessage := "cannot process a task with an empty address field"
		log.Error(errMessage)
		sherlock.onError(errors.New(errMessage))
		sherlock.setTaskState(FAILED)
		sherlock.taskErrorTry++
		sherlock.SetWaitGroupDone(waitGroup)
		return false
	}
	startTime := time.Now()

	client := &http.Client{}
	response, err := client.Get(sherlock.addr)

	if err != nil {
		sherlock.taskError = err
		sherlock.setTaskState(FAILED)
		sherlock.taskErrorTry++

		sherlock.SetWaitGroupDone(waitGroup)
		return false
	}
	t := time.Now()
	respT := t.Sub(startTime)
	sherlock.setResponseTime(respT)
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		sherlock.onError(err)
		sherlock.SetWaitGroupDone(waitGroup)
		return false
	}

	sherlock.setResponse(response)
	sherlock.setResponseHeader(&response.Header)
	sherlock.setResponseBody(string(bodyBytes))
	sherlock.setResponseBodyInBytes(bodyBytes)
	sherlock.setStatusCode(response.StatusCode)
	sherlock.setTaskError(nil)
	sherlock.setTaskState(FINISHED)

	log.WithFields(log.Fields{
		"addr": sherlock.addr,
	}).Info("MakeRequestAndStoreResponse End")

	sherlock.SetWaitGroupDone(waitGroup)
	return true
}

func (sherlock *CrawlerTaskRequest) SetWaitGroupDone(waitGroup *sync.WaitGroup) {
	if waitGroup != nil {
		defer waitGroup.Done()
	}
}
