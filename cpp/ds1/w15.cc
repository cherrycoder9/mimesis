
// # 연산자 오버로딩
// 기존의 연산자를 사용자 정의 클래스에 맞게 재정의함
// 직관적이고 자연스러운 방식으로 객체를 조작할 수 있게 함
// C++ 표준에서 허용된 연산자만 가능함 (::, sizeof, typeid 등은 불가)

// 연산 결과로 어떤 값을 반환할지 신중히 결정해야 함
// 보통 연산 결과를 새로운 객체로 반환하거나 참조를 반환함
// 이항 연산자의 경우 좌항과 우항의 위치를 고려해야 함.

// 연산자 오버로딩은 두가지 방식으로 구현됨
// 멤버함수: 연산자를 클래스 내부에서 정의하며 좌항은
// 호출하는 객체(this)로 암시적으로 전달됨
// 전역함수: 클래스 외부에서 정의하며 좌항과 우항 모두 매개변수로
// 명시적으로 전달됨. friend 키워드를 사용해 클래스 private 멤버 접근 가능

#include <iostream>

using namespace std;

class Point {
 private:
  int x;
  int y;

 public:
  // 생성자, x와 y좌표 초기화
  Point(int xCoord = 0, int yCoord = 0) : x(xCoord), y(yCoord) {
  }

  // + 연산자 오버로딩 (멤버함수)
  Point operator+(const Point& other) const {
    return Point(x + other.x, y + other.y);  // 새로운 Point 객체 반환
  }

  // 좌표 출력 함수
  void print() const {
    cout << "x: " << x << ", y: " << y << "\n";
  }
};

int main(int argc, char const* argv[]) {
  // 두개의 Point 객체 생성
  Point p1(3, 4);
  Point p2(1, 2);

  // + 연산자를 사용해 두 좌표를 더함
  Point result = p1 + p2;

  // 결과 출력
  cout << "두 좌표의 합 결과: \n";
  result.print();
  return 0;
}
