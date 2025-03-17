DROP TABLE IF EXISTS members;

CREATE TABLE members (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  age INT NOT NULL,
  city TEXT NOT NULL
);

INSERT INTO
  members (name, age, city)
VALUES
  ('민수', 21, '서울'),
  ('서연이', 23, '서울'),
  ('지호', 27, '인천'),
  ('하윤정김', 32, '인천');

-- 회원을 나이순으로 정렬 (나이가 적은 순서)
-- 오름차순(ASC, 생략가능), 내림차순(DESC)
SELECT
  name,
  age
FROM
  members
ORDER BY
  age;

-- 나이가 많은 순서로 정렬 
SELECT
  name,
  age
FROM
  members
ORDER BY
  age DESC;

-- 다중 컬럼으로 정렬하기
-- 컬럼 여러개를 쉼표로 구분해서 정렬 순서를 지정
SELECT
  name,
  city,
  age
FROM
  members
ORDER BY
  city ASC,
  age DESC;

-- 표현식과 함수의 결과값으로도 정렬 가능
-- 이름이 짧은 순서대로 정렬하기 
SELECT
  name,
  LENGTH(name) AS 이름길이
FROM
  members
ORDER BY
  LENGTH(name) ASC;

-- NULL 값 처리방식 주의
-- SQLite에서 NULL은 오름차순일때 가장 먼저 오고 내림차순일땐 맨 뒤에 위치
-- NULL 정렬을 원하는 대로 하려면 IS NULL 표현식 활용이 필요함