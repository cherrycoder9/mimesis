// 어댑터 패턴은 서로 호환되지 않는 인터페이스를 가진 클래스들이
// 함께 작동할 수 있도록 중간에서 변환 역할을 하는 디자인 패턴

import { PaymentAdapter } from "./adapters/PaymentAdapter";
import { NewPaymentGateway } from "./payment-gateways/NewPaymentGateway";
import { OldPaymentGateway } from "./payment-gateways/OldPaymentGateway";
import { PaymentGateway } from "./payment-gateways/PaymentGateway";

// 클라이언트 코드, 결제 처리 로직
function processPayment(gateway: PaymentGateway, amount: number): void {
  console.log(`결제 요청: ${amount}원`);
  gateway.makePayment(amount); // 결제 게이트웨이를 통해 결제 실행
  console.log("결제 완료!\n");
}

// 새로운 결제 시스템 사용
const newGateway = new NewPaymentGateway();
processPayment(newGateway, 10000);

// 기존 결제 시스템 사용 (어댑터 적용)
const oldGateway = new OldPaymentGateway();
const adapter = new PaymentAdapter(oldGateway);
processPayment(adapter, 20000);