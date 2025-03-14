// # 클래스 템플릿
// 함수 템플릿이 단일 함수의 동작을 일반화 한다면
// 클래스 템플릿은 객체의 구조와 동작을 데이터 타입에 구애받지 않도록 함
// std::vector, std::map 같은 컨테이너들이 클래스 템플릿의 사례

// 컴파일러는 클래스 템플릿을 기반으로 실제 사용되는 타입에
// 맞춰 구체적인 클래스를 인스턴스화
// 클래스 템플릿은 보통 헤더 파일에 선언과 정의를 함께 작성함
// 컴파일 타임에 인스턴스화되기 때문에, 링커가 아닌 컴파일러가 모든 코드를
// 볼 수 있어야 하기 때문

// 스택을 구현한 클래스 템플릿
#include <iostream>
#include <stdexcept>
#include <vector>

using namespace std;

// 템플릿 클래스 Stack, 제네릭 타입 T를 사용해 스택 구현
template <typename T>
class Stack {
 private:
  vector<T> elements;  // 데이터를 저장하는 벡터 컨테이너

 public:
  // 스택에 데이터 추가
  void push(const T& item) {
    // push_back: 벡터의 맨 끝에 새로운 요소를 추가함
    elements.push_back(item);
  }

  // 스택에서 데이터 제거 및 반환
  T pop() {
    if (isEmpty()) {
      throw out_of_range("스택이 비어 있습니다.");
    }
    // back: 벡터의 맨 마지막 요소를 참조함(제거하지 않고 확인만)
    T item = elements.back();
    // pop_back: 벡터의 맨 마지막 요소를 제거함
    elements.pop_back();
    return item;
  }

  // 스택의 최상단 데이터 확인
  T top() const {
    if (isEmpty()) {
      throw out_of_range("스택이 비어 있습니다.");
    }
    return elements.back();
  }

  // 스택이 비어 있는지 확인
  bool isEmpty() const {
    return elements.empty();
  }

  // 스택의 크기 반환
  size_t size() const {
    return elements.size();
  }
};

int main(int argc, char const* argv[]) {
  Stack<int> intStack;  // int 타입 스택 생성

  cout << "스택이 비어 있나요? " << (intStack.isEmpty() ? "예" : "아니오")
       << "\n";

  intStack.push(10);
  intStack.push(20);
  intStack.push(30);

  cout << "스택 크기: " << intStack.size() << "\n";
  cout << "최상단 값: " << intStack.top() << "\n";

  cout << "꺼낸 값: " << intStack.pop() << "\n";
  cout << "새로운 최상단 값: " << intStack.top() << "\n";

  Stack<string> stringStack;
  stringStack.push("안녕");
  stringStack.push("하세요");
  cout << "문자열 스택 최상단 값: " << stringStack.top() << "\n";

  return 0;
}
