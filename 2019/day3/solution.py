import copy
from matplotlib import pyplot as plt

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
        points.append(copy.deepcopy(coords))
    return points

with open('input', 'r') as f:
    input = [l.split(',') for l in [s.strip('\n') for s in f.readlines()]]

path1_points = path_points(input[0])
p1x = [p[0] for p in path1_points]
p1y = [p[1] for p in path1_points]

path2_points = path_points(input[1])
p2x = [p[0] for p in path2_points]
p2y = [p[1] for p in path2_points]

plt.plot(p1x, p1y)
plt.plot(p2x, p2y)
plt.show() # calculate distance on graph - point (527, -100), manhattan dist = 627