CREATE TABLE employees (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  salary INTEGER NOT NULL, -- INTEGER 대신 REAL 권장 
  title TEXT NOT NULL
);

INSERT INTO
  employees (name, salary, title)
VALUES
  ('김대환', 98000, 'Backend Engineer'),
  ('정희철', 86000, 'Backend Engineer'),
  ('박세희', 93000, 'Frontend Engineer'),
  ('이대승', 77000, 'Backend Engineer'),
  ('구유라', 95000, 'Frontend Engineer');

-- id가 1인 직원의 급여를 99000으로 수정
UPDATE employees
SET
  salary = 99000
WHERE
  id = 1;

-- 다중 열 수정 
-- id가 2인 직원의 급여와 직합을 동시에 수정
UPDATE employees
SET
  salary = 88000,
  title = 'ML Engineer'
WHERE
  id = 2;

-- WHERE 절을 생략하면 모든 행이 수정되니 주의
SELECT
  *
FROM
  employees;

-- UPDATE 문은 트랜잭션 내에서 실행될 수 있어 데이터 무결성을 보장 가능
BEGIN TRANSACTION;

UPDATE employees
SET
  salary = 79000
WHERE
  ID = 4;

SELECT
  *
FROM
  employees;

-- 오류 발생시 롤백됨 
COMMIT;

SELECT
  *
FROM
  employees;