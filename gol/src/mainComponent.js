import React, { Component } from 'react'


class MainComponent extends Component{
    ws = new WebSocket("ws://localhost:8090/ws");
    render(){
        return(
            <div>
                Here we go again
            </div>
        )
    }
}

export default MainComponent