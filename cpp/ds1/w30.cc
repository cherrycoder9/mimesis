// unique_ptr
// 단일 소유권을 보장하는 스마트 포인터
// 특정 메모리 블록을 단 하나의 unique_ptr만 소유할 수 있음
// 소유권은 복사가 아닌 이동을 통해 전달됨 (move semantics)
// unique_ptr 객체가 소멸되면 자신이 관리하던 메모리 자동 해제
// 복사 생성자나 복사 대입 연산자가 삭제되어 있어 복사가 불가능하며
// std::move를 사용해 소유권을 이동해야 함

#include <cstdlib>
#include <iostream>
#include <memory>

using namespace std;

int main(int argc, char const *argv[]) {
  // C++11 이상, unique_ptr 사용한 동적 메모리 할당
  auto ptr = make_unique<int>(10);  // 10으로 초기화된 int형 메모리 생성
  cout << "ptr이 가리키는 값: " << *ptr << endl;

  // 소유권 이동
  auto ptr2 = move(ptr);  // ptr의 소유권을 ptr2로 이동, ptr은 nullptr이 됨
  if (ptr == nullptr) {
    cout << "ptr은 nullptr입니다." << endl;
  }
  cout << "ptr2가 가리키는 값: " << *ptr2 << endl;

  return EXIT_SUCCESS;
}
