import { CarAccidentType } from './carAccidentType';
import { DegreeOfAccident } from './degreeOfAccident';

export class CarAccidentCreate {
  policeman_id: string;
  driver_identification_number_first: string;
  driver_identification_number_second: string;
  vehicle_licence_number_first: string;
  vehicle_licence_number_second: string;
  driver_email: string;
  date: string;
  location: string;
  description: string;
  car_accident_type: CarAccidentType;
  degree_of_accident: DegreeOfAccident;
  number_of_penalty_points: number;

  constructor(
    policeman_id: string,
    driver_identification_number_first: string,
    driver_identification_number_second: string,
    vehicle_licence_number_first: string,
    vehicle_licence_number_second: string,
    driver_email: string,
    date: string,
    location: string,
    description: string,
    car_accident_type: CarAccidentType,
    degree_of_accident: DegreeOfAccident,
    number_of_penalty_points: number
  ) {
    this.policeman_id = policeman_id;
    this.driver_identification_number_first = driver_identification_number_first;
    this.driver_identification_number_second = driver_identification_number_second;
    this.vehicle_licence_number_first = vehicle_licence_number_first;
    this.vehicle_licence_number_second = vehicle_licence_number_second;
    this.driver_email = driver_email;
    this.date = date;
    this.location = location;
    this.description = description;
    this.car_accident_type = car_accident_type;
    this.degree_of_accident = degree_of_accident;
    this.number_of_penalty_points = number_of_penalty_points;
  }
}
