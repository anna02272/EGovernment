import { HttpClient, HttpEvent } from '@angular/common/http';

import { Hearing } from '../../models/court/hearing';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HearingService {
  private apiUrl = 'http://localhost:8083/api/hearing';

  constructor(private http: HttpClient) {}

  createHearing(hearing: Hearing, options?: any): Observable<HttpEvent<Hearing>> {
    return this.http.post<Hearing>(`${this.apiUrl}/create`, hearing, options);
  }
}
