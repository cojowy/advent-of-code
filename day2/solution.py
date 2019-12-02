import copy

def compute(myinput):
    i = 0
    operation = myinput[i]
    while i < len(myinput) and operation != 99:
        val1, val2, pos = myinput[myinput[i+1]], myinput[myinput[i+2]], myinput[i+3]
        if operation == 1:
            myinput[pos] = val1 + val2
        elif operation == 2:
            myinput[pos] = val1*val2
        else:
            raise 'error'
        i += 4
        operation = myinput[i]
    return myinput

with open('input', 'r') as f:
    startinput = [int(i) for i in f.read().split(",")]

# startinput[1], startinput[2] = 12, 2 # part 1 overrides
# print(compute(myinput))

for noun in range(99):
    for verb in range(99):
        myinput = copy.deepcopy(startinput)
        myinput[1], myinput[2] = noun, verb
        candidate_result = compute(myinput)
        if candidate_result[0] == 19690720:
            print("noun:", noun)
            print("verb:", verb)
            print("Answer = 100 * noun + verb =", 100*noun+verb)