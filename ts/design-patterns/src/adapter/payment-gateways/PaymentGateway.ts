// 결제 게이트웨이 인터페이스
// 모든 결제 시스템은 이 인터페이스를 따라야 함
export interface PaymentGateway {
  makePayment(amount: number): void; // 결제 금액을 받아 처리하는 메서드 
}