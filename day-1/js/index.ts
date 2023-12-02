import { readFileSync } from "fs";

type Key = string;
type PatternType = {
	[Key]: number;
};
const patterns: PatternType[] = [
	{ "1": 1 },
	{ "2": 2 },
	{ "3": 3 },
	{ "4": 4 },
	{ "5": 5 },
	{ "6": 6 },
	{ "7": 7 },
	{ "8": 8 },
	{ "9": 9 },
];
function getFirstDigit(input: string) {
	for (let i = 0; i <= input.length; i++) {
		for (const pattern in patterns) {
			if ([...input][i] === pattern) {
				return Object.values(pattern);
			}
		}
	}
}
function main() {
	const file = readFileSync("../input.txt", "utf8");

	const lines = file.split("\n");

	console.log(lines.length);
}

main();
