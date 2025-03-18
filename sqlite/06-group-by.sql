-- 데이터를 의미 있게 묶기 - GROUP BY
CREATE TABLE sales (store TEXT, product TEXT, amount INTEGER);

INSERT INTO
  sales (store, product, amount)
VALUES
  ('강남점', '커피', 5000),
  ('강남점', '커피', 6000),
  ('강남점', '케이크', 15000),
  ('홍대점', '커피', 4500),
  ('홍대점', '케이크', 12000),
  ('홍대점', '케이크', 13000);

-- 각 매장의 총 매출
SELECT
  store,
  SUM(amount) AS total_sales
FROM
  sales
GROUP BY
  store;

-- 조건을 추가로 걸어 그룹 필터링하기 - HAVING
-- 그룹에 조건을 걸고 싶을때 씀, 매출이 27000원 이상인 매장만 보려면?
-- WHERE, HAVING을 혼동하지 말것 
-- 조건의 대상이 원본 데이터인 경우 WHERE, 묶인 후 집계 결과이면 HAVING
SELECT
  store,
  SUM(amount) AS total_sales
FROM
  sales
GROUP BY
  store
HAVING
  total_sales >= 27000;

-- 홍대점의 데이터만 가지고 그룹화
SELECT
  product,
  SUM(amount) AS total_sales
FROM
  sales
WHERE
  store = '홍대점'
GROUP BY
  product;