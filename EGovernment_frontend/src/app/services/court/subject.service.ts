import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';
import {Subject} from '../../models/court/subject';

@Injectable({
  providedIn: 'root'
})
export class SubjectService {
  private apiUrl = 'http://localhost:8083/api/subject'; // Adjust URL as per your backend API

  constructor(private http: HttpClient) {}

  getAllSubjects(): Observable<Subject[]> {
    console.log(this.http.get<any[]>(`http://localhost:8083/api/subject/subjects`))
    return this.http.get<Subject[]>(`${this.apiUrl}/subjects`);
  }

  getSubjectById(id: string): Observable<Subject> {
    const url = `${this.apiUrl}/subjects/${id}`;
    return this.http.get<Subject>(url);
  }
  getViolationDetails(violationId: string): Observable<any> {
    return this.http.get<any>(`http://localhost:8084/api/delict/get/${violationId}`);
  }

  createSubject(subject: Subject): Observable<Subject> {
    return this.http.post<Subject>(`${this.apiUrl}/create`, subject);
  }

  updateSubjectStatus(id: string, status: string): Observable<Subject> {
    const url = `${this.apiUrl}/subjects/${id}/status`;
    return this.http.put<Subject>(url, { status });
  }

  updateSubjectJudgment(id: string, judgment: string): Observable<Subject> {
    const url = `${this.apiUrl}/subjects/${id}/judgment`;
    return this.http.put<Subject>(url, { judgment });
  }

  updateSubjectCompromis(id: string, compromis: string): Observable<Subject> {
    const url = `${this.apiUrl}/subjects/${id}/compromis`;
    return this.http.put<Subject>(url, { compromis });
  }

}
