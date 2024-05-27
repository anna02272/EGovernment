import { Injectable } from "@angular/core";
import { ApiService } from "../api.service";
import { ConfigService } from "../config.service";
import { Request } from "src/app/models/statisics/request";

@Injectable()
export class RequestService {
    
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(request: Request){
    const url = `${this.config.request_url}/create`;
    return this.apiService.post(url, request);
   }
  
  getAll() {
    return this.apiService.get(this.config.request_url + "/all");
   }

   getById(id : string) {
    return this.apiService.get(this.config.request_url + "/get/" + id );
   }

   delete(id : string) {
    return this.apiService.delete(this.config.request_url + "/delete/" + id );
   }
}