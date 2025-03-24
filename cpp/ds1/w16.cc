// 전역 함수로 오버로딩 예제
// 전역 함수로 연산자 오버로딩이 필요한 경우
// 1. 좌항이 클래스 객체가 아닌 경우, 전역함수는 좌항과 우항을 자유롭게 정의가능
// 2. 대칭성을 보장하고 싶을 때, 위와 비슷한 맥락
// 3. 클래스 설계를 최소화하고 싶을 때
// 4. friend 함수로 private, protected 멤버에 접근해야 할때
// 5. 서드파티 클래스에 연산자를 추가하고 싶을 때

#include <iostream>

using namespace std;

class Point {
 private:
  int x;
  int y;

 public:
  Point(int xCoord = 0, int yCoord = 0) : x(xCoord), y(yCoord) {
  }
  void print() const {
    cout << "x: " << x << ", y: " << y << "\n";
  }
  // 전역 함수에서 private 멤버에 접근하기 위해 friend 선언
  friend Point operator+(const Point& lhs, const Point& rhs);
};

// + 연산자 오버로딩 (전역 함수)
// left-hand side, right-hand side
Point operator+(const Point& lhs, const Point& rhs) {
  return Point(lhs.x + rhs.x, lhs.y + rhs.y);
}

int main(int argc, char const* argv[]) {
  /* code */
  return 0;
}
