#!/usr/bin/python3

def solve_captcha(puzzle_input, step):
    sum = 0
    for index, value in enumerate(puzzle_input):
        next_index = (index + step) % len(puzzle_input)
        if value == puzzle_input[next_index]:
            sum += value
    return sum

if __name__ == "__main__":
    puzzle_input = open("input.txt").read()
    puzzle_input = list(map(int, puzzle_input))

    # Phase 1
    print(solve_captcha(puzzle_input, 1))

    # Phase 2
    print(solve_captcha(puzzle_input, len(puzzle_input) // 2))