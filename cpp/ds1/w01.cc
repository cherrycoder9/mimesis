#include <iostream>
#include <print>
#include <string>

// g++ w01.cc -std=c++23 -Wextra -march=native -fstack-protector-strong
using std::cout, std::u8string, std::println;

int main(int argc, char const *argv[]) {
    // # C++20에서 추가된 UTF-8 문자 자료형
    char8_t utf8_char = u8'A';
    // std::ostream은 기본적으로 char기반 출력을 지원하도록 설계됨
    // 따라서 char8_t에 대한 직접적인 출력 연산자는 표준에 포함되어 있지 않음
    // 그래서 아래 코드는 오류 남
    // cout << utf8_char << "\n";

    // 방법1. char로 캐스팅
    cout << static_cast<char>(utf8_char) << "\n";

    // 방법2. unsigned char로 캐스팅
    cout << static_cast<unsigned char>(utf8_char) << "\n";

    // 방법3. std::u8string으로 변환후 출력, string 헤더 필요
    u8string utf8_str = u8string(1, utf8_char);  // char8_t를 u8string으로 변환
    cout << reinterpret_cast<const char *>(utf8_str.c_str())
         << "\n";  // char*로 캐스팅

    // # C++20에서 UTF-8 문자열을 표현하기 위해 std::u8string이 도입됨
    // std::u8string은 std::basic_string<char8_t>의 별칭임
    u8string hello = u8"안녕하세요!";
    cout << reinterpret_cast<const char *>(hello.c_str()) << "\n";

    // * c_str()의 역할
    // c_str은 std::string, std::u8string과 같은 std::basic_string 계열 클래스
    // 멤버 함수 이 함수는 문자열 객체의 내부 데이터를 C 스타일 문자열(\0종료된
    // 문자 배열) 즉, const char* 또는 const char8_t 형태로 반환함
    // std::u8string은 UTF-8 문자열을 저장하는 타입으로 내부적으로 char8_t
    // 배열로 데이터 관리 하지만 std::cout은 std::basic_ostream 클래스 객체로,
    // 기본적으로 char 타입 기반 문자열을 출력하도록 설계되었음. std::u8string
    // 자체를 직접 출력하도록 오버로드된 연산자가 표준에 없기 때문에,
    // std::u8string 내용을 std::cout으로 출력하려면 내부 데이터를 char8_t*
    // 형태로 추출해야 함 c_str을 호출하면 const char8_t* 포인터를 반환해
    // 문자열의 시작 주소를 제공함

    // * reinterpret_cast<const char*>의 역할
    // C++에서 제공하는 캐스팅 연산자 중 하나로, 포인터나 참조 타입을 다른
    // 타입으로 강제로 재해석함. 타입안전성을 보장하지 않으며, 주로 저수준
    // 작업에서 사용됨 hello.c_str()은 const char8_t* 타입을 반환했음. UTF-8
    // 데이터를 가리킨 포인터 하지만 std::cout의 << 연산자는 const char*
    // 타입(즉, 일반 C스타일 문자열)을 받아서 출력하도록 오버로드되어 있음.
    // const char8_t*에 대한 출력 연산자는 표준 라이브러리에 정의되어 있지 않음
    // 따라서 const char8_t*를 const char*로 변환해야 std::cout이 이것을
    // 정상적으로 처리할 수 있음. reinterpret_cast<const char*>는 이 타입 변환을
    // 수행하며, 메모리 레이아웃은 그대로 유지한 채 타입만 재해석하는 것임

    // * 다른 캐스트는 안되나?
    // static_cast: 타입 사이의 안전한 변환(예: 상속 관계)을 위해 사용되지만,
    // char8_t와 char는 서로 다른 기본 타입이므로 불가능함 dynamic_cast: 다형성
    // 객체 변환에 사용되며, 여기에는 해당되지 않음 const_cast: const 속성만
    // 변경할 수 있지 타입 자체를 변경하지는 않음

    // * 또다른 대안은?
    // C++23에서는 <print> 헤더와 std::print가 도입되어
    // UTF-8 문자열을 더 자연스럽게 출력할 수 있음
    // 외부 라이브러리인 ICU나 Boost.Locale을 사용할 수 있음
    // char기반 std::string 으로 변환후 출력해도 됨

    const char8_t *hello2 = u8"안녕!!";
    u8string name = u8"박보람";
    println("{}", reinterpret_cast<const char *>(hello2));
    println("{}", reinterpret_cast<const char *>(name.c_str()));

    return 0;
}
