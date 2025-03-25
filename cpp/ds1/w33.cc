// RAII는 리소스 획득 시점을 객체의 초기화 시점과 일치시키는 것을 의미함
// 여기서 리소스란 메모리, 파일 핸들, 네트워크 소켓, 뮤텍스 등 프로그램 실행중
// 관리해야 하는 모든 자원을 포괄함

// RAII의 핵심 아이디어는 다음과 같음
// 1. 리소스 캡슐화, 리소스를 관리하는 객체를 정의하고 해당 객체가 리소스의
// 소유권을 가짐
// 2. 생성시 할당, 객체의 생성자에서 리소스를 할당하며, 할당 실패시 예외를
// 발생시켜 객체가 불완전한 상태로 생성되지 않도록 함
// 3. 소멸시 해제, 객체가 스코프를 벗어나 소멸될때 소멸자에서 리소스를 자동으로
// 해제함

// 파일 핸들을 관리하는 클래스를 예로 들어 RAII 설명
// 파일을 열고 닫는 작업을 RAII를 통해 자동화한 예제

#include <cstdlib>
#include <fstream>
#include <iostream>
#include <stdexcept>

using namespace std;

class FileHandler {
 private:
  ofstream file;  // 관리할 파일 스트림 객체

 public:
  // 생성자, 파일을 열고 초기화
  // explicit: c++ 컴파일러는 특정 상황에서 자동으로 한 타입을 다른 타입으로
  // 변환할수 있기에 이 키워드를 붙이면 해당 생성자를 통한 객체 생성은 명시적
  // 호출로만 가능해짐
  explicit FileHandler(const string& filename) : file(filename) {
    if (!file.is_open()) {
      throw runtime_error("파일 열기 실패: " + filename);
    }
    cout << "파일 '" << filename << "'이 성공적으로 열렸습니다." << endl;
  }

  // 소멸자
  ~FileHandler() {
    if (file.is_open()) {
      file.close();
      cout << "파일이 성공적으로 닫혔습니다." << endl;
    }
  }

  // 파일에 데이터 쓰기
  void write(const string& data) {
    if (file.is_open()) {
      file << data << endl;
      cout << "데이터 '" << data << "'이 파일에 기록되었습니다." << endl;
    } else {
      cerr << "오류: 파일이 열려 있지 않습니다." << endl;
    }
  }

  // 복사 생성자 삭제
  FileHandler(const FileHandler&) = delete;
  // 복사 대입 연산자 삭제
  FileHandler& operator=(const FileHandler&) = delete;
};

int main(int argc, char const* argv[]) {
  try {
    FileHandler fh("test.txt");      // 파일 열기
    fh.write("RAII 테스트 데이터");  // 데이터 쓰기
    // 스코프 종료시 fh 소멸, 파일 자동 닫힘
  } catch (const exception& e) {
    cerr << "예외 발생: " << e.what() << '\n';
    return EXIT_FAILURE;
  }
  cout << "프로그램이 정상적으로 종료됩니다." << endl;
  return EXIT_SUCCESS;
}
