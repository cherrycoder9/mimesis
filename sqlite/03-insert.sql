-- NULL을 허용하지 않는 컬럼에 NULL 값을 넣으면 에러 발생 
-- DEFAULT 값이 설정된 컬럼은 자동 채워짐
-- 한번에 여러개의 데이터를 삽입하면 단일 트랜잭션으로 처리되어 COMMIT 비용 절감되고 디스크 I/O 감소
-- INSERT INTO ... SELECT 이용하면 다른 테이블에서 데이터를 가져와 삽입할 수 있음 
-- 기존 데이터를 백업하거나 필터링해서 특정 조건의 데이터만 복사 가능 
-- INSERT 대신 REPLACE 사용할 수도 있음. 동일한 PRIMARY KEY나 UNIQUE 키가 존재하면 기존 데이터를 삭제하고 새로운 데이터를 삽입함. 기존 데이터를 완전히 삭제한 후 다시 넣기 때문에, 자동 증가 값이 증가할 수 있음
REPLACE INTO
    users (id, name, age)
VALUES
    (1, '박보람', 22);

-- SQLite 에서는 OR 절을 사용해 충돌시 동작을 제어할 수 있음
INSERT OR IGNORE INTO
    users (id, name, age)
VALUES
    (1, '박보람', 22);

-- UPSERT (INSERT ... ON CONFLICT)
-- 데이터가 없으면 삽입하고, 있으면 업데이트하는 문법
-- INSERT: 기존 데이터가 있으면 에러 발생 
-- INSERT OR REPLACE: 기존 데이터를 삭제후 새 데이터 삽입
-- INSERT OR IGNORE: 기존 데이터가 있으면 무시하고 아무 작업 안함
-- UPSERT: 기존 데이터가 있으면 업데이트
-- id = 1 없으면 새롭게 삽입되고 이미 존재하면 기존 데이터 업데이트 
INSERT INTO
    users (id, name, age)
VALUES
    (1, '박보람', 25)
ON CONFLICT (id) DO UPDATE
SET
    name = EXCLUDED.name,
    age = EXCLUDED.age;

-- 충돌시 아무 작업도 하지 않고 무시하려면 DO NOTHING 사용도 가능
INSERT INTO
    users (id, name, age)
VALUES
    (1, '박보람', 25)
ON CONFLICT (id) DO NOTHING;

-- 업데이트가 특정 조건일때만 실행되도록 WHERE을 추가할수도 있음
-- 현재 age 값보다 새로운 age 값이 클 때만 업데이트  
INSERT INTO
    users (id, name, age)
VALUES
    (1, '박보람', 25)
ON CONFLICT (id) DO UPDATE
SET
    AGE = EXCLUDED.age
WHERE
    users.age < EXCLUDED.age;

-- IGNORE: 충돌이 발생하면 아무 작업도 하지 않고 무시함
-- REPLACE: 기존 데이터를 삭제하고 새 데이터 삽입 
-- ROLLBACK / ABORT / FAIL 옵션도 있음 
-- AUTOINCREMENT는 INTEGER PRIMARY KEY 에서만 동작하며 자동증가하는 유일한 값을 부여함 
-- id값(AUTOINCREMENT)에 NULL 입력해도 자동 증가함 
-- 삭제후에도 사용했던 숫자는 재사용되지 않음
-- BEGIN TRANSACTION, 여러개의 INSERT를 한번의 트랜잭션으로 묶기
BEGIN TRANSACTION;

INSERT INTO
    users (name, age)
VALUES
    ('강민철', 29);

INSERT INTO
    users (name, age)
VALUES
    ('배민수', 21);

INSERT INTO
    users (name, age)
VALUES
    ('정유리', 32);

COMMIT;