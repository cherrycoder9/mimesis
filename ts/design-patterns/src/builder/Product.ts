// 의류 쇼핑몰의 제품 객체를 정의하는 클래스
// 제품 클래스, 의류 제품의 속성을 표현함
export class Product {
  private name: string;
  private price: number;
  private category: string;
  private description: string;

  // private 생성자, 외부에서 직접 new Product()로 객체를 생성하지
  // 못하게 하고, 빌더를 통해서만 생성 가능하게 함
  private constructor() {
    this.name = '';
    this.price = 0;
    this.category = '';
    this.description = '';
  }

  // 정적 팩토리 메서드 추가
  public static create(): Product {
    return new Product();
  }

  // 속성 값을 반환하는 getter 메서드들
  public getName(): string {
    return this.name;
  }

  public getPrice(): number {
    return this.price;
  }

  public getCategory(): string {
    return this.category;
  }

  public getDescription(): string {
    return this.description;
  }
}