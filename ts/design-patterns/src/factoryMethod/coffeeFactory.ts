// Coffee 인터페이스를 정의

import { Americano } from "./coffee/americano";
import { Espresso } from "./coffee/espresso";
import { Latte } from "./coffee/latte";

// 모든 커피 클래스는 이 인터페이스를 구현해야 함
export interface Coffee {
  brew(): void;
}

// 팩토리 메서드 패턴의 핵심 추상 클래스
// 서브클래스에서 createCoffee 메서드 구현해 구체적인 커피 객체 생성
export abstract class CoffeeFactory {
  // 팩토리 메서드, 서브클래스에서 구체적인 커피 객체를 생성하도록 강제 
  protected abstract createCoffee(type: string): Coffee;

  // 커피를 주문하는 메서드 
  public orderCoffee(type: string): void {
    const coffee = this.createCoffee(type); // 서브클래스에서 생성된 커피 객체
    console.log(`${type} 주문을 받았습니다.`);
    coffee.brew(); // 커피 제조 시작
  }
}

// CoffeeFactory를 상속받아 아메리카노를 생성
export class AmericanoFactory extends CoffeeFactory {
  protected createCoffee(type: string): Coffee {
    if (type === "아메리카노") {
      return new Americano();
    } else {
      throw new Error("지원하지 않는 커피 종류입니다.");
    }
  }
}

// 커피팩토리를 상속받아 라떼를 생성
export class LatteFactory extends CoffeeFactory {
  protected createCoffee(type: string): Coffee {
    if (type === "라떼") {
      return new Latte();
    } else {
      throw new Error("지원하지 않는 커피 종류입니다.");
    }
  }
}

// 커피팩토리를 상속받아 에스프레소를 생성
export class EspressoFactory extends CoffeeFactory {
  protected createCoffee(type: string): Coffee {
    if (type === "에스프레소") {
      return new Espresso();
    } else {
      throw new Error("지원하지 않는 커피 종류입니다.");
    }
  }
}