import itertools
class Computer:
    def __init__(self, data):
        self.data = data
        self.input_counter = 0

    def parse_operation(self, operation):
        in_op = [int(char) for char in str(operation)]
        while len(in_op) < 5:
            in_op = [0] + in_op
        return in_op

    def compute(self, inputs=[]):
        i, operation_params = 0, self.parse_operation(self.data[0])
        while i < len(self.data) and operation_params[3] != 9 and operation_params[4] != 9:
            if operation_params[3] == 0 and operation_params[4] == 1: # add
                val1 = self.data[i+1] if operation_params[2] == 1 else self.data[self.data[i+1]]
                val2 = self.data[i+2] if operation_params[1] == 1 else self.data[self.data[i+2]]
                position = self.data[i+3]
                self.data[position] = val1 + val2
                i += 4
            elif operation_params[3] == 0 and operation_params[4] == 2: # multiply
                val1 = self.data[i+1] if operation_params[2] == 1 else self.data[self.data[i+1]]
                val2 = self.data[i+2] if operation_params[1] == 1 else self.data[self.data[i+2]]
                position = self.data[i+3]
                self.data[position] = val1 * val2
                i += 4
            elif operation_params[3] == 0 and operation_params[4] == 3: # write input to position
                position = self.data[i+1]
                self.data[position] = inputs[self.input_counter]
                self.input_counter+=1
                i += 2
            elif operation_params[3] == 0 and operation_params[4] == 4: # print value at position
                return self.data[self.data[i+1]]
            elif operation_params[3] == 0 and operation_params[4] == 5: # jump if true
                val1 = self.data[i+1] if operation_params[2] == 1 else self.data[self.data[i+1]]
                val2 = self.data[i+2] if operation_params[1] == 1 else self.data[self.data[i+2]]
                if val1 != 0:
                    i = val2
                else:
                    i += 3
            elif operation_params[3] == 0 and operation_params[4] == 6: # jump if false
                val1 = self.data[i+1] if operation_params[2] == 1 else self.data[self.data[i+1]]
                val2 = self.data[i+2] if operation_params[1] == 1 else self.data[self.data[i+2]]
                if val1 == 0:
                    i = val2
                else:
                    i += 3
            elif operation_params[3] == 0 and operation_params[4] == 7: # less than
                val1 = self.data[i+1] if operation_params[2] == 1 else self.data[self.data[i+1]]
                val2 = self.data[i+2] if operation_params[1] == 1 else self.data[self.data[i+2]]
                position = self.data[i+3]
                if val1 < val2:
                    self.data[position] = 1
                else:
                    self.data[position] = 0
                i += 4
            elif operation_params[3] == 0 and operation_params[4] == 8: # equal to
                val1 = self.data[i+1] if operation_params[2] == 1 else self.data[self.data[i+1]]
                val2 = self.data[i+2] if operation_params[1] == 1 else self.data[self.data[i+2]]
                position = self.data[i+3]
                if val1 == val2:
                    self.data[position] = 1
                else:
                    self.data[position] = 0
                i += 4
            operation_params = self.parse_operation(self.data[i])

with open('input', 'r') as f:
    data = [int(s) for s in f.read().split(",")]

my_computer = Computer(data)

# part 1
maximum = 0
for combo in itertools.permutations([0,1,2,3,4], 5):
    second_input_value = 0
    for num in combo:
        second_input_value = my_computer.compute([num,second_input_value])
        my_computer.input_counter=0
    if second_input_value > maximum:
        maximum = second_input_value
print(maximum)