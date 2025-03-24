import sqlite3

with sqlite3.connect('test.db') as conn:
    cursor = conn.cursor()

    # 이미지 파일을 이진 데이터로 읽기
    with open('sample.jpg', 'rb') as file:
        blob_data = file.read()

    # 테이블 생성, 이미 존재하면 생략 가능
    # 생략

    # 파일 이름과 블롭 데이터를 테이블에 저장
    cursor.execute('INSERT INTO files (name, data) VALUES (?, ?)',
                   ('sample.jpg', blob_data))

    # 변경사항 저장
    conn.commit()

    # ID가 1인 레코드 조회
    cursor.execute('SELECT name, data FROM files WHERE id = ?', (1, ))
    row = cursor.fetchone()  # 다음 한 행을 가져오는 메서드
    # 다음 행이 있으면 해당 행을 튜플로 반환, 없으면 None 반환
    # fetchmany(size): 결과 집합에서 지정한 size만큼의 행을 리스트 형태로 가져옴
    # fetchall(): 결과 집합에서 남아 있는 모든 행을 리스트 형태로 한꺼번에 가져옴

    if row:
        name, data = row
        # 조회된 블롭 데이터를 새 파일로 저장
        with open(f'retrieved_{name}', 'wb') as file:
            file.write(data)
        print(f"파일 '{name}이 저장됐습니다.")
    else:
        print("해당 ID의 데이터가 없습니다.")

print("블롭 데이터가 저장됐습니다.")
