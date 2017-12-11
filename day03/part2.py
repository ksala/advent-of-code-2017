#!/usr/bin/env python3

grid = {}


def sum_neighboors(x, y):
    sum = 0
    for i in range(-1, 2):
        for k in range(-1, 2):
            sum = sum + grid.get((x + i, y + k), 0)
    return sum


def print_grid(grid):
    print("---")
    for i in range(-10, 10):
        for k in range(-10, 10):
            print(grid.get((i, k), 0), end=" ")
        print()
    print("---")


def move_right(x, y):
    return x, y + 1


def move_down(x, y):
    return x - 1, y


def move_up(x, y):
    return x + 1, y


def move_left(x, y):
    return x, y - 1


def change_move(move):
    if move == move_right:
        return move_down
    if move == move_down:
        return move_left
    if move == move_left:
        return move_up
    if move == move_up:
        return move_right


value = 289326

# Starting point
x = 0
y = 0

grid[(x, y)] = 1

steps = 1
next_steps = 1
move = move_right
index_loop = 0
while grid[(x, y)] < value:
    print("Looped: %s times" % index_loop)

    while steps > 0:
        print("Moving: %s for %s steps" % (move, steps))
        x, y = move(x, y)
        grid[(x, y)] = sum_neighboors(x, y)
        if grid[(x, y)] > value:
            break # Value found, exit early
        steps = steps - 1
    
    steps = next_steps
    index_loop = index_loop + 1
    next_steps = next_steps + (index_loop % 2)
    move = change_move(move)
    print_grid(grid)

print(grid[(x, y)])
