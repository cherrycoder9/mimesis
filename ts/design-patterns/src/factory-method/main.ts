// # 팩토리 메서드
// 객체 생성을 직접 호출하지 않고, 특정 메서드를 통해 생성 책임을
// 하위 클래스나 별도의 팩토리 로직에 위임하는 방식

import { AmericanoFactory, CoffeeFactory, EspressoFactory, LatteFactory } from "./coffeeFactory";

function main() {
  // 아메리카노 주문
  const americanoFactory: CoffeeFactory = new AmericanoFactory();
  americanoFactory.orderCoffee("아메리카노");

  // 라떼 주문
  const latteFactory: CoffeeFactory = new LatteFactory();
  latteFactory.orderCoffee("라떼");

  // 에스프레소 주문
  const espressoFactory: CoffeeFactory = new EspressoFactory();
  espressoFactory.orderCoffee("에스프레소");
}

main();