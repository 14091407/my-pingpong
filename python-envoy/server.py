from concurrent import futures
import logging

import grpc

import pingpong_pb2
import pingpong_pb2_grpc

class Pingpong():
    def pingPong(self, request, context):
        ping = request.ping
        return pingpong_pb2.PongResponse(pong="pong!!")

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pingpong_pb2_grpc.add_PingPongServiceServicer_to_server(Pingpong(), server)
    server.add_insecure_port('[::]:50051')
    print("Server start at port: 50051...")
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()