import java.io.BufferedReader
import java.io.InputStreamReader
import java.io.File
import kotlin.math.abs

fun part_1(arr1:MutableList<Int>, arr2:MutableList<Int>):Int{
	var arr11 = arr1.sorted()
	var arr21 = arr2.sorted()
	var i = 0
	var distant = 0
	while (i < arr1.size){
		distant += abs(arr11[i] - arr21[i])
		i += 1
	}
	return distant
}

fun getCount( arr2:MutableList<Int>):MutableMap<Int, Int>{
	val frequencyMap: MutableMap<Int, Int> = HashMap() 
    for (s in arr2)
    {
        var count = frequencyMap[s]
        if (count == null) count = 0
        frequencyMap[s] = count + 1
    }
 	return frequencyMap

}

fun part_2(arr1:MutableList<Int>, arr2:MutableList<Int>):Int{	
	var frequencyMap: MutableMap<Int, Int> = getCount(arr2)
	var simillarity_score = 0
	for (arr in arr1){
		var vals = frequencyMap.get(arr)
		var toUse= if (vals == null) 0 else vals
		simillarity_score += (arr * toUse)

	}
	


	return simillarity_score

}



fun main(){
	val lines = File("../input.txt").readLines()
	var listA: MutableList<Int> = mutableListOf()
	var listB: MutableList<Int> = mutableListOf()
	for (line in lines){
		val regex = Regex("(\\d+)")
    	val resultList = regex.findAll(line).map { it.value }.toList().map { it.toInt() }
    	listA.add(resultList[0])
    	listB.add(resultList[1])
   	}
   	println(part_1(listA, listB))    //1590491
   	println(part_2(listA, listB))	//22588371
}


	

