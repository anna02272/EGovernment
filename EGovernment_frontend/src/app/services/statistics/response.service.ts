import { Injectable } from "@angular/core";
import { ApiService } from "../api.service";
import { ConfigService } from "../config.service";

@Injectable()
export class ResponseService {
    
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(formData: FormData){
    const url = `${this.config.response_url}/create`;
    return this.apiService.post(url, formData);
   }
  
  getAll() {
    return this.apiService.get(this.config.response_url + "/all");
   }
  
   getById(id : string) {
    return this.apiService.get(this.config.response_url + "/get/" + id );
   }
}