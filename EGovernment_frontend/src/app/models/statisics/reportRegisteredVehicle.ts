import { Category } from "./category";

export class ReportRegisteredVehicle {
  id: string;
  title: string;
  date: string;
  total_number: number;
  category: Category;
  year: number;
 
 
  constructor(id: string, title: string, date: string, total_number: number, category: Category, year: number ) {
    this.id = id;
    this.title = title;
    this.date = date;
    this.total_number = total_number;
    this.category = category;
    this.year = year;
  }
}