import {Court} from '../../models/court/court';
import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CourtService {

  private apiUrl = `http://localhost:8083/api/court/courts`;

  constructor(private http: HttpClient) { }

  getAllCourts(): Observable<Court[]> {
    return this.http.get<Court[]>(`${this.apiUrl}/all`);
  }
}
