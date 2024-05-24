import { DatePipe } from '@angular/common';
import { Component } from '@angular/core';
import { Request } from 'src/app/models/statisics/request';
import { RefreshService } from 'src/app/services/refresh.service';
import { RequestService } from 'src/app/services/statistics/request.service';

@Component({
  selector: 'app-requests',
  templateUrl: './requests.component.html',
  styleUrls: ['./requests.component.css'],
  providers: [DatePipe]
})
export class RequestsComponent {
  requests: Request[] = [];
  
  constructor(
    private requestService: RequestService,
    private datePipe: DatePipe,
    private refreshService: RefreshService,
  ) {}
  
  ngOnInit() {
    this.load();
    this.subscribeToRefresh();
  }
  load() {
    this.requestService.getAll().subscribe((data: Request[]) => {
      this.requests = data;
    });
  }
    private subscribeToRefresh() {
      this.refreshService.getRefreshObservable().subscribe(() => {
        this.load();
      });
    }

    formatDate(date: string): string {
      const formattedDate = new Date(date);
      return this.datePipe.transform(formattedDate, 'HH:mm dd/MM/yyyy') || '';
    }
}
