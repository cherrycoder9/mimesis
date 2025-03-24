// std::find는 주어진 범위에서 특정 값을 가진 첫번째 요소를
// 선형 검색으로 찾아 해당 위치의 반복자를 반환함
// 값이 없으면 end 반복자를 반환함
// 시간 복잡도는 O(n)

#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

int main(int argc, char const *argv[]) {
  vector<int> numbers = {1, 2, 3, 4, 5};

  // 값 3을 검색
  auto it = find(numbers.begin(), numbers.end(), 3);

  // 검색 결과 확인 및 출력
  if (it != numbers.end()) {
    cout << "값 3을 찾았습니다. 위치: " << (it - numbers.begin()) << endl;
  } else {
    cout << "값 3을 찾지 못했습니다." << endl;
  }

  return EXIT_SUCCESS;
}
