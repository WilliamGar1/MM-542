import { Pipe, PipeTransform } from '@angular/core';
import { Observable } from 'rxjs';

@Pipe({
  name: 'page'
})
export class PagePipe implements PipeTransform {

  transform(value: any, page: number = 0): any {

    if(value === null){
      return []
    }else{

      return value.slice(page, page + 5)
    }


  }

}
