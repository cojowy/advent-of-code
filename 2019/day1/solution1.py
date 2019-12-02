with open('input', 'r') as f:
    input = [int(int(l.strip('\n'))/3)-2 for l in f.readlines()]
print(sum(input))
