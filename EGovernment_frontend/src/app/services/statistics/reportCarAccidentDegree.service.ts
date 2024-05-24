import { Injectable } from "@angular/core";
import { ApiService } from "../api.service";
import { ConfigService } from "../config.service";
import { Category } from "src/app/models/statisics/category";
import { ReportCarAccidentDegree } from "src/app/models/statisics/reportCarAccidentDegree";
import { DegreeOfAccident } from "src/app/models/statisics/degreeOfAccident";

@Injectable()
export class ReportCarAccidentDegreeService {
    
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(report: ReportCarAccidentDegree, degree: DegreeOfAccident, year: number){
    const url = `${this.config.carAccidentDegreeReport_url}/create/degree/${degree}/year/${year}`;
    return this.apiService.post(url, report);
   }
  
  getAll() {
    const url = `${this.config.carAccidentDegreeReport_url}/all`;
    return this.apiService.get(url);
   }
  
   getById(id : number) {
    const url = `${this.config.carAccidentDegreeReport_url}/get/${id}`;
    return this.apiService.get(url);
   }

   getByDegree(degree : DegreeOfAccident) {
    const url = `${this.config.carAccidentDegreeReport_url}/get/degree/${degree}`;
    return this.apiService.get(url);
   }

   getByDegreeAndYear(degree : DegreeOfAccident, year: number) {
    const url = `${this.config.carAccidentDegreeReport_url}/get/degree/${degree}/year/${year}`;
    return this.apiService.get(url);
   }
}