import { Injectable } from "@angular/core";
import { ApiService } from "../api.service";
import { ConfigService } from "../config.service";
import { Category } from "src/app/models/statisics/category";
import { ReportRegisteredVehicle } from "src/app/models/statisics/reportRegisteredVehicle";

@Injectable()
export class ReportRegisteredVehiclesService {
    
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(report: ReportRegisteredVehicle){
    const url = `${this.config.registeredVehiclesReport_url}/create/category/${report.category}/year/${report.year}`;
    return this.apiService.post(url, report);
   }
  
  getAll() {
    const url = `${this.config.registeredVehiclesReport_url}/all`;
    return this.apiService.get(url);
   }
  
   getById(id : number) {
    const url = `${this.config.registeredVehiclesReport_url}/get/${id}`;
    return this.apiService.get(url);
   }

   getByCategory(category : Category) {
    const url = `${this.config.registeredVehiclesReport_url}/get/category/${category}`;
    return this.apiService.get(url);
   }

   getByCategoryAndYear(category : Category, year: number) {
    const url = `${this.config.registeredVehiclesReport_url}/get/category/${category}/year/${year}`;
    return this.apiService.get(url);
   }
}