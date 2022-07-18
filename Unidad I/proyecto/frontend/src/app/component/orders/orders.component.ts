import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { HttpService } from 'src/app/http.service';
import Swal from 'sweetalert2';

@Component({
  selector: 'app-orders',
  templateUrl: './orders.component.html',
  styleUrls: ['./orders.component.css']
})
export class OrdersComponent implements OnInit {

  public customers: any = [];
  public employees: any = [];
  public shippers: any = [];

  public date = new Date();

  orderForm = new FormGroup({
    CustomerID: new FormControl('', Validators.required),
    EmployeeID: new FormControl('', Validators.required),
    ShipVia: new FormControl('', Validators.required),
    Freight: new FormControl('', Validators.required),
    ShipName: new FormControl('', Validators.required),
    ShipAddress: new FormControl('', Validators.required),
    ShipCity: new FormControl('', Validators.required),
    ShipRegion: new FormControl('', Validators.required),
    ShipPostalCode: new FormControl('', Validators.required),
    ShipCountry: new FormControl('', Validators.required)
  });

  constructor(
    private httpService: HttpService,
  ) { }

  ngOnInit(): void {
    this.getCustomers();
    this.getEmployees();
    this.getShippers();
    console.log(this.getDate(this.date))
  }

  public newOrder(): void {
    let o = this.orderForm.value;

    var order = {
      OrderID:        o.OrderID,
      CustomerID:     o.CustomerID,
      EmployeeID:     o.EmployeeID,
      OrderDate:      this.getDate(this.date),
      RequiredDate:   this.getDate(this.sumDays(this.date)),
      ShipVia:        o.ShipVia,
      Freight:        o.Freight,
      ShipName:       o.ShipName,
      ShipAddress:    o.ShipAddress,
      ShipCity:       o.ShipCity,
      ShipRegion:     o.ShipRegion,
      ShipPostalCode: o.ShipPostalCode,
      ShipCountry:    o.ShipCountry
    }

    if (this.orderForm.valid){
      this.httpService.newOrder(order)
        .subscribe(res => {

          if(res.Success){
            Swal.fire(
              'Éxito',
              `${res.Msj}`,
              'success',
            );
          }

        });
    }else{
      Swal.fire(
        'Error',
        'Complete todos los campos',
        'error',
      );
    }

  }

  public getCustomers() {
    this.httpService.getCustomers()
      .subscribe(data => {
        this.customers = data;
        //console.log(this.customers);
      })
  }

  public getEmployees() {
    this.httpService.getEmployees()
      .subscribe(data => {
        this.employees = data;
        //console.log(this.employees);
      })
  }

  public getShippers() {
    this.httpService.getShippers()
      .subscribe(data => {
        this.shippers = data;
        //console.log(this.shippers);
      })
  }

  public sumDays(date: Date){
    date.setDate(date.getDate() + 2);
    return date;
  }

  public getDate(date: Date){
    return date.toISOString().split('T')[0];
  }
}
