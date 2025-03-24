// 조건 기반 검색 std::find_if
// 단항 함수를 인자로 받아 조건을 정의
// 첫번째 짝수를 찾는 코드
#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

int main(int argc, char const *argv[]) {
  vector<int> numbers = {1, 3, 5, 2, 4};

  // 첫 번째 짝수를 찾기 위한 조건 정의
  auto it =
      find_if(numbers.begin(), numbers.end(), [](int n) { return n % 2 == 0; });

  // 검색 결과 확인 및 출력
  if (it != numbers.end()) {
    cout << "첫번째 짝수: " << *it << ", 위치: " << (it - numbers.begin())
         << endl;
  } else {
    cout << "짝수를 찾지 못했습니다." << endl;
  }

  return EXIT_SUCCESS;
}