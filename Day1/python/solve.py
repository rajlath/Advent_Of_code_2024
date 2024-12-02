lines = open("../input.txt").read().splitlines()
listA = []
listB = []
for line in lines:
	A, B = map(int, line.split())
	listA.append(A)
	listB.append(B)
listA.sort()
listB.sort()

def part_1():
	return  sum([abs(x - y) for (x,y) in zip(listA, listB)])
	# total_distance = 0
	# for x, y in zip(listA, listB):
	# 	total_distance += abs(x - y)
	# return total_distance

def part_2():
	return sum([x * (listB.count(x)) for x in listA])
	# simillarity_score = 0
	# for x in listA:
	# 	simillarity_score += x * (listB.count(x))
	# return simillarity_score

print(part_1())



