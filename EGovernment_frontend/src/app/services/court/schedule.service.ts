import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Schedule } from '../../models/court/schedule';

@Injectable({
  providedIn: 'root'
})
export class ScheduleService {
  private apiUrl = 'http://localhost:8083/api/schedule';

  constructor(private http: HttpClient) { }

  createSchedule(schedule: Schedule): Observable<Schedule> {
    return this.http.post<Schedule>(`http://localhost:8083/api/schedule/create`, schedule);
  }
}
