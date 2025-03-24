// Product 객체를 단계적으로 생성하기 위한 빌더 클래스
import { Product } from "./Product";

export class ProductBuilder {
  private product: Product; // 생성할 제품 객체를 저장함

  // 생성자, 새로운 Product 객체를 초기화
  constructor() {
    this.product = Product.create();
  }

  // 제품명을 설정하는 메서드
  public setName(name: string): ProductBuilder {
    this.product['name'] = name; // private 속성에 접근하기 위해 [''] 문법 사용
    return this;
  }

  // 가격을 설정하는 메서드
  public setPrice(price: number): ProductBuilder {
    this.product['price'] = price;
    return this;
  }

  // 카테고리를 설정하는 메서드
  public setCategory(category: string): ProductBuilder {
    this.product['category'] = category;
    return this;
  }

  // 제품 설명을 설정하는 메서드 
  public setDescription(description: string): ProductBuilder {
    this.product['description'] = description;
    return this;
  }

  // 완성된 제품 객체를 반환하는 메서드
  public build(): Product {
    return this.product;
  }
}