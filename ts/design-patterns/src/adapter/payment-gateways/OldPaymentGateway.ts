// 기존 결제 게이트웨이
export class OldPaymentGateway {
  public processPayment(amount: number): void {
    console.log(`기존 결제 시스템으로 ${amount}원을 결제합니다.`);
  }
}