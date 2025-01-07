#!/usr/bin/env python3

import sys
import argparse


class MultiplierAction(argparse.Action):
	def __call__(self, parser, namespace, values, option_string=None):
		setattr(namespace, self.dest, [each * 10 for each in values])


def main():
	parser = argparse.ArgumentParser(description="taking argparse for a ride")
	parser.add_argument(
		"integers",
		default=[1, 2, 3],
		metavar="int",
		nargs="+",
		type=int,
		help="an integer",
		action=MultiplierAction,
	)
	parser.add_argument(
		"--log",
		default=sys.stdout,
		type=argparse.FileType("w"),
		help="a file to write stuff",
	)

	args = parser.parse_args()
	args.log.write("something somethings\n")
	if args.log != sys.stdout:
		args.log.close()
	print(args.integers)


if __name__ == "__main__":
	main()
