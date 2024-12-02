from collections import Counter


def main() -> None:
    file_name = "./day-1/input.txt"
    left = []
    right = []
    with open(file_name) as file:
        for line in file.readlines():
            left_value, right_value = map(int, line.split("   "))
            left.append(left_value)
            right.append(right_value)
    right_counts = Counter(right)

    similarity_score = 0
    for number in left:
        similarity_score += number * right_counts.get(number, 0)
    left.sort()
    right.sort()
    result = 0
    for index, left_value in enumerate(left):
        result += abs(left_value - right[index])
    print(f"{result=}")
    print(f"{similarity_score=}")


if __name__ == "__main__":
    main()
