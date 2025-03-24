<?php

// php에서 배열을 순회하는 방법은 반복문의 종류와 배열의 
// 구조(인덱스 배열, 연관 배열)에 따라 달라진다.

// 학생 점수 배열 정의 (연관 배열)
$scores = [
  "철수" => 85,
  "영희" => 92,
  "민수" => 78,
];

// 1. for 반복문: 인덱스 배열 순회 
$fruits = ["사과", "바나나", "오렌지"];
echo "=== 과일 목록 (for) ===\n";
for ($i = 0; $i < count($fruits); $i++) {
  echo "과일 " . ($i + 1) . ": " . $fruits[$i] . "\n";
}

// 2. foreach 반복문: 연관 배열 순회 
echo "\n=== 학생 점수 (foreach) ===\n";
foreach ($scores as $name => $score) {
  echo "$name 학생의 점수는 {$score}점입니다.\n";
}

// 3. foreach로 배열 수정, 참조 사용
echo "\n=== 점수 5점 추가 ===\n";
foreach ($scores as $name => &$score) {
  // 참조(&)로 원본 배열 값 수정 
  $score += 5;
}

unset($score); // 참조 변수 해제 (이후 변수 재사용시 문제 방지)
foreach ($scores as $name => $score) {
  echo "$name 학생의 수정된 점수는 {$score}점입니다.\n";
}

// 4. 다차원 배열 순회
$students = [
  ["이름" => "지민", "나이" => 20, "전공" => "컴퓨터공학"],
  ["이름" => "수진", "나이" => 21, "전공" => "전자공학"],
];
echo "\n=== 학생 정보 ===\n";
foreach ($students as $student) {
  // 다차원 배열의 각 요소를 순회하며 정보 출력 
  echo "{$student['이름']} (나이: {$student['나이']}, 전공: {$student['전공']})\n";
}