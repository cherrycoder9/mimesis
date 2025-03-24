// std:sort는 STL에서 가장 널리 사용되는 정렬 알고리즘
// 주어진 범위의 요소를 기본적으로 오름차순으로 정렬함
// 내부적으로는 인트로소트(Introsort)라는 하이브리드 정렬 알고리즘을
// 사용해 평균 O(n log n) 시간 복잡도를 보장함
// std::sort는 두 개의 반복자를 인자로 받아 범위를 지정함

#include <algorithm>
#include <cstdlib>
#include <iostream>
#include <vector>

using namespace std;

int main(int argc, char const *argv[]) {
  vector<int> numbers = {5, 3, 1, 4, 2};  // 초기화된 벡터

  // 벡터를 오름차순으로 정렬
  sort(numbers.begin(), numbers.end());

  // 정렬된 결과 출력
  cout << "정렬된 숫자: ";
  for (int num : numbers) {
    cout << num << " ";
  }
  cout << endl;

  return EXIT_SUCCESS;
}
