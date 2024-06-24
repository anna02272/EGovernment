import {Injectable} from "@angular/core";
import {ApiService} from "../api.service";
import {ConfigService} from "../config.service";
import {Vehicle} from "src/app/models/police/vehicle";

@Injectable()
export class VehicleService {

  constructor(
    private apiService: ApiService,
    private config: ConfigService
  ) {
  }
  create(vehicle: Vehicle){
    const url = this.config.createVehicle_url;
    return this.apiService.post(url, vehicle);
   }

  getAll() {
    const url = this.config.allVehicles_url;
    return this.apiService.get(url);
   }

//    getById(id : string) {
//     return this.apiService.get(this.config.request_url + "/get/" + id );
//    }

//    delete(id : string) {
//     return this.apiService.delete(this.config.request_url + "/delete/" + id );
//    }
}
