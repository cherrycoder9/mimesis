// 어댑터 클래스
import { OldPaymentGateway } from "../payment-gateways/OldPaymentGateway";
import { PaymentGateway } from "../payment-gateways/PaymentGateway";

// 기존 결제 시스템을 새 인터페이스에 맞게 변환
export class PaymentAdapter implements PaymentGateway {
  private oldGateway: OldPaymentGateway;

  constructor(oldGateway: OldPaymentGateway) {
    this.oldGateway = oldGateway; // 기존 결제 시스템 인스턴스를 저장 
  }

  // PaymentGateway 인터페이스의 메서드를 구현
  // 내부적으로 기존 시스템의 메서드를 호출함
  public makePayment(amount: number): void {
    this.oldGateway.processPayment(amount);
  }
}