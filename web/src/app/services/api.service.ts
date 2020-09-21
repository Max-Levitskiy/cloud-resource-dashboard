import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {environment} from '../../environments/environment';
import {timeout} from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class ApiService {

  constructor(private http: HttpClient) { }

  ping(): Promise<boolean> {
    return this.http.get(`${environment.api.host}:${environment.api.port}/ping`)
      .pipe(timeout(2000))
      .toPromise()
      .then(() => {
        return true;
      }).catch(() => {
        return false;
      });
  }
}
