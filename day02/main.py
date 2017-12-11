#!/usr/bin/env python3

if __name__ == "__main__":
    # Oneliner! First phase
    print(sum(map(lambda x: max(x) - min(x), map(lambda x: map(int, x.split("\t")), open("input.txt").read()[:-1].split("\n")))))


    # Second phase
    puzzle_input = open("input.txt").read()[:-1]

    checksum = 0
    speadsheet = map(lambda x: map(int, x.split("\t")), puzzle_input.split("\n"))
    for row in speadsheet:
        for num1 in row:
            for num2 in row:
                if num1 != num2 and num1 % num2 == 0:
                    checksum = checksum + (num1 // num2)
    print(checksum)