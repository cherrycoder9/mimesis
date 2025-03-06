<!-- 
  [실행시 출력됨]
  변수는 $ 기호로 시작
  초기화하지 않은 변수에 접근하면 기본적으로 null 반환 
  php에서 변수의 스코프는 크게 전역, 지역, 정적으로 나뉨
  # 전역 스코프 
  함수나 클래스 외부에서 선언된 변수, 파일 전체에서 접근 가능
  만약 함수 내부에서 전역 변수를 수정하려면 global 키워드나 $GLOBALS 배열을 통해야함
  # 지역 스코프
  함수 내부에서 선언된 변수, 함수가 종료되면 메모리에서 해제됨
  # 정적 스코프
  함수 내부에서 static 키워드로 선언된 변수
  함수가 종료된 후에도 값이 유지됨, 함수 호출간 데이터를 유지해야 할때 유용
-->

<?php
// 전역 변수 선언
$globalCounter = 10;

// 함수 정의
function incrementCounter()
{
  //지역 변수 선언
  $localCounter = 5;

  // 함수 안에서 전역 변수를 사용하려면 global 키워드로 명시적 선언 해야함
  global $globalCounter;
  $globalCounter++;

  // 지역 변수 출력, 함수 내부에서만 유효 
  echo "함수 내부에서의 지역 변수 값: $localCounter\n";
  echo "함수 내부에서의 전역 변수 값: $globalCounter\n";
}

// 정적 변수 사용
function staticCounter()
{
  // 정적 변수 선언
  static $staticCounter = 0;
  $staticCounter++;

  echo "정적 변수 값: $staticCounter\n";
}

// 함수 호출 
echo "함수 호출전 전역변수 값: $globalCounter\n";
incrementCounter();
echo "함수 호출후 전역변수 값: $globalCounter\n";

// 정적 변수 동작 확인
staticCounter();
staticCounter();
staticCounter();

// 초기화되지 않은 변수 접근
$undefinedVariable = null;
// null 병합연산자 사용 
echo "초기화되지 않은 변수값: " . ($undefinedVariable ?? "없음") . "\n";
