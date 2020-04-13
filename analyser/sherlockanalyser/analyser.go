package sherlockanalyser

/*
AnalyserServiceHandler will be AnalyserService representation.
*/
type AnalyserServiceHandler struct {
	AnalyserQueue *AnalyserQueue
	Dependencies  *AnalyserDependency
}

/*
NewAnalyserServiceHandler will return an new AnalyserServiceHandler instance.
*/
func NewAnalyserServiceHandler() *AnalyserServiceHandler {
	return &AnalyserServiceHandler{}
}

/*
InjectDependency will inject the dependencies for the a sherlockcrawler instance
into the actual instance.
*/
func (analyser AnalyserServiceHandler) InjectDependency(deps *AnalyserDependency) {
	analyser.Dependencies = deps
}

/*
CreateTask will append the current queue a task.
*/
func (analyser AnalyserServiceHandler) CreateTask() {

}

/*
SendResult will send the result to the crawler.
*/
func (analyser AnalyserServiceHandler) SendResult() {

}
