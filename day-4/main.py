def count_xmas_occurrences(grid, word="XMAS"):
    rows, cols = len(grid), len(grid[0])
    word_length = len(word)
    directions = [
        (0, 1),  # Horizontal right
        (0, -1),  # Horizontal left
        (1, 0),  # Vertical down
        (-1, 0),  # Vertical up
        (1, 1),  # Diagonal down-right
        (-1, -1),  # Diagonal up-left
        (1, -1),  # Diagonal down-left
        (-1, 1),  # Diagonal up-right
    ]

    def is_valid(x, y):
        """Check if a position is within bounds."""
        return 0 <= x < rows and 0 <= y < cols

    def check_direction(x, y, dx, dy):
        """Check if 'word' exists starting from (x, y) in direction (dx, dy)."""
        for i in range(word_length):
            nx, ny = x + i * dx, y + i * dy
            if not is_valid(nx, ny) or grid[nx][ny] != word[i]:
                return False
        return True

    count = 0
    for x in range(rows):
        for y in range(cols):
            for dx, dy in directions:
                if check_direction(x, y, dx, dy):
                    count += 1
    return count


with open("./day-4/input.txt") as file:
    grid = [line.strip() for line in file.readlines()]


# Count the occurrences of "XMAS"
print(
    "Total occurrences of XMAS:",
    count_xmas_occurrences(grid),
)
