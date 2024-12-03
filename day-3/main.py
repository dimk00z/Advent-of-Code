import re


def sum_valid_mul_instructions(memory):
    pattern = r"mul\((\d{1,3}),(\d{1,3})\)"
    matches = re.findall(pattern, memory)
    total = 0
    for x, y in matches:
        total += int(x) * int(y)

    return total


with open("input.txt") as file:
    result = sum_valid_mul_instructions(file.read())
print(result)
