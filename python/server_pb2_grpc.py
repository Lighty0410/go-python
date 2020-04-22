# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import server_pb2 as server__pb2


class FilterTextStub(object):
    """Missing associated documentation comment in .proto file"""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.FilterText = channel.unary_unary(
                '/grpc.FilterText/FilterText',
                request_serializer=server__pb2.Text.SerializeToString,
                response_deserializer=server__pb2.Answer.FromString,
                )


class FilterTextServicer(object):
    """Missing associated documentation comment in .proto file"""

    def FilterText(self, request, context):
        """Missing associated documentation comment in .proto file"""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_FilterTextServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'FilterText': grpc.unary_unary_rpc_method_handler(
                    servicer.FilterText,
                    request_deserializer=server__pb2.Text.FromString,
                    response_serializer=server__pb2.Answer.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'grpc.FilterText', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class FilterText(object):
    """Missing associated documentation comment in .proto file"""

    @staticmethod
    def FilterText(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/grpc.FilterText/FilterText',
            server__pb2.Text.SerializeToString,
            server__pb2.Answer.FromString,
            options, channel_credentials,
            call_credentials, compression, wait_for_ready, timeout, metadata)