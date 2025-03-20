import { PaymentGateway } from "./PaymentGateway";

// 새로운 결제 게이트웨이
export class NewPaymentGateway implements PaymentGateway {
  public makePayment(amount: number): void {
    console.log(`새로운 결제 시스템으로 ${amount}원을 결제합니다.`);

  }
}