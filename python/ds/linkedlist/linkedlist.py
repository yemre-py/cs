from .node import Node

class LinkedList:
    def __init__(self):
        self.head = None
    
    def display(self) -> None:
        """
        Display the linked list
        """
        if self.head is None:
            raise ValueError("cannot found list")

        temp = self.head
        while temp is not None:
            if temp.next is None:
                print(f"{temp.data}\n", end="")
            else: 
                print(f"{temp.data} -> ", end="")
            temp = temp.next
            
    def append(self, val: int) -> None:
        """
        Append a new node to the linked list
        
        Params:
            val (int): The value to be appended
        """
        newNode = Node(val)
        
        if not self.head:
            self.head = newNode
            return
        
        temp = self.head
        while temp.next is not None:
            temp = temp.next
            
        temp.next = newNode
    
    def insert(self, val, pos: int):
        if pos < 0:
            raise ValueError("index out of range")

        newNode = Node(val)
        if pos == 0:
            newNode.next = self.head
            self.head = newNode
            return
        
        index = 0
        temp = self.head
        while temp is not None and index < pos-1:
            temp = temp.next
            index += 1
        
        if temp is None:
            raise ValueError("index out of range")
        
        newNode.next = temp.next
        temp.next = newNode 
    
    
    def indexOf(self, pos) -> int:
        if self.head is None:
            raise ValueError("cannot found list")  
        
        index = 0
        temp = self.head
        while temp is not None:
            if pos is index:
                break
            temp = temp.next
            index += 1

        if temp is None:
            raise ValueError("index out of range")
        
        return temp.data