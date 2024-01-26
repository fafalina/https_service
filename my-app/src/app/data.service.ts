import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, interval } from 'rxjs';
import { switchMap } from 'rxjs/operators';
import { TimestampData } from './data.model';

@Injectable({
  providedIn: 'root'
})

export class DataService {
  private apiUrl = 'https://localhost:8080';

  constructor(private http: HttpClient) { }

  getTimestampData(): Observable<TimestampData> {
    return this.http.get<TimestampData>(this.apiUrl);
  }

  getUpdatedTimestampData(): Observable<TimestampData> {
    return interval(1000)
      .pipe(
        switchMap(() => this.http.get<TimestampData>(this.apiUrl))
      );
  }
}
