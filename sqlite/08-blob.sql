-- Blob: Binary Large Object
-- 이진 데이터를 저장하기 위해 설계된 데이터 타입 
-- 이미지, 오디오, 비디오, 문서파일 등과 같은 비구조화된 데이터를 DB에 직접 저장
-- 파일 시스템에 의존하지 않고 DB 내에서 모든 데이터를 통합적으로 관리 가능
-- DB가 제공하는 ACID 속성을 통해 데이터 무결성 보장 
-- SQLite에서 Blob 데이터 타입 자체에 명시적인 크기 제한이 없음
-- 블롭 데이터를 DB에 저장하면 파일 시스템 접근에 비해 성능 저하될 가능성 있음 
CREATE TABLE files (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  data BLOB NOT NULL
);