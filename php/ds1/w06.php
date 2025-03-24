<?php
$num1 = 10;
$num2 = 3;

$sum = $num1 + $num2;
$difference = $num1 - $num2;
$product = $num1 * $num2;
$quotient = $num1 / $num2; // 부동소수점 결과 
$remainder = $num1 % $num2;
$power = $num1 ** $num2;

echo "산술 연산 결과: \n";
echo "덧셈: $sum\n";
echo "뺄셈: $difference\n";
echo "곱셈: $product\n";
echo "나눗셈: $quotient\n";
echo "나머지: $remainder\n";
echo "거듭제곱: $power\n";

$intValue = 5;
$stringValue = "5";

$looseEqual = $intValue == $stringValue; // 느슨한 동등 비교 
$strictEqual = $intValue === $stringValue; // 엄격한 일치 비교
$notEqual = $intValue != $stringValue; // 부등 비교
$greaterThan = $intValue > 3; // 크기 비교
$spaceship = $intValue <=> 7; // 우주선 연산자 
// -1: 왼쪽 값이 오른쪽 값보다 작을때
// 0: 두 값이 같을 때
// 1: 왼쪽 값이 오른쪽 값보다 클때 

echo "\n비교 연산 결과:\n";
echo ($looseEqual ? "참" : "거짓") . "\n"; // 참 
echo ($strictEqual ? "참" : "거짓") . "\n"; // 거짓
echo ($notEqual ? "참" : "거짓") . "\n"; // 거짓
echo ($greaterThan ? "참" : "거짓") . "\n"; // 참 
echo "$spaceship\n"; // -1