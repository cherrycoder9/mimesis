// # 빌더 패턴
// 복잡한 객체를 단계적으로 구성하게 하는 패턴, 생성 과정을 추상화

import { ProductBuilder } from "./ProductBuilder";

function main() {
  const product = new ProductBuilder()
    .setName('트렌디 후드티')
    .setPrice(39000)
    .setCategory('상의')
    .setDescription('따뜻하고 멋진 후드티입니다.')
    .build();

  // 생성된 제품 정보 출력
  console.log('제품명: ' + product.getName());
  console.log('가격: ' + product.getPrice() + '원');
  console.log('카테고리: ' + product.getCategory());
  console.log('설명: ' + product.getDescription());

}

main();