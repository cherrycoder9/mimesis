import { Coffee } from "../coffeeFactory";

// Espresso 클래스는 Coffee 인터페이스를 구현
export class Latte implements Coffee {
  public brew(): void {
    console.log("에스프레소 샷을 추출해 에스프레소를 만듭니다.");
  }
}