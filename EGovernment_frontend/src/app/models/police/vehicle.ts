import {Category} from "./category";
import {VehicleModel} from "./vehicleModel";

export interface Vehicle {
    registration_plate: string;
    vehicle_model: VehicleModel;
    vehicle_owner: string;
    registration_date?: string;
    category: Category;
  }
