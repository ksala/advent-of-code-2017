import math


def find_top_corner(steps):
    max_value = 1
    while steps:
        max_value = max_value + (8 * steps)
        steps = steps - 1
    return max_value


def find_sides(top_corner):
    side_lenght = int(math.sqrt(top_corner)) - 1
    return [
        (top_corner, top_corner - (side_lenght)),
        (top_corner - (side_lenght), top_corner - 2 * side_lenght),
        (top_corner - 2 * (side_lenght), top_corner - 3 * side_lenght),
        (top_corner - 3 * (side_lenght), top_corner - 4 * side_lenght + 1),
    ]


def find_side(value, sides):
    for side in sides:
        if side[0] > value > side[1]:
            return side


if __name__ == "__main__":
    value = 289326
    i = 0
    top_corner = 0
    while True:
        i = i + 1
        top_corner = find_top_corner(i)
        if top_corner > value:
            break

    sides = find_sides(top_corner)
    side = find_side(value, sides)
    middle = (side[0] + side[1]) // 2
    step_side = abs(middle - value)
    total_steps = step_side + i
    print(total_steps)
