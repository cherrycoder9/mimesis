DROP TABLE IF EXISTS customers;

CREATE TABLE customers (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL
);

DROP TABLE IF EXISTS orders;

CREATE TABLE orders (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  customer_id INTEGER NOT NULL,
  product TEXT NOT NULL,
  FOREIGN KEY (customer_id) REFERENCES customers (id)
);

INSERT INTO
  customers (name)
VALUES
  ('수지'),
  ('지훈'),
  ('민서');

INSERT INTO
  orders (customer_id, product)
VALUES
  (1, '노트북'),
  (2, '이어폰'),
  (2, '키보드');

-- 교집합으로 데이터 연결 - INNER JOIN
SELECT
  customers.name,
  orders.product
FROM
  customers
  INNER JOIN orders ON customers.id = orders.customer_id;

-- 왼쪽 테이블 중심으로 데이터 연결 - LEFT JOIN
-- 왼쪽 테이블 데이터는 무조건 모두 표시됨
-- 오른쪽 테이블 데이터가 없으면 NULL로 채움 
SELECT
  customers.name,
  orders.product
FROM
  customers
  LEFT JOIN orders ON customers.id = orders.customer_id;

-- SQLite는 RIGHT JOIN을 직접 지원하지는 않음 
-- 동일한 효과를 얻으려면 LEFT JOIN의 테이블 위치를 바꾸면 됨 
SELECT
  customers.name,
  orders.product
FROM
  orders
  LEFT JOIN customers ON orders.customer_id = customers.id;

-- 모든 데이터를 전부 엮어버리기 - CROSS JOIN
-- 흔히 카르테시안 곱(cartesian product)이라고 부른다
-- 데이터가 폭발적으로 늘어나므로 주의한다 
SELECT
  customers.name,
  orders.product
FROM
  customers
  CROSS JOIN orders;

/*
INNER JOIN은 교집합(∩)
공통된 데이터만 남긴다.

LEFT JOIN은 왼쪽 집합 우선 (왼쪽 집합 ⊂ 오른쪽 집합)
왼쪽은 전부 남기고 오른쪽만 골라 연결한다.

RIGHT JOIN은 오른쪽 집합 우선 (왼쪽 집합 ⊃ 오른쪽 집합)
오른쪽은 전부 남기고 왼쪽만 골라 연결한다.

CROSS JOIN은 카르테시안 곱 (모든 조합)
모든 데이터를 무작정 연결한다.
*/