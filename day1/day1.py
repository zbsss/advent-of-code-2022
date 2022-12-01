import heapq


with open('./input.txt', 'r') as file:
    min_heap = []

    count = 0
    for line in file.readlines():
        if stripped := line.strip():
            count += int(stripped)
        else:
            heapq.heappush(min_heap, count)

            if len(min_heap) > 3:
                heapq.heappop(min_heap)

            count = 0

print(f"task 1: {min_heap[0]}")
print(f"task 2: {sum(min_heap)}")
