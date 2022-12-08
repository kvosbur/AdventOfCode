from abc import ABC, abstractmethod


class AbstractFile(ABC):
    name: str

    @abstractmethod
    def size(self):
        pass

    @abstractmethod
    def output(self, tabs=0):
        pass


class File(AbstractFile):

    def __init__(self, name, file_size):
        self.name = name
        self.file_size = file_size

    def size(self):
        return self.file_size

    def output(self, tabs=0):
        print(tabs * "\t", "File", self.name, self.size())


class Directory(AbstractFile):

    def __init__(self, name):
        self.name = name
        self.files = {}

    def add_file(self, file: AbstractFile):
        self.files[file.name] = file

    def size(self):
        size = 0
        for file in self.files.values():
            size += file.size()
        return size

    def output(self, tabs=0):
        print(tabs * "\t", "Directory", self.name, self.size())
        for file in self.files.values():
            file.output(tabs + 1)

