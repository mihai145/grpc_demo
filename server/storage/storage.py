from abc import ABC, abstractmethod

class Storage(ABC):
    @abstractmethod
    def insert(self, kvPair):
        pass

    @abstractmethod
    def getAll(self):
        pass