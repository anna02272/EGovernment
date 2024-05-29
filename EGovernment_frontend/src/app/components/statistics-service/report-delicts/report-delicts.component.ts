import { Component, OnInit } from '@angular/core';
import { ReportDelicTypeService } from 'src/app/services/statistics/reportDelicType.service';

@Component({
  selector: 'app-report-delicts',
  templateUrl: './report-delicts.component.html',
  styleUrls: ['./report-delicts.component.css']
})
export class ReportDelictsComponent implements OnInit {
  reportData: any[] = [];
  types = [ 
    { name: 'Speeding', total_number: 0 },
    { name: 'DrivingUnderTheInfluenceOfAlcohol', total_number: 0 },
    { name: 'DrivingUnderTheInfluence', total_number: 0 },
    { name: 'ImproperOvertaking', total_number: 0 },
    { name: 'ImproperParking', total_number: 0 },
    { name: 'FailureTooComplyWithTrafficLightsAndSigns', total_number: 0 },
    { name: 'ImproperUseOfSeatBeltsAndChildSeats', total_number: 0 },
    { name: 'UsingMobilePhoneWhileDriving', total_number: 0 },
    { name: 'ImproperUseOfMotorVehicle', total_number: 0 },
  ];
  typeTranslation: { [key: string]: string } = {
    'Speeding': 'Prekoračenje brzine',
    'DrivingUnderTheInfluenceOfAlcohol': 'Vožnja pod uticajem alkohola',
    'DrivingUnderTheInfluence': 'Vožnja pod uticajem droge',
    'ImproperOvertaking': 'Nepravilno preticanje',
    'ImproperParking': 'Nepravilno parkiranje',
    'FailureTooComplyWithTrafficLightsAndSigns': 'Nepridržavanje saobraćajnih znakova i svetlosnih signalizacija',
    'ImproperUseOfSeatBeltsAndChildSeats': 'Nepravilna upotreba sigurnosnih pojaseva i dečijih sedišta',
    'UsingMobilePhoneWhileDriving': 'Upotreba mobilnog telefona tokom vožnje',
    'ImproperUseOfMotorVehicle': 'Nepravilno rukovanje vozilom',
    
  };
  maxDelicts: number = 0;

  constructor(private reportService: ReportDelicTypeService) {}

  ngOnInit(): void {
    this.reportService.getAll().subscribe((data: any[]) => {
      this.groupDataByYear(data);
      this.calculateMaxDelicts();
      this.calculateMaxDelictsByYear();
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
            existingItem.types.push({ type: item.type, total_number: item.total_number });
        } else {
            groupedData.push({
                year: year,
                title: item.title,
                date: item.date,
                types: [{ type: item.type, total_number: item.total_number }]
            });
        }
    });

    groupedData.forEach((group) => {
        group.date = latestDataByYear[group.year].date.toISOString();
        group.title = latestDataByYear[group.year].title;
    });

    groupedData.sort((a, b) => b.year - a.year);

    this.reportData = groupedData;
    this.aggregateTypeData();
  }

  getTotalNumberByType(item: any, type: string): number {
    const typeData = item.types.find((typ: { type: string }) => typ.type === type);
    return typeData ? typeData.total_number : "-";
  }
  
  getTotalDelicts(): number {
    return this.reportData.reduce((total, item) => total + item.types.reduce((subTotal: any, typ: { total_number: any; }) => subTotal + typ.total_number, 0), 0);
  }

  getTypeWithMostDelicts(): string {
    const types = this.reportData.flatMap(item => item.types);
    const maxType = types.reduce((max, typ) => (max.total_number > typ.total_number ? max : typ), { total_number: 0 });
    return maxType.type;
  }

  getYearWithMostDelicts(): number {
    const maxYear = this.reportData.reduce((max, item) => (max.types?.reduce((subMax: { total_number: number; }, 
      typ: { total_number: number; }) => (subMax.total_number > typ.total_number ? subMax : typ), 
      { total_number: 0 }).total_number > item.types.reduce((subMax: { total_number: number; }, 
        typ: { total_number: number; }) => (subMax.total_number > typ.total_number ? subMax : typ), 
        { total_number: 0 }).total_number ? max : item), { year: 0 });
    return maxYear.year;
  }

  aggregateTypeData(): void {
    const typeMap: { [name: string]: number } = {};
    
    this.reportData.forEach(item => {
      item.types.forEach((typ: { type: string, total_number: number }) => {
        if (typeMap[typ.type]) {
          typeMap[typ.type] += typ.total_number;
        } else {
          typeMap[typ.type] = typ.total_number;
        }
      });
    });

    this.types = this.types.map(typ => ({
      ...typ,
      total_number: typeMap[typ.name] || 0
    }));
  }

  calculateMaxDelicts(): void {
    this.maxDelicts = Math.max(...this.types.map(typ => typ.total_number));
  }

  calculateMaxDelictsByYear(): void {
    this.maxDelicts = Math.max(...this.reportData.map(item => this.getTotalNumberForYear(item)));
  }

  getBarHeight(type: { name: string, total_number: number }): number {
    return this.maxDelicts ? (type.total_number / this.maxDelicts) * 100 : 0;
  }

  getTotalNumberForYear(item: any): number {
    return item.types.reduce((total: number, typ: { total_number: number }) => total + typ.total_number, 0);
  }

  getBarHeightForYear(item: any): number {
    const totalNumber = this.getTotalNumberForYear(item);
    return this.maxDelicts ? (totalNumber / this.maxDelicts) * 100 : 0;
  }
}
