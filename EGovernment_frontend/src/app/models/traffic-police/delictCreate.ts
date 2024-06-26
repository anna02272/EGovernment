import { DelictType } from "./delictType";
import { DelictStatus } from "./delictStatus";

export class DelictCreate {
  driver_identification_number: string;
  vehicle_licence_number: string;
  driver_email: string;
  location: string;
  description: string;
  delict_type: DelictType;
  delict_status: DelictStatus;
  price_of_fine: number;
  number_of_penalty_points: number;

  constructor(
    driver_identification_number: string,
    vehicle_licence_number: string,
    driver_email: string,
    location: string,
    description: string,
    delict_type: DelictType,
    delict_status: DelictStatus,
    price_of_fine: number,
    number_of_penalty_points: number
  ) {
    this.driver_identification_number = driver_identification_number;
    this.vehicle_licence_number = vehicle_licence_number;
    this.driver_email = driver_email;
    this.location = location;
    this.description = description;
    this.delict_type = delict_type;
    this.delict_status = delict_status;
    this.price_of_fine = price_of_fine;
    this.number_of_penalty_points = number_of_penalty_points;
  }
}
