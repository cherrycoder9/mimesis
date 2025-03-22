// std::map
// 키-값 쌍을 저장하는 연관 컨테이너
// 내부적으로 레드-블랙 트리를 사용해 키를 정렬함
// 키를 기준으로 오름차순 정렬됨
// O(log n) 시간 복잡도로 키 검색, 삽입, 삭제 가능
// 중복된 키는 허용되지 않음

// insert(pair<K, V>): 키-값 쌍 삽입
// find(K key): 특정 키 검색
// erase(K key): 특정 키 삭제
// operator[]: 키로 값 접근(없으면 새 요소 삽입)

#include <cstdlib>
#include <iostream>
#include <map>
#include <string>

using namespace std;

int main(int argc, char const* argv[]) {
  map<string, int> ageMap;

  // 요소 삽입
  ageMap["홍길동"] = 30;
  ageMap["김철수"] = 25;
  ageMap.insert({"이영희", 28});

  // 요소 접근
  cout << "홍길동 나이: " << ageMap["홍길동"] << endl;

  // 모든 요소 출력
  cout << "맵의 모든 요소:" << endl;
  for (const auto& pair : ageMap) {
    cout << pair.first << ": " << pair.second << endl;
  }

  // 요소 삭제
  ageMap.erase("김철수");
  cout << "삭제 후 맵:" << endl;
  for (const auto& pair : ageMap) {
    cout << pair.first << ": " << pair.second << endl;
  }

  // 키 존재 여부 확인
  if (ageMap.find("이영희") != ageMap.end()) {
    cout << "이영희가 맵에 존재합니다." << endl;
  }

  return EXIT_SUCCESS;
}
