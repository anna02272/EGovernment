import { DelictType } from "./delictType";
import { DelictStatus } from "./delictStatus";

export class Delict {
  id: string;
  policeman_id: string;
  driver_identification_number: string;
  vehicle_licence_number: string;
  driver_email: string;
  driver_jmbg: string;
  date: string;
  location: string;
  description: string;
  delict_type: DelictType;
  delict_status: DelictStatus;
  price_of_fine: number;
  number_of_penalty_points: number;
 
 
  constructor(id: string,
    policeman_id: string, 
    driver_identification_number: string, 
    vehicle_licence_number: string, 
    driver_email: string, 
    driver_jmbg: string, 
    date: string, 
    location: string, 
    description: string, 
    delict_type: DelictType, 
    delict_status: DelictStatus, 
    price_of_fine: number, 
    number_of_penalty_points: number ) 
    {
    this.id = id;
    this.policeman_id = policeman_id;
    this.driver_identification_number = driver_identification_number;
    this.vehicle_licence_number = vehicle_licence_number;
    this.driver_email = driver_email;
    this.driver_jmbg = driver_jmbg;
    this.date = date;
    this.location = location;
    this.description = description;
    this.delict_type = delict_type;
    this.delict_status = delict_status;
    this.price_of_fine = price_of_fine;
    this.number_of_penalty_points = number_of_penalty_points;
  }
}