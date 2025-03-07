<?php
// null은 값이 없거나 정의되지 않은 상태를 나타내는 특별 타입
// 변수가 초기화되지 않았거나 명시적으로 null로 설정한 경우 
// 데이터베이스에서 필드가 비어있는 경우나 함수가 반환값을 가지지 않을때
// 느슨한 타입 평가를 방지하고 싶다면 엄격한 비교 연산자 사용 (===)
// unset()을 호출하면 변수 자체가 제거되지만, null로 설정하면 값만 없는 상태
// php7 이후 도입된 ??연산자는 null값 간편하게 처리 가능

// PHP › Format: Code Style K&R
function isAdult(int $age): bool {
  return $age >= 19; // 19세 이상이면 true, 아님 false 반환
}

function findUser(string $name): ?string {
  $users = ['김철수' => '철수', '이영희' => '영희'];
  return $users[$name] ?? NULL; // 이름이 없으면 null 반환 
}

$age = 20;
$isAdult = isAdult($age);
echo "나이 {$age}세는 성인인가요? ";
echo $isAdult ? "네, 성인입니다.\n" : "아니오, 미성년자입니다\n";

$user = findUser('김철수');
echo "찾은 사용자: " . ($user ?? "사용자가 없습니다.") . "\n";

$zero = 0;
$emptyString = "";
echo "0은 false인가요? " . ($zero == false ? "네\n" : "아니요\n");
echo "0은 false인가요? " . ($zero === false ? "네\n" : "아니요\n");

$nothing = null;
if (is_null($nothing)) {
  echo "변수는 null 상태입니다.\n";
}