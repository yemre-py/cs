from ds import LinkedList

if __name__ == "__main__":
    ll = LinkedList()
    for i in range(10):
        ll.append(i)
    
    ll.insert(5, 8)
    val = ll.indexOf(0)
    print(val)
    ll.display()
