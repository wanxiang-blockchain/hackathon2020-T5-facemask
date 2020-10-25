import { Injectable } from '@angular/core';
import { HttpClient,HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { map } from"rxjs/operators";


@Injectable({
  providedIn: 'root'
})
export class BlockService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };
  constructor(private http: HttpClient) { }

  public getUpstream(): Observable<any>
  {
    let url:string = "/assets/data/block/upstream.json";
    return this.http.get(url);
  }

  public getDownstream(): Observable<any>
  {
    let url:string = "/assets/data/block/downstream.json";
    return this.http.get(url);
  }
}
