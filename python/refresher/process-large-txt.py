#!/usr/bin/env python3


import sys

from functools import partial
from itertools import chain
from collections import Counter


def main():
    filepath = sys.argv[1]
    with open(filepath) as file:
        data = file.read()

    print(
        Counter(chain(*map(partial(str.split, sep=" "), data.split("\n")))).most_common(
            10
        )
    )


if __name__ == "__main__":
    main()
