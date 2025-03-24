// weak_ptr
// weak_ptr은 shared_ptr과 함께 사용됨
// 메모리에 대한 약한 참조를 제공
// 참조 카운트를 증가시키지 않으므로 메모리 해제에 직접적 영향을 주지 않음
// 주로 shared_ptr 간의 순환 참조 문제를 해결하는데 사용됨
// 소유권을 가지지 않음
// lock() 메서드: shared_ptr로 변환해 안전하게 메모리 접근 가능
// expired() 메서드: 참조하던 메모리가 해제되었는지 확인 가능
#include <cstdlib>
#include <iostream>
#include <memory>

using namespace std;

int main(int argc, char const *argv[]) {
  // shared_ptr 생성
  auto sp = make_shared<int>(30);
  weak_ptr<int> wp = sp;  // weak_ptr로 약한 참조 설정
  cout << "참조카운트: " << sp.use_count() << endl;
  cout << "참조카운트: " << wp.use_count() << endl;

  // weak_ptr를 shared_ptr로 변환해 사용
  if (auto locked = wp.lock()) {
    cout << "locked가 가리키는 값: " << *locked << endl;  // 30
  } else {
    cout << "메모리가 이미 해제되었습니다." << endl;
  }

  sp.reset();  // shared_ptr 해제, 참조 카운트가 0이 됨

  // 메모리 해제 여부 확인
  if (wp.expired()) {
    cout << "메모리가 해제되었습니다." << endl;
    cout << "참조카운트: " << sp.use_count() << endl;  // 0
    cout << "참조카운트: " << wp.use_count() << endl;  // 0
  }

  return EXIT_SUCCESS;
}
