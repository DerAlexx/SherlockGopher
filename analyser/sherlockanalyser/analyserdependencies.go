package sherlockanalyser

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	crawlerproto "github.com/ob-algdatii-20ss/SherlockGopher/sherlockcrawler/proto/crawlertoanalyser"
)

/*
AnalyserDependency represents all services which are needed to run the analyser.
*/
type AnalyserDependency struct {
	Neo4J   *neo4j.Session
	Crawler func() crawlerproto.AnalyserInterfaceService
}

/*
NewAnalyserDependencies will return a new analyserDependency instance to put it in the dependencies
in a analyser object.
*/
func NewAnalyserDependencies() *AnalyserDependency {
	return &AnalyserDependency{}
}
