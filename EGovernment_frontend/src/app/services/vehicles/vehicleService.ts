import {Injectable} from "@angular/core";
import {ApiService} from "../api.service";
import {ConfigService} from "../config.service";
import {Vehicle} from "src/app/models/police/vehicle";
import { VehicleDriver } from "src/app/models/police/vehicleDriver";

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

   getAllVehicleDrivers() {
    const url = this.config.allVehicleDrivers_url;
    return this.apiService.get(url);
   }
   getAllRegisteredVehicles() {
    const url = this.config.getRegisteredVehicles_url;
    return this.apiService.get(url);
   }

   getById(id : string) {
    const url = this.config.getVehicleById_url + id;
    return this.apiService.get(url, id);
   }

   getDriverById(id : string) {
    const url = this.config.getVehicleDriverById_url + id;
    return this.apiService.get(url, id);
   }


   getByCategoryAndYear(category: string, year: number) {
    const url = this.config.getVehicleByCategoryAndYear_url(category, year);
    return this.apiService.get(url);
  }

  createVehicleDriver(vehicleDriver: VehicleDriver){
    const url = this.config.createVehicleDriver;
    return this.apiService.post(url, vehicleDriver);
   }


//    delete(id : string) {
//     return this.apiService.delete(this.config.request_url + "/delete/" + id );
//    }
}
