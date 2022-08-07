from storage.storage import Storage
import redis
import os

class RedisStorage(Storage):
    def __init__(self):
        self.r = redis.StrictRedis(host=os.getenv("REDIS_HOST_NAME"), port=os.getenv("REDIS_PORT"), db=os.getenv("REDIS_DB"), password=os.getenv("REDIS_PASS"))

    def insert(self, kvPair):
        res = self.r.set(kvPair.key, kvPair.value)
        return res

    def getAll(self):
        res = []
        for key in self.r.scan_iter("*"):
            res.append((key, self.r.get(key)))
        return res