import { Injectable } from "@angular/core";
import { ApiService } from "../api.service";
import { ConfigService } from "../config.service";
import { ReportCarAccidentType } from "src/app/models/statisics/reportCarAccidentType";
import { CarAccidentType } from "src/app/models/statisics/carAccidentType";

@Injectable()
export class ReportCarAccidentTypeService {
    
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(report: ReportCarAccidentType){
    const url = `${this.config.carAccidentTypeReport_url}/create/carAccidentType/${report.type}/year/${report.year}`;
    return this.apiService.post(url, report);
   }
  
  getAll() {
    const url = `${this.config.carAccidentTypeReport_url}/all`;
    return this.apiService.get(url);
   }
  
   getById(id : number) {
    const url = `${this.config.carAccidentTypeReport_url}/get/${id}`;
    return this.apiService.get(url);
   }

   getByType(carAccidentType : CarAccidentType) {
    const url = `${this.config.carAccidentTypeReport_url}/get/carAccidentType/${carAccidentType}`;
    return this.apiService.get(url);
   }

   getByTypeAndYear(carAccidentType : CarAccidentType, year: number) {
    const url = `${this.config.carAccidentTypeReport_url}/get/carAccidentType/${carAccidentType}/year/${year}`;
    return this.apiService.get(url);
   }
}