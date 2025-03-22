// std::set
// 유일한 요소를 정렬된 순서로 저장하는 연관 컨테이너
// 내부적으로 레드-블랙 트리 사용
// 요소가 삽입되면 자동으로 오름차순으로 정렬됨
// 중복된 요소는 삽입되지 않음
// 삽입, 삭제, 검색이 모두 O(log n) 시간에 이루어짐
// 반복자를 통해 요소를 정렬된 순서로 접근할 수 있음

#include <iostream>
#include <set>

using namespace std;

int main(int argc, char const* argv[]) {
  set<int> s;  // 정수형 set 생성

  // 요소 삽입
  s.insert(10);
  s.insert(20);
  s.insert(10);

  // 요소 출력
  cout << "set 요소:";
  for (const auto& elem : s) {
    cout << elem << " ";
  }
  cout << endl;

  // 요소 검색
  if (s.find(20) != s.end()) {
    cout << "20이 set에 있습니다." << endl;
  }

  // 요소 삭제
  s.erase(10);
  cout << "10 삭제후 set: ";
  for (const auto& elem : s) {
    cout << elem << " ";
  }
  cout << endl;

  return 0;
}
