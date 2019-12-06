class Universe:

    def __init__(self, graph):
        self.graph = graph
    
    def find_all_orbiters_for_key(self, key):
        orbiters = []
        for node in self.graph[key]:
            orbiters += [node] + self.find_all_orbiters_for_key(node)
        return orbiters
    
    def generate_all_orbiters_for_all_keys(self):
        all_orbiters_graph = {}
        for key in self.graph:
            all_orbiters_graph[key] = self.find_all_orbiters_for_key(key)
        self.all_orbiters_graph = all_orbiters_graph
    
    def count_all_orbiters(self):
        c = 0
        for key in self.all_orbiters_graph:
            c += len(self.all_orbiters_graph[key])
        return c 

# read and parse input
with open('input', 'r') as f:
    data = [(l.strip("\n").split(")")) for l in f.readlines()]
orbit_map={} # key:orbited_by
for node in data:
    if node[0] not in orbit_map:
        orbit_map[node[0]] = [node[1]]
    else:
        orbit_map[node[0]].append(node[1])
    if node[1] not in orbit_map:
        orbit_map[node[1]] = []

# instantiate universe object with graph, generate and count orbiters
universe = Universe(orbit_map)
universe.generate_all_orbiters_for_all_keys()
print(universe.count_all_orbiters())
