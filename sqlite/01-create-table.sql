-- PRIMARY KEY 각 행을 고유하게 식별하는 값 
-- AUTOINCREMENT 자동으로 증가하는 기본 키
-- UNIQUE 중복되지 않아야 하는 값
-- NOT NULL 반드시 값이 있어야 함 
-- CHECK 특정 조건을 만족해야 함
-- DEFAULT 기본값 지정
-- users 테이블 생성 
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    age INTEGER CHECK (age >= 0),
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 테이블 삭제
DROP TABLE IF EXISTS users;

-- 데이터 삽입
INSERT INTO
    users (name, email, age)
VALUES
    ('김철수', '철수@hello.com', 70),
    ('박보라', '보라@hello.com', 20);

-- 데이터 조회
SELECT
    *
FROM
    users;

-- 외래 키 사용 
-- orders 테이블이 users 테이블을 참조하도록 만들어보자
-- SQLite는 동적 타입 시스템을 사용하므로 REAL 필드에 정수를 넣어도 오류 발생하지 않음
-- 그러나 INTEGER PRIMARY KEY는 반드시 정수만 저장해야 함 
-- TEXT 문자열, BLOB 바이너리 데이터
CREATE TABLE orders (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER,
    product TEXT NOT NULL,
    amount REAL NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO
    orders (user_id, product, amount)
VALUES
    (1, '카페모카', 1),
    (2, '아메리카노', 2),
    (1, '카페모카', 1),
    (2, '카페모카', 2);

SELECT
    *
FROM
    orders;

-- 카페모카 주문한 사람들 이름 출력 
SELECT
    users.name
FROM
    users
    JOIN orders ON users.id = orders.user_id
WHERE
    orders.product = '카페모카';

-- 카페모카를 여러 잔 주문한 경우 출력 이름 중복 제거하려면 DISTINCT 추가 
SELECT DISTINCT
    users.name
FROM
    users
    JOIN orders ON users.id = orders.user_id
WHERE
    orders.product = '카페모카';

-- 주문수량, 주문일시 함께 출력
SELECT
    users.name,
    orders.amount,
    orders.created_at
FROM
    users
    JOIN orders ON users.id = orders.user_id
WHERE
    orders.product = '카페모카';