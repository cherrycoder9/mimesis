-- 테이블을 생성한 후에도 요구사항 변화나 설계 개선을 위해 테이블 구조를 수정하는 경우가 빈번함
-- 이때 ALTER TABLE을 사용하며 테이블 이름변경, 새로운 열 추가, 열 이름 변경, 열 삭제 수행가능
-- 그러나 SQLite의 ALTER TABLE 명령은 MySQL, PostgreSQL 등 다른 DBMS에 비해 기능 제한적
-- 열의 데이터 타입 변경이나 제약 조건 수정은 직접 지원되지 않으므로 
-- 테이블을 새로 생성하고 기존 데이터를 복사하는 우회적인 방법을 사용해야 함
CREATE TABLE students (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  major TEXT,
  grade TEXT
);

INSERT INTO
  students (name, major, grade)
VALUES
  ('박영호', '컴퓨터공학', '3.8'),
  ('김희철', '통계', '3.6'),
  ('박보람', '컴퓨터공학', '3.7'),
  ('윤수아', '컴퓨터공학', '3.5'),
  ('정혜민', '통계', '3.9'),
  ('박혜경', '통계', '4.3 '),
  ('나동규', '통계', '4.1'),
  ('이동수', '컴퓨터공학', '4.0');

SELECT
  *
FROM
  students;

-- 기존 테이블 students를 student로 이름 변경
ALTER TABLE students
RENAME TO student;

SELECT
  *
FROM
  student;

-- 테이블에 새로운 속성을 추가해야 할때 ALTER TABLE 명령에 ADD COLUMN 절을 사용
-- 추가된 열에 기본값 지정 가능, 미지정시 NULL로 채워짐 
-- student 테이블에 age라는 INTEGER 타입 열 추가, 기본값은 20
ALTER TABLE student
ADD COLUMN age INTEGER DEFAULT 20;

SELECT
  *
FROM
  student;

-- 열 이름 변경, SQLite 3.25.0 버전부터 지원
ALTER TABLE student
RENAME COLUMN major to department;

SELECT
  *
FROM
  student;

-- 열 삭제, 3.35.0 버전부터 사용 가능
-- 해당 열에 의존하는 뷰, 트리거, 인덱스 등이 있으면 오류발생하므로 주의 
ALTER TABLE student
DROP COLUMN age;

SELECT
  *
FROM
  student;

-- 열의 데이터 타입 변경 
CREATE TABLE students (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT,
  department TEXT,
  grade REAL
);

-- 데이터 복사 (TEXT를 REAL로 변환)
INSERT INTO
  students (name, department, grade)
SELECT
  name,
  department,
  CAST(grade AS REAL)
FROM
  student;

-- 새 테이블 확인
SELECT
  *
FROM
  students;

-- 기존 테이블 삭제
DROP TABLE student;

-- 새로운 테이블 이름을 기존 이름으로 변경
ALTER TABLE students
RENAME TO student;

-- 다시 확인 
SELECT
  *
FROM
  student;

-- 기존 TEXT에서 4.0 점이 REAL에서는 4점으로 나옴.
-- 굳이 소수점까지 표시하려면 printf 함수 사용, 3.8.3 이상에서 지원 
SELECT
  name,
  department,
  printf('%.1f', grade) AS grade
FROM
  student;