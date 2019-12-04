def two_adjacent_equal_numbers(num):
    numstring = str(num)
    for i in range(len(numstring) - 1):
        if numstring[i] == numstring[i+1]:
            return True
    return False

def never_decreases(num):
    numstring = str(num)
    for i in range(len(numstring) - 1):
        if numstring[i+1] < numstring[i]:
            return False
    return True

def check_surroundings(num, i):
    if i > 0 and num[i] == num[i-1]:
        return False
    if i < len(num)-2 and num[i] == num[i+2]:
        return False
    return True

def exactly_two_adjacent_equal_numbers(num):
    numstring = str(num)
    for i in range(len(numstring) - 1):
        if numstring[i] == numstring[i+1] and check_surroundings(numstring, i):
            return True
    return False

num_range = (353096, 843212)

c = 0
for num in range(num_range[0],num_range[1]+1):
    if two_adjacent_equal_numbers(num) and never_decreases(num):
        c += 1
print("part 1:", c)

c = 0
for num in range(num_range[0],num_range[1]+1):
    if exactly_two_adjacent_equal_numbers(num) and never_decreases(num):
        c += 1
print("part 2:", c)