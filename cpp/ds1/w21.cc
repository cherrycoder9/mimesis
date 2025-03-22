// std::list
// 이중 연결 리스트를 구현한 시퀀스 컨테이너
// 각 요소가 다음 및 이전 요소와 연결되어 있음
// 요소가 메모리 상에서 흩어져 있어 캐시 효율성이 낮음
// 인덱스 접근 대신 반복자를 통해 순회해야 함
// 중간 요소 삽입 및 삭제가 O(1) 시간 복잡도로 수행됨

// push_back(), push_front(): 끝 또는 시작에 요소 추가
// pop_back(), pop_front(): 끝 또는 시작 요소 제거
// insert(iterator pos, T value): 특정 위치에 요소 삽입
// erase(iterator pos): 특정 위치의 요소 삭제

#include <cstdlib>
#include <iostream>
#include <list>
#include <string>

using namespace std;
int main(int argc, char const* argv[]) {
  list<string> lst;  // string 타입 요소를 저장하는 리스트 생성

  // 요소 추가
  lst.push_back("사과");
  lst.push_front("바나나");
  lst.push_back("체리");

  // 모든 요소 출력
  cout << "리스트의 모든 요소: ";
  for (const auto& elem : lst) {
    cout << elem << " ";
  }
  cout << endl;

  // 두 번째 위치에 요소 삽입
  auto it = lst.begin();
  ++it;                         // 두번째 위치로 이동
  it = lst.insert(it, "딸기");  // 딸기 삽입후 반환된 반복자를 it에 저장

  // 삽입 후 출력
  cout << "삽입 후 리스트: ";
  for (const auto& elem : lst) {
    cout << elem << " ";
  }
  cout << endl;

  // 요소 삭제
  lst.erase(it);  // 딸기 삭제
  cout << "삭제 후 리스트: ";
  for (const auto& elem : lst) {
    cout << elem << " ";
  }
  cout << endl;

  return EXIT_SUCCESS;
}
