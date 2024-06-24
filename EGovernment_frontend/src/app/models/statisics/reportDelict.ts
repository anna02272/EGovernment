import {DelictType} from "./delictType";

export class ReportDelict {
  id: string;
  title: string;
  date: string;
  total_number: number;
  type: DelictType;
  year: number;


  constructor(id: string, title: string, date: string, total_number: number, type: DelictType, year: number ) {
    this.id = id;
    this.title = title;
    this.date = date;
    this.total_number = total_number;
    this.type = type;
    this.year = year;
  }
}
