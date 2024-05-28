import { Component, OnInit } from '@angular/core';
import { ReportRegisteredVehiclesService } from 'src/app/services/statistics/reportRegisteredVehicles.service';

@Component({
  selector: 'app-report-registered-vehicles',
  templateUrl: './report-registered-vehicles.component.html',
  styleUrls: ['./report-registered-vehicles.component.css']
})
export class ReportRegisteredVehiclesComponent implements OnInit {
  reportData: any[] = [];
  categories = [
    { name: 'A', total_number: 0 },
    { name: 'B', total_number: 0 },
    { name: 'B1', total_number: 0 },
    { name: 'A1', total_number: 0 },
    { name: 'C', total_number: 0 },
    { name: 'AM', total_number: 0 },
    { name: 'A2', total_number: 0 }
  ];
  maxRegisteredVehicles: number = 0;

  constructor(private reportService: ReportRegisteredVehiclesService) {}

  ngOnInit(): void {
    this.reportService.getAll().subscribe((data: any[]) => {
      this.groupDataByYear(data);
      this.calculateMaxRegisteredVehicles();
      this.calculateMaxRegisteredVehiclesByYear();
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
            existingItem.categories.push({ category: item.category, total_number: item.total_number });
        } else {
            groupedData.push({
                year: year,
                title: item.title,
                date: item.date,
                categories: [{ category: item.category, total_number: item.total_number }]
            });
        }
    });

    groupedData.forEach((group) => {
        group.date = latestDataByYear[group.year].date.toISOString();
        group.title = latestDataByYear[group.year].title;
    });

    groupedData.sort((a, b) => b.year - a.year);

    this.reportData = groupedData;
    this.aggregateCategoryData();
  }

  getTotalNumberByCategory(item: any, category: string): number {
    const categoryData = item.categories.find((cat: { category: string }) => cat.category === category);
    return categoryData ? categoryData.total_number : "-";
  }
  
  getTotalRegisteredVehicles(): number {
    return this.reportData.reduce((total, item) => total + item.categories.reduce((subTotal: any, cat: { total_number: any; }) => subTotal + cat.total_number, 0), 0);
  }

  getCategoryWithMostRegisteredVehicles(): string {
    const categories = this.reportData.flatMap(item => item.categories);
    const maxCategory = categories.reduce((max, cat) => (max.total_number > cat.total_number ? max : cat), { total_number: 0 });
    return maxCategory.category;
  }

  getYearWithMostRegisteredVehicles(): number {
    const maxYear = this.reportData.reduce((max, item) => (max.categories?.reduce((subMax: { total_number: number; }, 
      cat: { total_number: number; }) => (subMax.total_number > cat.total_number ? subMax : cat), 
      { total_number: 0 }).total_number > item.categories.reduce((subMax: { total_number: number; }, 
        cat: { total_number: number; }) => (subMax.total_number > cat.total_number ? subMax : cat), 
        { total_number: 0 }).total_number ? max : item), { year: 0 });
    return maxYear.year;
  }

  aggregateCategoryData(): void {
    const categoryMap: { [name: string]: number } = {};
    
    this.reportData.forEach(item => {
      item.categories.forEach((cat: { category: string, total_number: number }) => {
        if (categoryMap[cat.category]) {
          categoryMap[cat.category] += cat.total_number;
        } else {
          categoryMap[cat.category] = cat.total_number;
        }
      });
    });

    this.categories = this.categories.map(cat => ({
      ...cat,
      total_number: categoryMap[cat.name] || 0
    }));
  }

  calculateMaxRegisteredVehicles(): void {
    this.maxRegisteredVehicles = Math.max(...this.categories.map(cat => cat.total_number));
  }

  calculateMaxRegisteredVehiclesByYear(): void {
    this.maxRegisteredVehicles = Math.max(...this.reportData.map(item => this.getTotalNumberForYear(item)));
  }

  getBarHeight(category: { name: string, total_number: number }): number {
    return this.maxRegisteredVehicles ? (category.total_number / this.maxRegisteredVehicles) * 100 : 0;
  }

  getTotalNumberForYear(item: any): number {
    return item.categories.reduce((total: number, cat: { total_number: number }) => total + cat.total_number, 0);
  }

  getBarHeightForYear(item: any): number {
    const totalNumber = this.getTotalNumberForYear(item);
    return this.maxRegisteredVehicles ? (totalNumber / this.maxRegisteredVehicles) * 100 : 0;
  }
}
