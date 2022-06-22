import { Component, OnDestroy, OnInit } from '@angular/core';
import { Subscription } from 'rxjs';
import { HttpService } from 'src/app/http.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit, OnDestroy {

  public products: any = [];
  public keyword: string = "";
  public page: number = 0;

  private subscription: Subscription

  public product = {
    ProductID: 0,
    UnitPrice: 0,
    UnitsInStock: 0
  }

  constructor(
    private httpService: HttpService,
  ) { }

  ngOnDestroy(): void {
    this.subscription.unsubscribe();
    console.log('Obvservable cerrado');
  }

  ngOnInit(): void {
    this.getProducts();

    this.subscription = this.httpService.refresh$
      .subscribe(() => {
        this.getProducts();
      })
  }

  public getProducts() {
    this.httpService.getProducts()
      .subscribe(data => {
        this.products = data;
        console.log(this.products);
      });
  }

  public productFilter(value: string) {
    this.page = 0;
    this.httpService.productFilter(value)
      .subscribe(data => {
        this.products = data;
      });
  }

  public updateProduct(price: string, stock: string) {

    let pro = {
      ProductID: this.product.ProductID,
      UnitPrice: parseFloat(price),
      UnitsInStock: parseInt(stock)
    }
    
    this.httpService.updateProduct(this.product.ProductID, pro)
      .subscribe(data => {
        Swal.fire(
          'Éxito',
          'Producto Actualizado',
          'success',
        );
      });
  }

  public productDetails(ProductID: number, UnitPrice: number, UnitsInStock: number){
    this.product = {ProductID, UnitPrice, UnitsInStock}
  }

  public nextPage() {
    this.page += 5;
  }

  public prevPage() {
    if(this.page > 0)
      this.page -= 5;
  }
}
