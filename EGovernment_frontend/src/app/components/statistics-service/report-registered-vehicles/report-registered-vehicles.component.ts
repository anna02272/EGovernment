import { Component, OnInit } from '@angular/core';
import { ReportRegisteredVehiclesService } from 'src/app/services/statistics/reportRegisteredVehicles.service';

@Component({
  selector: 'app-report-registered-vehicles',
  templateUrl: './report-registered-vehicles.component.html',
  styleUrls: ['./report-registered-vehicles.component.css']
})
export class ReportRegisteredVehiclesComponent implements OnInit {
  reportData: any[] = [];

  constructor(private reportService: ReportRegisteredVehiclesService) {}

  ngOnInit(): void {
    this.reportService.getAll().subscribe((data: any[]) => {
      this.groupDataByYear(data);
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
  }



  getTotalNumberByCategory(item: any, category: string): number {
    const categoryData = item.categories.find((cat: { category: string }) => cat.category === category);
    return categoryData ? categoryData.total_number : "-";
  }
}
