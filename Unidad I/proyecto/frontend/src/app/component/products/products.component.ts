import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { HttpService } from 'src/app/http.service';

@Component({
  selector: 'app-products',
  templateUrl: './products.component.html',
  styleUrls: ['./products.component.css']
})
export class ProductsComponent implements OnInit {

  public products: any = [];

  constructor(
    private httpService: HttpService,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.getProducts();
  }

  public getProducts() {
    this.httpService.getProducts()
      .subscribe(data => {
        this.products = data;
        console.log(this.products);
      });
  }

  public productFilter() {
    this.httpService.productFilter()
      .subscribe(data => {
        this.products = data;
        console.log(this.products)
      })
  }

}
