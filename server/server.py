import app_pb2
import app_pb2_grpc

import grpc
from concurrent import futures

import os
from dotenv import load_dotenv

# from storage.ramStorage import RamStorage
from storage.redisStorage import RedisStorage

class Server(app_pb2_grpc.DictServiceServicer):
    def __init__(self):
        self.storage = RedisStorage()

    def AddPair(self, request, context):
        print(f'New insert request: {request}')

        result = self.storage.insert(request)

        res = app_pb2.StatusResponse(ok=result)
        return res

    def GetPairs(self, request, context):
        print(f'New retrieve request: {request}')
        
        kvs = self.storage.getAll() # list of tuples
        l = []                      # list of app_pb2.KeyValuePair
        for kv in kvs:
            l.append(app_pb2.KeyValuePair(key=kv[0], value=kv[1]))

        res = app_pb2.ListPairs()
        res.pairs.extend(l)

        return res


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    app_pb2_grpc.add_DictServiceServicer_to_server(Server(), server)
    server.add_insecure_port(f'[::]:{os.getenv("GRPC_SERVER_PORT")}')
    
    server.start()
    print("Server Started...")
    
    server.wait_for_termination()
    


if __name__ == "__main__":
    load_dotenv()
    serve()