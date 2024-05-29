import { Component, OnInit } from '@angular/core';
import { ReportCarAccidentDegreeService } from 'src/app/services/statistics/reportCarAccidentDegree.service';


@Component({
  selector: 'app-report-car-accident-degree',
  templateUrl: './report-car-accident-degree.component.html',
  styleUrls: ['./report-car-accident-degree.component.css']
})
export class ReportCarAccidentDegreeComponent implements OnInit {
  reportData: any[] = [];
  degrees = [
    { name: 'NoHarm', total_number: 0 },
    { name: 'WithMaterialDamage', total_number: 0 },
    { name: 'WithInjuredPersons', total_number: 0 },
    { name: 'WithDeadBodies', total_number: 0 },
  ];
  degreeTranslation: { [key: string]: string } = {
    'NoHarm': 'Bez štete',
    'WithMaterialDamage': 'Sa materijalnom štetom',
    'WithInjuredPersons': 'Sa povređenim licima',
    'WithDeadBodies': 'Sa poginulim licima'
  };
  maxCarAccidents: number = 0;

  constructor(private reportService: ReportCarAccidentDegreeService) {}

  ngOnInit(): void {
    this.reportService.getAll().subscribe((data: any[]) => {
      this.groupDataByYear(data);
      this.calculateMaxCarAccidents();
      this.calculateMaxCarAccidentsByYear();
    });
  }

  groupDataByYear(data: any[]): void {
    const groupedData: any[] = [];
    const latestDataByYear: { [year: number]: { date: Date, title: string } } = {};

    data.forEach((item) => {
        const year = item.year;
        const currentDate = new Date(item.date);

        if (!latestDataByYear[year] || currentDate > latestDataByYear[year].date) {
            latestDataByYear[year] = { date: currentDate, title: item.title };
        }

        const existingItem = groupedData.find((group) => group.year === year);
        if (existingItem) {
            existingItem.degrees.push({ degree: item.degree, total_number: item.total_number });
        } else {
            groupedData.push({
                year: year,
                title: item.title,
                date: item.date,
                degrees: [{ degree: item.degree, total_number: item.total_number }]
            });
        }
    });

    groupedData.forEach((group) => {
        group.date = latestDataByYear[group.year].date.toISOString();
        group.title = latestDataByYear[group.year].title;
    });

    groupedData.sort((a, b) => b.year - a.year);

    this.reportData = groupedData;
    this.aggregateDegreeData();
  }

  getTotalNumberByDegree(item: any, degree: string): number {
    const degreeData = item.degrees.find((deg: { degree: string }) => deg.degree === degree);
    return degreeData ? degreeData.total_number : "-";
  }
  
  getTotalCarAccidents(): number {
    return this.reportData.reduce((total, item) => total + item.degrees.reduce((subTotal: any, deg: { total_number: any; }) => subTotal + deg.total_number, 0), 0);
  }

  getDegreeWithMostCarAccidents(): string {
    const degrees = this.reportData.flatMap(item => item.degrees);
    const maxDegree = degrees.reduce((max, deg) => (max.total_number > deg.total_number ? max : deg), { total_number: 0 });
    return maxDegree.degree;
  }

  getYearWithMostCarAccidents(): number {
    const maxYear = this.reportData.reduce((max, item) => (max.degrees?.reduce((subMax: { total_number: number; }, 
      deg: { total_number: number; }) => (subMax.total_number > deg.total_number ? subMax : deg), 
      { total_number: 0 }).total_number > item.degrees.reduce((subMax: { total_number: number; }, 
        deg: { total_number: number; }) => (subMax.total_number > deg.total_number ? subMax : deg), 
        { total_number: 0 }).total_number ? max : item), { year: 0 });
    return maxYear.year;
  }

  aggregateDegreeData(): void {
    const degreeMap: { [name: string]: number } = {};
    
    this.reportData.forEach(item => {
      item.degrees.forEach((deg: { degree: string, total_number: number }) => {
        if (degreeMap[deg.degree]) {
          degreeMap[deg.degree] += deg.total_number;
        } else {
          degreeMap[deg.degree] = deg.total_number;
        }
      });
    });

    this.degrees = this.degrees.map(cat => ({
      ...cat,
      total_number: degreeMap[cat.name] || 0
    }));
  }

  calculateMaxCarAccidents(): void {
    this.maxCarAccidents = Math.max(...this.degrees.map(deg => deg.total_number));
  }

  calculateMaxCarAccidentsByYear(): void {
    this.maxCarAccidents = Math.max(...this.reportData.map(item => this.getTotalNumberForYear(item)));
  }

  getBarHeight(degree: { name: string, total_number: number }): number {
    return this.maxCarAccidents ? (degree.total_number / this.maxCarAccidents) * 100 : 0;
  }

  getTotalNumberForYear(item: any): number {
    return item.degrees.reduce((total: number, deg: { total_number: number }) => total + deg.total_number, 0);
  }

  getBarHeightForYear(item: any): number {
    const totalNumber = this.getTotalNumberForYear(item);
    return this.maxCarAccidents ? (totalNumber / this.maxCarAccidents) * 100 : 0;
  }
}
