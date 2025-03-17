import { ReservationManager } from "./ReservationManager";

// 싱글톤 인스턴스를 두번 가져와 동일성 테스트
const manager1 = ReservationManager.getInstance();
const manager2 = ReservationManager.getInstance();

// 동일 인스턴스인지 확인
console.log(`동일 인스턴스인가?: ${manager1 === manager2}`);

// 예약 추가 (두개가 같은 인스턴스를 공유하므로 데이터가 통합됨)
manager1.addReservation("3월 27일, 오후 2시 - 3명");
manager2.addReservation("3월 29일, 오후 8시 - 2명");

// 예약 목록 출력
manager1.listReservations();