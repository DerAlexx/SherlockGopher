import React from 'react';
import 'bootstrap/dist/css/bootstrap.css';
import SearchBar from './searchbar.js';
import ReactPaginate from 'react-paginate';
import axios from 'axios'
import "./imagemetadataservice.css";

export default class Imagemetadataservice extends React.Component {
    METADATA = "http://0.0.0.0:8081/graph/v1/getmetadata"

    state = {
        postdata: [],
        currentPage: 0,
        maxPage: 1,
        pageRange: 0,
        sdatamessage: "",
        hasSdataError: false
    }

    constructor(props) {
        super(props);
        this.handlePageClick = this
            .handlePageClick
            .bind(this);
    }

    receivedData() {
        axios.get(this.METADATA + "/" + this.state.currentPage).then(res => {
            try {
                const data = res.data.map;
                const postData = data.map(pd => <React.Fragment>

                        <div class="card">
                        <div class="card-body">
                            <ul class="list-group list-group-flush">
                            <li class="list-group-item">Image URL: {pd.img_url}</li>
                            <li class="list-group-item">ID: {pd.img_id}</li>
                            <li class="list-group-item">Zustand: {pd.condition}</li>
                            <li class="list-group-item">Datum: {pd.datetime_original}</li>
                            <li class="list-group-item">Gerät: {pd.model}</li>
                            <li class="list-group-item">Kameraherstellers: {pd.make}</li>
                            <li class="list-group-item">Beschreibung: {pd.maker_note}</li>
                            <li class="list-group-item">Software: {pd.software}</li>
                            <li class="list-group-item">GPS-Längengrad: {pd.gps_latitude}</li>
                            <li class="list-group-item">GPS-Breitengrad: {pd.gps_longitude}</li>
                        </ul>
                        </div>
                        </div>
                        <hr></hr>
                    </React.Fragment>)
                this.setState({
                    sdatamessage: "No error",
                    hasSdataError: false,
                    postData,
                    currentPage: res.data.currentpage,
                    maxPage: res.data.maxpage,
                    pageRange: res.data.pagerange,
                })
            } catch (error) {
                console.log(error)
                this.setState({
                    sdatamessage: "For some Reason an Error occured. Cannot process response.",
                    hasMetaError: true,
                })  
            }
            }).catch(error => {
                console.log(error)
                this.setState({
                    sdatamessage: "For some Reason an Error occured. Is the Webserver up?",
                    hasSdataError: true,
            })
        })
    }

    handlePageClick = (e) => {
        const selectedPage = e.selected;
        this.setState({
            currentPage: selectedPage,
        }, () => {
            this.receivedData()
        });
    };

    componentDidMount(){
        this.interval = setInterval(() => {
            try {
                this.receivedData()
            } catch(error) {
                console.log(error)
            }
        }, 10000)
    }

    componentWillUnmount() {
        clearInterval(this.interval)
    }

    render() {
        const {
            sdatamessage,
            hasSdataError
        } = this.state
        return(
            <div>
                <SearchBar></SearchBar>
                <div class="body">
                    <p>Metadata of all Images</p>
                        { 
                            hasSdataError ? 
                                <div class="alert alert-danger">
                                    {sdatamessage}
                                </div>
                                :
                                <div>
                                    <hr></hr>
                                    <br></br>
                                    {this.state.postData}
                                    <ReactPaginate
                                        previousLabel={"<<"}
                                        nextLabel={">>"}
                                        breakLabel={"..."}
                                        breakClassName={"break-me"}
                                        pageCount={this.state.maxPage}
                                        marginPagesDisplayed={1}
                                        pageRangeDisplayed={this.state.pageRange}
                                        onPageChange={this.handlePageClick}
                                        containerClassName={"pagination"}
                                        subContainerClassName={"pages pagination"}
                                        activeClassName={"active"} 
                                    /> 
                                </div>     
                        }            
                </div>
            </div>
        )
    }
}