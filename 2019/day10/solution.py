def points_on_same_line(a, b, c):
    return abs((a[0]*b[1] + b[0]*c[1] + c[0]*a[1] - a[1]*b[0] - b[1]*c[0] - c[1]*a[0])/2) == 0

def distance_between(a, b):
    return((b[0]-a[0])**2 +  (b[1]-a[1])**2)**0.5

def asteroid_positions(data):
    positions=[]
    for i in range(len(data[0])):
        for j in range(len(data)):
            if data[i][j] == '#':
                positions.append((j,i))
    return positions

def third_point_falls_between(a, b, c):
    return round(distance_between(a,c) + distance_between(b,c), 10) == round(distance_between(a,b),10)

def third_point_blocks(a, b, c):
    return points_on_same_line(a, b, c) and distance_between(a, c) < distance_between(a, b) and third_point_falls_between(a, b, c)

with open('input', 'r') as f:
    data = [[char for char in line.strip("\n")] for line in f.readlines()]

positions = asteroid_positions(data)

visibles_map={}
for pt in positions:
    visibles_map[pt] = []
    for p1 in list(set(positions) - set([pt])):
        visible = True
        for p2 in list(set(positions) - set([pt, p1])):
            if third_point_blocks(pt, p1, p2):
                visible = False

        if visible:
            visibles_map[pt].append(p1)

max_key, max_count = (0,0), 0
for key in visibles_map:
    if len(visibles_map[key]) > max_count:
        max_key = key
        max_count = len(visibles_map[key])
print(max_key, max_count)