<?php
// php에서 숫자형은 크게 정수형(int)과 부동소수점형(float)으로 나뉨 
// 타입 확인 함수: is_int(), is_float()
// 캐스팅 연산자: (int), (float)
// 문자열은 string 타입으로 표현됨, 작은따옴표나 큰따옴표로 감싸서 선언
// 큰따옴표 문자열은 변수와 일부 특수문자(\n, \t 등)를 해석해 동적으로 처리함 
// 히어독(heredoc) 및 나우독(nowdoc): 복잡한 문자열이나 HTML 코드를 포함할때
// 히어독은 변수 해석 지원, 나우독은 지원 안함 
// 히어독은 <<<식별자 형태로 작성 <<<BYE
// 나우독은 <<<'식별자' 형태로 작성 <<<'HELLO'
// 문자열을 다루기 위한 대표적 내장 함수들
// strlen(), str_replace(), substr(), trim() 등
// 정규 표현식 (preg_* 함수) 조작 가능
// 문자열과 숫자형 간의 변환: (string), strval(), (int), intval()

$integerNumber = 42;
$floatNumber = 3.14;

$sum = $integerNumber + $floatNumber;
$product = $integerNumber * $floatNumber;

echo "덧셈결과: $sum\n";
echo "곱셈결과: $product\n";

// .연산자는 문자열 연결을 위한 연산자임, +연산자는 숫자형 데이터의 덧셈에만 사용 
echo "정수형 여부: " . (is_int($integerNumber) ? "참" : "거짓") . "\n";
echo "부동소수점 여부: " . (is_float($floatNumber) ? "참" : "거짓") . "\n";

// 타입 변환
$floatToInt = (int) $floatNumber; // 소수점 이하 버림
echo "부동소수점을 정수로 변환: $floatToInt\n";
$stringNumber = "123.45";
$floatFromString = (float) $stringNumber;
echo "문자열을 부동소수점으로 변환: $floatFromString\n";


$singleQuoteString = '안녕, 정민수';
$name = "정민수";
$doubleQuoteString = "안녕, $name";

// 문자열 길이 확인
$length = strlen($doubleQuoteString);
echo "문자열 길이: $length\n"; // 17, 한글은 3바이트 

// 문자열 치환 
$replacedString = str_replace("안녕", "반갑습니다", $doubleQuoteString);
echo "문자열 치환 결과: $replacedString\n";

// 문자열 자르기
$substring = substr($doubleQuoteString, 0, 6);
echo "문자열 자르기 결과: $substring\n"; // 안녕

// 히어독 문자열 선언
$heredocString = <<<EOD
여러 줄로 구성된 문자열
이름: $name
줄바꿈 가능
EOD;
echo "히어독 문자열:\n$heredocString\n";