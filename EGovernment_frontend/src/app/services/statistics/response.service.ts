import { Injectable } from "@angular/core";
import { ApiService } from "../api.service";
import { ConfigService } from "../config.service";
import { Response } from "src/app/models/statisics/response";

@Injectable()
export class ResponseService {
    
  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(response: Response){
    const url = `${this.config.response_url}/create`;
    return this.apiService.post(url, response);
   }
  
  getAll() {
    return this.apiService.get(this.config.response_url + "/all");
   }
  
   getById(id : string) {
    return this.apiService.get(this.config.response_url + "/get/" + id );
   }
}