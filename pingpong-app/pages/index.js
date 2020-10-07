import React from 'react'
import { PingPongServiceClient } from '../api/pingpong_grpc_web_pb'
import { PingRequest, PongResponse } from '../api/pingpong_pb'

var client = new PingPongServiceClient('http://localhost:10000');
var client2 = new EntExamInfoServiceClient('http://localhost:10000')

const Index = () => {
    const pingpong = () => {
        const request = new PingRequest()
        request.setPing('Ping')

        client.pingPong(request, {}, (err, response) => {
            if (response == null) {
                console.log(err)
            } else {
                console.log(response.getPong())
            }
        })
    }

    return(
        <div>
            <h1>Call gRPC PingPong Service</h1>
            <p>A response text will appear in browser console.</p>
            <button onClick={pingpong}>Send Ping!!</button>
        </div>
    )
}

export default Index
