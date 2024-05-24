import { DegreeOfAccident } from "./degreeOfAccident";

export class ReportCarAccidentDegree {
  id: string;
  title: string;
  date: string;
  total_number: number;
  degree: DegreeOfAccident;
  year: number;
 
 
  constructor(id: string, title: string, date: string, total_number: number, degree: DegreeOfAccident, year: number ) {
    this.id = id;
    this.title = title;
    this.date = date;
    this.total_number = total_number;
    this.degree = degree;
    this.year = year;
  }
}