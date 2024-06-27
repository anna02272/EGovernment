import {Injectable} from "@angular/core";
import {ApiService} from "../api.service";
import {ConfigService} from "../config.service";
import {Vehicle} from "src/app/models/police/vehicle";
import { VehicleDriver } from "src/app/models/police/vehicleDriver";
import { DriverLicence } from "src/app/models/police/driverLicence";
import { Observable } from "rxjs/internal/Observable";
import { HttpClient } from "@angular/common/http";
import { Category } from "src/app/models/police/category";

@Injectable()
export class VehicleService {

  constructor(
    private apiService: ApiService,
    private http: HttpClient,
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

   getAllLicences() {
    const url = this.config.getAllDriverLicences_url;
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

   getDriverLicenceById(id: string) {
    const url = this.config.getDriverLicenceById_url + id;
    return this.apiService.get(url,id);
   }

  getDriverLicenceByDriver(id: string) {
    const url = this.config.getDriverLicenceByDriver_url + id;
    return this.apiService.get(url,id);
  }

   getByCategoryAndYear(category: string, year: number) {
    const url = this.config.getVehicleByCategoryAndYear_url(category, year);
    return this.apiService.get(url);
  }

  createVehicleDriver(vehicleDriver: VehicleDriver){
    const url = this.config.createVehicleDriver;
    return this.apiService.post(url, vehicleDriver);
   }

   createDriverLicence(driverLicence: DriverLicence){
    const url = this.config.createDriverLicence_url;
    return this.apiService.post(url, driverLicence);
   }

   getNumberOfRegVehiclesCategory(category: string) {
    const url = this.config.getCountRegisteredVehiclesCategory + category;
    return this.apiService.get(url);
   }

   getRegisteredVehiclesPdf(): Observable<Blob> {
    return this.http.get('http://localhost:8080/api/vehicle/registeredVehicles/pdf', { responseType: 'blob' });
  }

  getRegisteredVehiclesCategoryPdf(searchCategory: string): Observable<Blob> {
    return this.http.get(`http://localhost:8080/api/vehicle/registeredVehicles/category/${searchCategory}/pdf`, { responseType: 'blob' });
  }
}
