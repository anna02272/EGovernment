import { CarAccidentType } from "./carAccidentType";

export class ReportCarAccidentType {
  id: string;
  title: string;
  date: string;
  total_number: number;
  type: CarAccidentType;
  year: number;
 
 
  constructor(id: string, title: string, date: string, total_number: number, type: CarAccidentType, year: number ) {
    this.id = id;
    this.title = title;
    this.date = date;
    this.total_number = total_number;
    this.type = type;
    this.year = year;
  }
}