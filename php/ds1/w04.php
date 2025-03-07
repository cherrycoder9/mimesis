<?php
// # 배열 선언하는 방법
// 1. array() 함수를 활용한 선언
// 2. 짧은 배열 문법([])을 사용한 선언, php5.4 이후 도입 

// 인덱스 배열 선언
$fruits = ['사과', '바나나', '오렌지'];
// 배열 끝에 요소 추가
$fruits[] = '포도';
// 특정 인덱스 값 수정
$fruits[1] = '망고';

// 배열 출력 (디버깅용)
echo "과일 목록:\n";
foreach ($fruits as $index => $fruit) {
  echo "{$index}번 과일: $fruit\n";
}

// 연관 배열 선언
$student = [
  'name' => '김정민',
  'age' => 20,
  'major' => '컴퓨터공학'
];

// 연관 배열에 새 키-값 쌍 추가
$student['grade'] = 'A';

// 연관 배열 출력 
echo "\n학생 정보:\n";
foreach ($student as $key => $value) {
  // . 연결 연산자 대신 중괄호를 사용한 보간 방식 적용 
  echo "{$key}: {$value}\n";
}

unset($student['age']);

// 삭제 후 결과 확인
echo "\n나이 삭제후 학생정보:\n";
foreach ($student as $key => $value) {
  // printf 사용 방식 예시 
  printf("%s: %s\n", $key, $value);
}