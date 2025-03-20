import { Coffee } from "../coffeeFactory";

// Latte 클래스는 Coffee 인터페이스를 구현
export class Espresso implements Coffee {
  public brew(): void {
    console.log("우유를 붓고 에스프레소를 추가해 라떼를 만듭니다.");
  }
}