/**
 * 스터디카페의 예약을 관리하는 싱글톤 클래스 
 */
export class ReservationManager {
  // 유일한 인스턴스를 저장하는 정적 변수 (초기값은 null)
  private static instance: ReservationManager | null = null;

  // 예약 정보를 저장하는 배열
  private reservations: string[] = [];

  // private 생성자, 외부에서 직접 인스턴스를 생성하지 못하도록 제한
  private constructor() {
    console.log("예약관리자 인스턴스가 생성되었습니다.");
  }

  /**
   * 싱글톤 인스턴스를 반환하는 메서드
   *
   * @returns {ReservationManager} 
   */
  public static getInstance(): ReservationManager {
    if (ReservationManager.instance === null) {
      ReservationManager.instance = new ReservationManager();
    }
    return ReservationManager.instance;
  }

  /**
   * 새로운 예약을 추가하는 메서드
   *
   * @param {string} reservation 
   */
  public addReservation(reservation: string): void {
    this.reservations.push(reservation);
    console.log(`예약이 추가되었습니다: ${reservation}`);
  }

  public listReservations(): void {
    console.log("현재 예약 목록:");
    if (this.reservations.length === 0) {
      console.log("예약이 없습니다.");
    } else {
      this.reservations.forEach((reservation, index) => {
        console.log(`${index + 1}. ${reservation}`);
      });
    }

  }
}

