-- SQLite는 다른 RDBMS보다 제한적인 ALTER TABLE 기능을 가짐 
-- # SQLite는 아래 기능만을 공식적으로 지원 
-- 컬럼 추가 (ADD COLUMN)
-- 테이블 이름 변경 (RENAME TO)
-- 컬럼 이름 변경 (RENAME COLUMN)
-- 테이블 합치기 (ALTER) TABLE ... MERGE), SQLite 3.35.0 이상 
-- # 아래와 같은 변경은 ALTER TABLE에서 직접 지원되지 않음 
-- 컬럼 삭제 (DROP COLUMN)
-- 컬럼 타입 변경
-- 컬럼 타입 변경
-- 컬럼 순서 변경  
-- # 기존 컬럼에 제약 조건 추가  
-- 주의사항: ADD COLUMN은 항상 테이블의 마지막 컬럼에 추가됨
-- 새로운 컬럼에는 기본적으로 NULL 값이 들어감
-- NOT NULL이나 DEFAULT를 추가하지 않으면 기존 행에는 NULL이 들어감
ALTER TABLE users
ADD COLUMN password TEXT NOT NULL DEFAULT 'WELCOME3';

SELECT
    *
FROM
    users;

-- # 테이블 이름 변경
-- 테이블 이름 변경시 기존 데이터는 유지됨
-- 변경후 뷰(View)나 트리거(Trigger)도 함께 업데이트 해야함 
ALTER TABLE users
RENAME TO customers;

SELECT
    *
FROM
    customers;

SELECT DISTINCT
    customers.name
FROM
    customers
    JOIN orders ON customers.id = orders.user_id
WHERE
    orders.product = '카페모카';

-- # 컬럼 이름 변경
ALTER TABLE customers
RENAME COLUMN age TO birth_year;

SELECT
    *
FROM
    customers;

CREATE TABLE dormant_customers (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    birth_year INTEGER CHECK (birth_year >= 0),
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL DEFAULT 'WELCOME1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE dormant_customers;

INSERT INTO
    dormant_customers (name, birth_year, email, password)
VALUES
    ('차오름', 33, '오름@hi.com', 'good77');

SELECT
    *
FROM
    dormant_customers;

-- # 테이블 합치기 (INSERT INTO ... SELECT ...)
INSERT INTO
    dormant_customers (name, birth_year, email)
SELECT
    name,
    birth_year,
    email
FROM
    customers;

SELECT
    *
FROM
    dormant_customers;

ALTER TABLE customers
ADD COLUMN password TEXT NOT NULL DEFAULT 'WELCOME3';

CREATE TABLE abc (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE def (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

INSERT INTO
    abc (name)
VALUES
    ('가갸'),
    ('거겨');

INSERT INTO
    def (name)
VALUES
    ('나냐'),
    ('너녀');

SELECT
    *
FROM
    abc;

SELECT
    *
FROM
    def;

INSERT INTO
    abc (name)
SELECT
    name
FROM
    def;

SELECT
    *
FROM
    abc;

-- 컬럼 삭제 기능이 없으므로 새 테이블 생성 -> 필요한 컬럼만 복사 -> 기존 테이블을 삭제하고 새 테이블 이름 변경을 하는 방식으로 하면 됨