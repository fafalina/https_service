import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { DataModel } from './data.model';

@Injectable({
  providedIn: 'root'
})

export class DataService {
  private apiUrl = 'https://localhost:8080';

  constructor(private http: HttpClient) { }

  getData(): Observable<{ data: DataModel }> {
    return this.http.get<{ data: DataModel }>(this.apiUrl);
  }
}
