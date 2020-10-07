from __future__ import print_function
import logging

import grpc

import pingpong_pb2
import pingpong_pb2_grpc

def letPing():
    with grpc.insecure_channel('localhost:10000') as channel:
        stub = pingpong_pb2_grpc.PingPongServiceStub(channel)
        response = stub.pingPong(pingpong_pb2.PingRequest(
            ping='PING!!!',
        ))
        print(response.pong)

if __name__ == '__main__':
    logging.basicConfig()
    letPing()