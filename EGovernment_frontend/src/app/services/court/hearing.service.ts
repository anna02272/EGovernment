import { HttpClient, HttpEvent } from '@angular/common/http';

import { Hearing } from '../../models/court/hearing';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Subject } from 'src/app/models/court/subject';
import { SubjectService } from './subject.service';

@Injectable({
  providedIn: 'root'
})
export class HearingService {
  private apiUrl = 'http://localhost:8083/api/hearing';

  constructor(private http: HttpClient, private subjectService :SubjectService) {}

  createHearing(hearing: Hearing, options?: any): Observable<HttpEvent<Hearing>> {
    return this.http.post<Hearing>(`${this.apiUrl}/create`, hearing, options);
  }
  getAllHearingsByJudge(): Observable<Hearing[]> {
    return this.http.get<Hearing[]>(`${this.apiUrl}/getByIdJudge`);
  }
  getSubjectById(subjectId: string): Observable<Subject> {
    return this.subjectService.getSubjectById(subjectId);
  }
}
