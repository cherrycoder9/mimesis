// 사용자 정의 비교 기준
// std::sort는 세번째 인자로 비교 함수를 받아 정렬 기준을 커스터마이징
#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

int main(int argc, char const *argv[]) {
  vector<int> numbers = {5, 3, 1, 4, 2};

  // 내림차순 정렬을 위한 람다 함수 사용
  sort(numbers.begin(), numbers.end(), [](int a, int b) { return a > b; });

  // 정렬된 결과 출력
  cout << "내림차순 정렬된 숫자: ";
  for (int num : numbers) {
    cout << num << " ";
  }
  cout << endl;

  return EXIT_SUCCESS;
}
