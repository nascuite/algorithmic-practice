#Depth First Search (LIFO Стек)
Проверить существование пути из одной вершины в другую или определить, является ли неориентированный граф связным.
```
Algorithm DFS(G, v)
    if v is already visited
        return
    Mark v as visited.
    // Perform some operation on v.
    for all neighbors x of v
        DFS(G, x)
```

#Breadth First Search (FIFO queue)
Найти кратчайший путь в невзвешенном графе.
```
Algorithm BFS(G, v)
    Q ← new empty FIFO queue
    Mark v as visited.
    Q.enqueue(v)
    while Q is not empty
        a ← Q.dequeue()
        // Perform some operation on a.
        for all unvisited neighbors x of a
            Mark x as visited.
            Q.enqueue(x)
```