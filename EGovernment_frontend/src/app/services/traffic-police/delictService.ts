import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Delict } from 'src/app/models/traffic-police//delict';
import { DelictCreate } from 'src/app/models/traffic-police/delictCreate';
import { DelictType } from "src/app/models/traffic-police/delictType";
import { ConfigService } from '../config.service';
import { ApiService } from "../api.service";

@Injectable({
  providedIn: 'root'
})
export class DelictService {

  constructor(
    private http: HttpClient,
    private config: ConfigService,
    private apiService: ApiService,
  ) {
  }

  insertDelict(delictCreate: DelictCreate) {
    const url = `${this.config.traffic_police_url}/createDelict`;
    return this.apiService.post(url, delictCreate);
  }

  getAllDelicts(){
    const url = `${this.config.traffic_police_url}/all`;
    return this.apiService.get(url);
  }

  getDelictById(delictId: string){
    const url = `${this.config.traffic_police_url}/get/${delictId}`;
    return this.apiService.get(url);
  }

  getAllDelictsByDelictType(delictType: DelictType) {
    const url = `${this.config.traffic_police_url}/get/delictType/${delictType}`;
    return this.apiService.get(url);
  }

  checkDriverAlcoholDelicts(driverId: string) {
    const url = `${this.config.traffic_police_url}/getDriver/${driverId}`;
    return this.apiService.get(url);
  }

  getAllDelictsByPolicemanID() {
    const url = `${this.config.traffic_police_url}/getPolicemanDelicts`;
    return this.apiService.get(url);
  }

  getAllDelictsByDriver() {
    const url = `${this.config.traffic_police_url}/getDriverDelicts`;
    return this.apiService.get(url);
  }

  getAllDelictsByDelictTypeAndYear(delictType: DelictType, year: number) {
    const url = `${this.config.traffic_police_url}/get/delictType/${delictType}/year/${year}`;
    return this.apiService.get(url);
  }

  updateDelictStatus(delictId: string) {
    const url = `${this.config.traffic_police_url}/pay/${delictId}`;
    return this.apiService.patch(url, null);
  }

  uploadImages(folderName: string, formData: FormData) {
    return this.apiService.post(`${this.config.traffic_police_url}/upload/${folderName}`, formData);
  }

  getImagesUrls(folderName: string): Observable<any>  {
    return this.http.get<any>(`${this.config.traffic_police_url}/getImagesUrls/${folderName}`);
  }

  getImages(folderName: string, imageName: string): Observable<any>{
    return this.http.get(`${this.config.traffic_police_url}/getImages/${folderName}/${imageName}`, { responseType: 'blob' });
  }

  getPdfByDelictId(id: string): Observable<Blob> {
    return this.http.get(`${this.config.traffic_police_url}/getPdf/${id}`, { responseType: 'blob' });
  }
}
