def maximum_overlapping_intervals(intervals):
    events = []
    for start, end in intervals:
        events.append((start, 'start'))
        events.append((end, 'end'))
    # Sort events: first by time, then 'end' before 'start' if times are equal
    events.sort(key=lambda x: (x[0], x[1] == 'start'))
    current_overlap = 0
    max_overlap = 0
    for event in events:
        if event[1] == 'start':
            current_overlap += 1
            max_overlap = max(max_overlap, current_overlap)
        elif event[1] == 'end':
            current_overlap -= 1
    return max_overlap
# Example usage:
intervals = [(1, 3), (2, 4), (3, 6), (5, 7), (6, 8)]
result = maximum_overlapping_intervals(intervals)
print("Maximum number of overlapping intervals:", result) # 2
