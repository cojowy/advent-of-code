with open('input', 'r') as f:
    data = f.read()

def number_frequency_in_list_of_strings(strings, number):
    c=0
    for s in strings:
        c+=s.count(str(number))
    return c

def winning_pixel(pixels):
    for px in pixels:
        if px == '0' or px == '1':
            return px
    return 2

pixels = (25,6)

i=0
lines = []
while i < len(data):
    lines.append(data[i:i+pixels[0]])
    i += pixels[0]

i=0
layers = []
while i < len(lines):
    layers.append(lines[i:i+pixels[1]])
    i += pixels[1]

min_zero_count, min_zero_layer = number_frequency_in_list_of_strings(layers[0], 0), layers[0]
for l in layers:
    if number_frequency_in_list_of_strings(l, 0) < min_zero_count:
        min_zero_count = number_frequency_in_list_of_strings(l, 0)
        min_zero_layer = l 

result_pixels=[]
for i in range(pixels[1]):
    result_pixels.append([])
    for j in range(pixels[0]):
        result_pixels[-1].append(winning_pixel([l[i][j] for l in layers]))

for l in result_pixels:
    print(",".join(l))
