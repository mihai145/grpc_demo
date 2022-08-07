from storage.storage import Storage

class RamStorage(Storage):
    def __init__(self):
        self.dict = dict()

    def insert(self, kvPair):
        self.dict[kvPair.key] = kvPair.value
        return True

    def getAll(self):
        return self.dict.items()