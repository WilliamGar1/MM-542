import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { Observable, Subject} from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  private URL = "http://localhost:3000";


  private _refresh$ = new Subject<void>();


  constructor(private http: HttpClient) { }

  public get refresh$() {
    return this._refresh$;
  }

  public getProducts(): Observable<any> {
    return this.http.get(this.URL + '/products');
  }

  public productFilter(keyword: string): Observable<any> {
    return this.http.get(this.URL + `/prod/${keyword}`);
  }

  public updateProduct(id: number, body: any): Observable<any> {
    return this.http.put(this.URL + `/updateProduct/${id}`, body)
      .pipe(
        tap(() => {
          this._refresh$.next();
        })
      )
  }

  public newOrder( body: any): Observable<any>{
    return this.http.post(this.URL + '/newOrder', body);
  }

  public newOrderDetail( body: any ): Observable<any> {
    return this.http.post(this.URL + '/newOrderDetail', body);
  }

  public getCustomers(): Observable<any> {
    return this.http.get(this.URL + '/customers');
  }

  public getEmployees(): Observable<any> {
    return this.http.get(this.URL + '/employees')
  }

  public getShippers(): Observable<any> {
    return this.http.get(this.URL + '/shippers')
  }

}
