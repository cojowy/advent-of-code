def calc_fuel(num):
    return int(num/3) - 2 if int(num/3) - 2 > 0 else 0
numsum = 0
input = [int(l.strip('\n')) for l in open('input', 'r').readlines()]
for n in input:
    num = int(n)
    while num != 0:
        num = calc_fuel(num)
        numsum += num 
print(numsum)