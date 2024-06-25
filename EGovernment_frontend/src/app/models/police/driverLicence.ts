import { Category } from "./category";

export interface DriverLicence {
    vehicle_driver: string;
    licence_number: string;
    location_licenced: Location;
    categories: Category[];
}