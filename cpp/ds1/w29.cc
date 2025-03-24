#include <cstdlib>
#include <iostream>

using namespace std;

int main(int argc, char const* argv[]) {
  int* ptr = new int;  // int 크기의 메모리 할당
  *ptr = 42;
  cout << "동적 할당된 정수 값: " << *ptr << endl;
  delete ptr;
  ptr = nullptr;  // 안전을 위한 조치

  // 정수형 배열 동적 할당
  int* arr = new int[5];
  for (size_t i = 0; i < 5; ++i) {
    arr[i] = i * 10;
  }
  cout << "동적 할당된 배열 요소: ";
  for (size_t i = 0; i < 5; ++i) {
    cout << arr[i] << " ";
  }
  cout << endl;
  delete[] arr;  // 배열 메모리 해제
  arr = nullptr;

  // 메모리 할당 실패 테스트
  int* bigArr = new (nothrow) int[1000000000000000000];
  if (bigArr == nullptr) {
    cout << "메모리 할당에 실패했습니다." << endl;
    exit(EXIT_FAILURE);  // 실패시 프로그램 종료
  } else {
    cout << "메모리 할당에 성공했습니다." << endl;
    delete[] bigArr;
    bigArr = nullptr;
  }

  // C++11 이후로는 new와 delete를 직접 사용하는 대신 스마트포인터 권장
  // 메모리 누수와 이중 해제를 방지하며 RAII 원칙을 따름

  exit(EXIT_SUCCESS);
}
