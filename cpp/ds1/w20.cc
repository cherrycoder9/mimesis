// 컨테이너는 데이터를 저장하고 관리하기 위한 템플릿 클래스
// 사용자가 원하는 데이터 타입을 저장할 수 있도록 설계됨
// 크게 시퀀스 컨테이너와 연관 컨테이너로 나뉨
// 시퀀스 컨테이너: 데이터를 순차적으로 저장하며 요소의 위치에 따라 접근함
// std::vector, std::list, std::deque
// 연관 컨테이너: 키-값 쌍 또는 정렬된 순서로 데이터를 저장하며 키를 통해
// 효율적으로 접근 std::map, std::set

// std::vector
// 동적 배열을 구현한 시퀀스 컨테이너
// 고정 크기 배열과 달리 크기가 가변적이며 메모리 관리를 자동으로 처리함
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

// push_back(T value): 벡터 끝에 요소를 추가
// pop_back(): 벡터 끝의 요소를 제거
// size(): 현재 요소 개수를 반환
// operator[]: 인덱스로 요소에 접근
// begin(), end(): 반복자를 반환해 순회 가능

int main(int argc, char const* argv[]) {
  vector<int> vec;  // int 타입 요소를 저장하는 벡터 생성

  // 요소 추가
  vec.push_back(10);
  vec.push_back(20);
  vec.push_back(30);

  // 첫 번째 요소와 크기 출력
  cout << "벡터의 첫번째 요소: " << vec[0] << endl;
  cout << "벡터의 크기: " << vec.size() << endl;

  // 모든 요소 순회 및 출력 (범위 기반 for문 사용)
  cout << "벡터의 모든 요소: ";
  for (const auto& elem : vec) {
    cout << elem << " ";
  }
  cout << endl;

  // 마지막 요소 제거
  vec.pop_back();
  cout << "pop_back 후 벡터 크기: " << vec.size() << endl;

  return EXIT_SUCCESS;
}
