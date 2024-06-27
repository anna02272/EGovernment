import { Gender } from "./gender";

export interface VehicleDriver {
    identification_number: string;
    name: string;
    last_name: string;
    date_of_birth?: string;
    gender: Gender;
    number_of_penalty_points:number;
    has_delict: boolean;
  }
