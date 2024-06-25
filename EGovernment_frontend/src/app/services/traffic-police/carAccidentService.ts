import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CarAccident } from 'src/app/models/traffic-police/carAccident';
import { CarAccidentType } from 'src/app/models/traffic-police/carAccidentType';
import { DegreeOfAccident } from 'src/app/models/traffic-police/degreeOfAccident';
import { ConfigService } from '../config.service';
import { ApiService } from "../api.service";

@Injectable({
  providedIn: 'root'
})
export class CarAccidentService {

  constructor(
    private http: HttpClient,
    private config: ConfigService,
    private apiService: ApiService,
  ) {
  }

  insertCarAccident(carAccident: CarAccident, policemanID: string) {
    const url = `${this.config.car_accident_url}/createCarAccident?policemanID=${policemanID}`;
    return this.apiService.post(url, carAccident);
  }

  getAllCarAccidents(){
    const url = `${this.config.car_accident_url}/all`;
    return this.apiService.get(url);
  }

  getCarAccidentById(carAccidentId: string) {
    const url = `${this.config.car_accident_url}/get/${carAccidentId}`;
    return this.apiService.get(url);
  }

  getAllCarAccidentsByType(carAccidentType: CarAccidentType) {
    const url = `${this.config.car_accident_url}/get/type/${carAccidentType}`;
    return this.apiService.get(url);
  }

  getAllCarAccidentsByTypeAndYear(carAccidentType: CarAccidentType, year: number) {
    const url = `${this.config.car_accident_url}/get/type/${carAccidentType}/year/${year}`;
    return this.apiService.get(url);
  }

  getAllCarAccidentsByDegree(degreeOfAccident: DegreeOfAccident) {
    const url = `${this.config.car_accident_url}/get/degree/${degreeOfAccident}`;
    return this.apiService.get(url);
  }

  getAllCarAccidentsByDegreeAndYear(degreeOfAccident: DegreeOfAccident, year: number) {
    const url = `${this.config.car_accident_url}/get/degree/${degreeOfAccident}/year/${year}`;
    return this.apiService.get(url);
  }

  getAllCarAccidentsByPolicemanID(){
    const url = `${this.config.car_accident_url}/getPolicemanCarAccident`;
    return this.apiService.get(url);
  }

  getAllCarAccidentsByDriver() {
    const url = `${this.config.car_accident_url}/getDriverCarAccident`;
    return this.apiService.get(url);
  }
}