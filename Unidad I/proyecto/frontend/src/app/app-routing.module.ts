import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { OrdersComponent } from './component/orders/orders.component';
import { ProductsComponent } from './component/products/products.component';

const routes: Routes = [
  {
    path: 'productos', component: ProductsComponent
  },
  {
    path: 'ordenes', component: OrdersComponent
  },
  {
    path: '**', redirectTo: 'productos', pathMatch: 'full'
  }
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
