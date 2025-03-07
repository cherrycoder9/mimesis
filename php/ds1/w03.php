<?php
// php7.0 이후부터 스칼라 타입 선언과 반환 타입 선언이 도입되었음
// 엄격한 타입 선언을 활용하면 타입 관련 오류를 줄일수 있음
declare(strict_types=1);

function addNumbers(float $a, float $b): float
{
  return $a + $b;
}

// $result = addNumbers(3.14, "1");
$result = addNumbers(3.14, 1);
echo "덧셈결과: $result\n";


// 높은 정밀도가 요구되는 경우 bcmath 확장 모듈 사용권장
// php.ini에서 extension=bcmath 설정
$num1 = "1234.657373";
$num2 = "241.546919";
$bcSum = bcadd($num1, $num2, 10); // 세번째 인자는 소수점 아래 자릿수
echo "bcmath로 덧셈: $bcSum\n";
$bcProduct = bcmul($num1, $num2, 10);
echo "곱셈: $bcProduct\n";
$bcDiv = bcdiv($num1, $num2, 10);
echo "나눗셈: $bcDiv\n";
// 제곱근 계산
$bcSqrt = bcsqrt("2", 10); // 2의 제곱근
echo "제곱근: $bcSqrt\n";


// bcstring은 멀티바이트 문자열을 처리하기 위한 PHP 확장 모듈 
// 기본 문자열 함수 사용 
$koreanString = "안녕hi하세요, PHP";
$byteLength = strlen($koreanString);
$byteSubstring = substr($koreanString, 0, 9); // 9바이트만 자름
echo "기본 문자열 길이: $byteLength\n"; // 22
echo "기본 문자열 자르기: $byteSubstring\n"; // 안녕hi�

// mbstring을 사용
$mbLength = mb_strlen($koreanString, "UTF-8");
$mbSubstring = mb_substr($koreanString, 0, 4, "UTF-8"); // 처음 4글자 자르기 
echo "멀티바이트 문자열 길이: $mbLength\n"; // 12
echo "멀티바이트 문자열 자르기: $mbSubstring\n";

// 대소문자 변환
$mbUpper = mb_strtoupper("php 이즈 굿", "UTF-8");
echo "대문자 변환: $mbUpper\n";

// 문자열 검색
$position = mb_strpos($koreanString, "hi", 0, "UTF-8");
echo "문자열 위치: $position\n";

// 문자열 치환
mb_regex_encoding("UTF-8"); // 멀티바이트 정규 표현식 인코딩 설정
$mbReplace = mb_ereg_replace("안녕", "hello ", $koreanString);
echo "문자열 치환: $mbReplace\n";