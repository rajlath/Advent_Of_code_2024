const fs = require("fs")

let data = fs.readFileSync("../input.txt").toString().split(/\r?\n/);
let listA = []
let listB = []

for (line of data){
	let [A, B] = line.split("   ").map(v=>parseInt(v))
	listA.push(A)
	listB.push(B)	
}

listA.sort((a, b) => a - b);
listB.sort((a, b) => a - b);

const countOccurrences = (arr, val) => arr.reduce((a, v) => (v === val ? a + 1 : a), 0);

function part1(){
	let totalDistance = 0
	for (indx in listA){
		totalDistance += Math.abs(listA[indx] - listB[indx])
	}
	return totalDistance
}

function part2(){
	let simillarityScore = 0
	for (valA of listA){
		simillarityScore += valA * countOccurrences(listB, valA)
	}
	return simillarityScore
}
console.log(part1())  //1590491 	
console.log(part2())  //22588371