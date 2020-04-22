import grpc
import logging
import os

import server_pb2_grpc, server_pb2, client_pb2_grpc, client_pb2
from concurrent import futures

client_port = os.getenv('GRPC_PYTHON_CLIENT_PORT')
if not client_port:
    client_port = "localhost:8080"

server_port = os.getenv('GRPC_PYTHON_SERVER_PORT')
if not server_port:
        server_port = "[::]:50051"

channel = grpc.insecure_channel(client_port)

def GetTopic(word):
    stub = client_pb2_grpc.GetTopicsStub(channel)
    wor = client_pb2.Word()
    wor.Word = word
    response = stub.GetTopics(wor)
    return response.topics

class Filter(server_pb2_grpc.FilterTextServicer):
    def FilterText(self, request, context):
        topics = request.Text.split()
        topics_list = []
        for topic in topics:
            top = GetTopic(topic)
            topics_list.append(top)
        answer = server_pb2.Answer()
        for t in topics_list:
            for topic in t:
                if answer.result.get(topic.key, None) is None:
                    answer.result[topic.key] = topic.value
                    continue
                if answer.result.get(topic.key, None) < topic.value:
                    answer.result[topic.key] = topic.value
        return answer


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=8))
    server_pb2_grpc.add_FilterTextServicer_to_server(Filter(), server)
    server.add_insecure_port(server_port)
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    logging.basicConfig()
    serve()