import app_pb2
import app_pb2_grpc

import grpc

import os
from dotenv import load_dotenv


class Client(object):
    def __init__(self):
        self.host = 'localhost'
        self.server_port = os.getenv("GRPC_SERVER_PORT")

        # instantiate a channel
        self.channel = grpc.insecure_channel(
            '{}:{}'.format(self.host, self.server_port))

        # bind the client and the server
        self.stub = app_pb2_grpc.DictServiceStub(self.channel)

    def addPair(self, key, value):
        request = app_pb2.KeyValuePair(key=key, value=value)        
        return self.stub.AddPair(request)

    def getPairs(self):
        request = app_pb2.Empty()        
        return self.stub.GetPairs(request)


def printResult(result):
    print("Key value pairs: ")
    for p in result.pairs:
        print(f'{p.key}={p.value}')
    print()


if __name__ == "__main__":
    load_dotenv()
    client = Client()

    client.addPair("k1", "v1")
    result = client.getPairs()
    printResult(result)

    client.addPair("k2", "v2")
    client.addPair("k3", "v3")
    result = client.getPairs()
    printResult(result)