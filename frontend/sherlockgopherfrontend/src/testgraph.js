//React Standard
import 'bootstrap/dist/css/bootstrap.css';
import React from "react";
import ForceGraph3D from 'react-force-graph-3d';
import ForceGraph2D from 'react-force-graph-2d';
import SpriteText from "three-spritetext";
import Axios from 'axios';

//Javascripts
import SearchBar from './searchbar.js';

//Stylesheets
import './assets/css/App.css';

export default class Testgraph extends React.Component {
    
    METAINFORMATION = "http://0.0.0.0:8081/graph/v1/meta"
    ALLINFORMATIONPERFORMANCE = "http://0.0.0.0:8081/graph/v1/alloptimized"

    state = {
        //Show 2D or 3D
        is2d: true,

        seconds: 60,
        rootid: "",

        //MetaInformation State
        metaError: "",
        metahasError: false,
        showMetaError: false,

        //Meta Data
        metadata: [],

        //Data State
        dataError: "",
        dataHasError: false,
        showDataError: false,

        //Data
        responseData:[],
        data: {
            "nodes":[],
            "links": []
        },
        prunedTree: {
            "nodes":[],
            "links": []
        },
    }

    constructor(props) {
        super(props);
        this.getChildrenOfNode = this.getChildrenOfNode.bind(this);
        this.getRoot = this.getRoot.bind(this);
        this.getMetaInformation = this.getMetaInformation.bind(this);
        this.getGraphData = this.getGraphData.bind(this);
        this.handleChange = this.handleChange.bind(this);
        this.getGraph = this.getGraph.bind(this);
        this.interval = 0;
        this.timerInterval = 0;
    }

    /*
    Fetch all Metainformation like number of nodes, etc.
    */
    getMetaInformation() {
        Axios.get(
            this.METAINFORMATION
        ).then(res => {
            const response = res.data
            this.setState({
                metadata: response,
            })
        }).catch(error => {
            this.setState({
                dataError: "Cannot fetch the Metadata, Error: " + String(error),
                dataHasError: true,
                showDataError: true,
            })
        })
    }

    /*
    getGraphData will fetch the GraphData.
    */
    getGraphData() {
        try {
            Axios.get(
                this.ALLINFORMATIONPERFORMANCE + "/0"
            ).then(res => {
                const response = res.data
                console.log("all", response)
                this.setState({
                        data: response
                })
            }).catch(error => {
                this.setState({
                    metaError: "Cannot fetch the graphdata, Error: " + String(error),
                    metahasError: true,
                    showMetaError: true,
                })
            });
        } catch (error) {
            console.log("An error occured while trying to get the Metadata(all).");
        }
    }

    getRoot(){
        try {
            Axios.get(
                this.ALLINFORMATIONPERFORMANCE + "/1"
            ).then(res => {
                const response = res.data
                this.setState({
                        data: response, 
                })
            }).catch(error => {
                this.setState({
                    metaError: "Cannot fetch the rootgraphdata, Error: " + String(error),
                    metahasError: true,
                    showMetaError: true,
                })
            });
        } catch (error) {
            console.log("An error occured while trying to get the Metadata(root).");
        }
    }


    /*
    componentDidMount will mount the component.
    */
    componentDidMount() {
        this.getRoot()
        this.interval = setTimeout(async () => {
            try {
                this.getMetaInformation()
                if(!this.state.is2d){
                    this.getGraphData()
                }
                console.log("hallo")
            } catch (error) {
                console.log("An error occured while trying to get the Metadata.")
                }
        }, 100)  
        this.interval = setInterval(async () => {
            try {
                this.getMetaInformation()
                if(!this.state.is2d){
                    this.getGraphData()
                }
                this.setState({
                    seconds: 60,
                })
            } catch (error) {
                console.log("An error occured while trying to get the Metadata.")
                }
        }, 60000) 
        this.timerInterval = setInterval(async () => {
            try {
                this.setState(({ seconds }) => ({
                    seconds: seconds - 1
                }))
            } catch (error) {
                console.log("An error occured in the timer function.")
                }
        }, 1000)  
    }

    /*
    componentWillUnmount will unmount the Component.
    */
    componentWillUnmount() {
        clearInterval(this.ConstantInterval)
    }

    /*
    handleChanges will handle a event so switch between 2D and 3D.
    */
    handleChange(event) {   
        this.setState(
          { 
            is2d: !this.state.is2d
          }
        ); 
        if (this.state.is2d){
            this.getRoot()
        } 
    }
    
    /*
    makeMetaDataIterable will format all data in the metadata array in a
    way to perform .map on the new data.
    */
    makeMetaDataIterable(data) {
        try {
            let larray = []
            for (let key in data) {
                for (let entrie in data[key]) {
                    larray.push([entrie, data[key][entrie]])
                }
            }
            return larray
        } catch (error) {
            return []
        }
    }    

    getChildrenOfNode(node, event){  
        try {
            Axios.post(
                this.ALLINFORMATIONPERFORMANCE + "/2",  JSON.stringify({
                    "address": node.id,
                })
            ).then(res => {
                const resp = res.data
                var tmpdata = {
                    "nodes":[],
                    "links": []
                }
                tmpdata.nodes = this.state.data.nodes.concat(resp.nodes)
                tmpdata.links = this.state.data.links.concat(resp.links)
                this.setState({ 
                    data: tmpdata,  
                })

            }).catch(error => {
                this.setState({
                    metaError: "Cannot fetch the childrengraphdata, Error: " + String(error),
                    metahasError: true,
                    showMetaError: true,
                })
            });
        } catch (error) {
                console.log("An error occured while trying to get the Metadata.");
        }  
    }

    getPrunedTree(node){
        const visibleNodes = [];
        const visibleLinks = [];
        (function traverseTree(node = nodesById[rootId]) {
            visibleNodes.push(node);
            if (node.collapsed) return;
            visibleLinks.push(...node.childLinks);
            node.childLinks
              .map(link => ((typeof link.target) === 'object') ? link.target : nodesById[link.target]) // get child node
              .forEach(traverseTree);
        })
    }

    click(node, event){ 
        node.collapsed = !node.collapsed;
        this.getPrunedTree(node)
    }
    
    /*
    getGraph will return the graph or a message.
    */
    getGraph(data) {
        if (data["nodes"] !== undefined && data["links"] !== undefined) {
            return (<ForceGraph2D
            graphData={data}
            nodeLabel="id"
            enableNodeDrag={false}
            width={1145}
            height={600}
            linkDirectionalArrowLength={3.5}
            linkDirectionalArrowRelPos={1}
            linkLabel={"label"}
            showNavInfo={false}
            onNodeClick={this.click}
        />)
        } else {
            return (
                <div class="alert alert-warning">No nodes in the Database or there is no connection</div>
            )
        }                 
    }

    render() {     
        const {data, metaError, metahasError, showMetaError, metadata, seconds} = this.state
        var finalmetadata = this.makeMetaDataIterable(metadata)
        return (
          <div>
            
              <SearchBar></SearchBar>
              <div class="body">
                <p>
                    <p>Metainformation</p>            
                    {showMetaError && metahasError ? 
                        <div class="alert alert-danger">
                            {metaError}
                        </div> 
                        : 
                        <div>
                            { 
                            finalmetadata.map(item => (
                                <span><span class="badge badge-pill badge-secondary">{String(item[0]).toUpperCase()}: {item[1]} </span> <span style={{color: "#d5d8dc", display: "None"}}> | </span>
                                </span>
                                ))
                            }
                        </div>
                        }
                    <hr></hr>
                    <div class="custom-control custom-switch">
                        <input onClick={this.handleChange} type="checkbox" class="custom-control-input" id="customSwitches"></input>
                        <span style={{backgroundColor:"#7CA9EF"}} class="badge badge-pill badge-dark htmlnode"> HTML </span> 
                        <span style={{color: "#d5d8dc"}}>   </span>
                        <span style={{backgroundColor:"#E891BC"}}  class="badge badge-pill badge-success"> Stylesheets </span>
                        <span style={{color: "#d5d8dc"}}>   </span>
                        <span style={{backgroundColor:"#F0B85B"}}  class="badge badge-pill badge-success"> Javascript </span>
                        <span style={{color: "#d5d8dc"}}>   </span>
                        <span style={{backgroundColor:"#85E196"}}  class="badge badge-pill badge-warning"> Images </span>
                        <span style={{color: "#273746"}}> | </span>
                        <span style={{backgroundColor:"#08206A"}}  class="badge badge-pill badge-danger"> Links  &#10138;</span>
                        <span style={{color: "#d5d8dc"}}>   </span>
                        <span style={{backgroundColor:"#24A144"}}  class="badge badge-pill badge-danger"> Requires &#10138;</span>
                        <span style={{color: "#d5d8dc"}}>  </span>
                        <span style={{backgroundColor:"#99BA51"}}  class="badge badge-pill badge-danger"> Shows &#10138;</span>
                        <span style={{color: "#273746"}}> | </span>
                        <span class="badge badge-pill badge-info"> Graph will automatically update in {seconds} sec.</span>
                    </div>     
                    <hr></hr>
                </p>
                    <div style={{margin: "auto",
                                width: "100%",
                                padding: "10px"}}>
                        { 
                            this.getGraph(data)
                        }
                    </div>
              </div>
          </div>
        )
    }
}

