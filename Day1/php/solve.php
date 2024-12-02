<?php
$file  = fopen("../input.txt", "r");
$lines = [];
$listA = [];
$listB = [];
while (! feof($file)) {
    $lines[] = fgets($file);
}
fclose($file);

foreach ($lines as $line) {
    preg_match_all('/(\d+)/', $line, $matches);
    $listA[] = $matches[0][0];
    $listB[] = $matches[0][1];
}

function part_1($list1, $list2)
{
    sort($list1);
    sort($list2);
    $lens     = count($list1);
    $indx     = 0;
    $distance = 0;
    while ($indx < $lens) {
        $distance += abs($list1[$indx] - $list2[$indx]);

        $indx++;
    }
    return $distance;
}

function part_2($list1, $list2)
{
    $lens              = count($list1);
    $indx              = 0;
    $simillarity_score = 0;

    while ($indx < $lens) {
        $curr = $list1[$indx];
        $simillarity_score += array_key_exists($curr, $list2) ? $curr * $list2[$curr] : 0;
        $indx++;
    }
    return $simillarity_score;

}

print(part_1($listA, $listB)); //1590491
$countValues = array_count_values($listB);
print("\n");
print(part_2($listA, $countValues)); //22588371
