// shared_ptr
// 다중 소유권을 지원하는 스마트 포인터
// 여러 shared_ptr 객체가 동일한 메모리 블록을 공유할 수 있으며
// 참조 카운팅 메커니즘을 통해 메모리 해제를 관리함
// 참조 카운트가 0이 되는 순간, 즉 모든 shared_ptr이 소멸될 때 메모리가 해제됨
// 복사 생성자와 대입 연산자를 통해 동일한 메모리를 여러 포인터가 공유가능
// use_count() 메서드로 현재 참조 수를 확인할 수 있음
// 마지막 shared_ptr 소멸될때 메모리가 해제됨
#include <cstdlib>
#include <iostream>
#include <memory>

using namespace std;

int main(int argc, char const *argv[]) {
  // shared_ptr 사용해 동적 메모리 할당
  auto ptr1 = make_shared<int>(20);
  cout << "ptr1이 가리키는 값: " << *ptr1 << endl;

  // 복사를 통한 소유권 공유
  auto ptr2 = ptr1;  // ptr2가 ptr1과 동일한 메모리 공유
  cout << "ptr1 가리키는 값: " << *ptr1 << endl;
  cout << "ptr2 가리키는 값: " << *ptr2 << endl;
  cout << "참조 카운트: " << ptr1.use_count() << endl;

  return EXIT_SUCCESS;
}
