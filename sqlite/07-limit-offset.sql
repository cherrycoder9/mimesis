-- 결과집합(result set)의 행(row) 수를 제한 - LIMIT
-- 여기에 OFFSET이 결합하면, 데이터를 어디서부터 가져올지 출발 지점을 정할수 있음
CREATE TABLE products (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  created_at TEXT DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO
  products (name)
VALUES
  ('태블릿1'),
  ('태블릿2'),
  ('태블릿3'),
  ('태블릿4'),
  ('태블릿5'),
  ('태블릿6'),
  ('태블릿7'),
  ('태블릿8'),
  ('태블릿9'),
  ('태블릿10'),
  ('태블릿11'),
  ('태블릿12'),
  ('태블릿13'),
  ('태블릿14'),
  ('태블릿15');

-- 테이블에서 10개만 데이터를 가져옴 
SELECT
  *
FROM
  products
LIMIT
  10;

-- 11번째부터 10개의 데이터를 가져옴 
-- 처음 10개를 건너뛴후 다음 데이터부터 가져옴 
SELECT
  *
FROM
  products
LIMIT
  10
OFFSET
  10;

-- OFFSET의 숫자가 커지면 성능이 저하되므로 주의 
-- 대안1: 인덱스를 이용한 키 기반 페이지네이션 구현, 단점: 특정 페이지로 바로 점프하기 어려움
-- 대안2: 범위를 좁혀 OFFSET 사용하기, 정렬된 상태에서 검색 범위를 좁힌 다음 사용
-- ORDER BY를 함께 사용하지 않으면 데이터가 항상 일정한 순서로 나온다는 보장 없음