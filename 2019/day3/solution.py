import copy
from matplotlib import pyplot as plt

def fill_in_gaps(points, coords):
    if len(points) > 0:
        startX, startY = points[-1][0], points[-1][1]
    else:
        startX, startY = 0,0
    endX, endY = coords[0], coords[1]
    new_points = []
    for i in range(startX, endX, 1 if endX > startX else -1):
        if (i, startY) != (startX, startY) or (startX, startY) == (0, 0):
            new_points.append((i, startY)) 
    for i in range(startY, endY, 1 if endY > startY else -1):
        if (startX, i) != (startX, startY) or (startX, startY) == (0, 0):
            new_points.append((startX, i))
    new_points.append(copy.deepcopy(coords))
    return new_points

def path_points(path):
    coords = [0,0]
    points = []
    for move in path:
        direction, steps = move[0], int(move[1:])
        if direction == 'R':
            coords[0] += steps
        elif direction == 'L':
            coords[0] -= steps
        elif direction == 'U':
            coords[1] += steps
        elif direction == 'D':
            coords[1] -= steps
        points += fill_in_gaps(points, tuple(coords))
    return points

def mhd(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

with open('input', 'r') as f:
    input = [l.split(',') for l in [s.strip('\n') for s in f.readlines()]]

path1_points = path_points(input[0])
path2_points = path_points(input[1])
common_points = list(set(path1_points).intersection(path2_points))
common_points.remove((0,0))

print("Minimum manhattan distance:",min([mhd([0,0], p) for p in common_points]))
print("Minimum steps taken:", min([path1_points.index(p) + path2_points.index(p) for p in common_points]))