import { Coffee } from "../coffeeFactory";

// Americano 클래스는 Coffee 인터페이스를 구현
export class Americano implements Coffee {
  public brew(): void {
    console.log("물을 붓고 에스프레소를 추가해 아메리카노를 만듭니다.");
  }
}