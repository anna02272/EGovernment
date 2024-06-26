import {Injectable} from "@angular/core";
import {ApiService} from "../api.service";
import {ConfigService} from "../config.service";
import {ReportDelict} from "src/app/models/statisics/reportDelict";
import {DelictType} from "src/app/models/statisics/delictType";

@Injectable()
export class ReportDelicTypeService {

  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(report: ReportDelict){
    const url = `${this.config.delictReport_url}/create/delictType/${report.type}/year/${report.year}`;
    return this.apiService.post(url, report);
   }

  getAll() {
    const url = `${this.config.delictReport_url}/all`;
    return this.apiService.get(url);
   }

   getById(id : number) {
    const url = `${this.config.delictReport_url}/get/${id}`;
    return this.apiService.get(url);
   }

   getByType(delictType : DelictType) {
    const url = `${this.config.delictReport_url}/get/delictType/${delictType}`;
    return this.apiService.get(url);
   }

   getByTypeAndYear(delictType : DelictType, year: number) {
    const url = `${this.config.delictReport_url}/get/delictType/${delictType}/year/${year}`;
    return this.apiService.get(url);
   }
}
