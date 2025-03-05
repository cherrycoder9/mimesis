#include <iomanip>
#include <ios>
#include <iostream>
#include <string>

using namespace std;

struct 직원 {
  string 이름;
  int 나이;
  double 급여;

  void 정보출력() const {
    // 10진수 고정, 소수점 0자리, iomanip 헤더 필요
    // cout은 스트림 객체로, 출력 형식을 제어하는 내부 상태 플래그를 가지고 있음
    // fixed와 setprecision 같은 조작자는 이 상태를 변경함
    // 한번 설정되면 그 이후의 모든 출력에 적용됨
    // fixed: 고정 소수점 표기법을 사용하도록 설정
    // setprecision(n): 소수점 이하 출력 자리 수를 n으로 설정
    cout << fixed << setprecision(0);
    cout << "이름: " << 이름 << ", 나이: " << 나이 << ", 급여: " << 급여
         << "원\n";

    // 과학적 표기법으로 되돌리려면 아래 코드 사용
    // cout << defaultfloat; // 기본 동작으로 되돌아감
  }

  void 급여인상(double 증가율) { 급여 *= (1.0 + 증가율); }
};

int main(int argc, char const *argv[]) {
  직원 employee{"김철수", 30, 5000000.0};

  employee.정보출력();
  employee.급여인상(0.1);
  employee.정보출력();

  return 0;
}
