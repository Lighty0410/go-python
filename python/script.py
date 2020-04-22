import grpc
import logging

import server_pb2_grpc, server_pb2, client_pb2_grpc, client_pb2
from concurrent import futures


def GetTopic(word):
    channel = grpc.insecure_channel('localhost:8080')
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
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()


serve()

if __name__ == '__main__':
    logging.basicConfig()
    serve()