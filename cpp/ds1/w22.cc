// std::deque
// 양쪽 끝에서 요소를 추가하거나 제거할 수 있는 시퀀스 컨테이너
// 양쪽 끝에서 요소를 추가하거나 제거할때 빠른 연산 가능함 O(1)
// 인덱스를 사용해 특정 위치의 요소에 빠르게 접근할 수 있음
// 중간에 요소를 삽입하거나 삭제하면 O(n) 시간
// 연속된 메모리 블록이 아닌, 필요에 따라 블록 단위로 메모리를 할당
// 큐나 스택을 구현할때 유용함

#include <deque>
#include <iostream>

using namespace std;

int main(int argc, char const* argv[]) {
  deque<int> deq;  // 정수형 deque 생성

  // 요소 추가
  deq.push_back(10);
  deq.push_front(20);
  deq.push_back(30);

  // 요소 출력
  cout << "deque 요소: ";
  for (const auto& elem : deq) {
    cout << elem << " ";
  }
  cout << endl;

  // 앞과 뒤 요소 제거
  deq.pop_front();  // 앞 요소 제거
  deq.pop_back();   // 뒤 요소 제거

  // 제거 후 출력
  cout << "제거 후 deque: ";
  for (const auto& elem : deq) {
    cout << elem << " ";
  }
  cout << endl;

  return 0;
}
