import { DatePipe } from '@angular/common';
import { Component } from '@angular/core';
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Request } from 'src/app/models/statisics/request';
import { RefreshService } from 'src/app/services/refresh.service';
import { RequestService } from 'src/app/services/statistics/request.service';
import { ResponseComponent } from '../response/response.component';

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
    private snackBar: MatSnackBar,
    public dialog: MatDialog
  ) { }

  ngOnInit() {
    this.load();
    this.subscribeToRefresh();
  }

  load() {
    this.requestService.getAll().subscribe((data: Request[]) => {
      this.requests = data;
    });
  }

  delete(request: Request) {
    this.requestService.delete(request.id).subscribe({
      next: () => {
        this.snackBar.open('Request deleted successfully.', 'Close', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.refreshService.refresh();
      },
      error: () => {
        this.snackBar.open('Failed to delete request. Please try again later.', 'Close', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  private subscribeToRefresh() {
    this.refreshService.getRefreshObservable().subscribe(() => {
      this.load();
    });
  }

  openResponseDialog(request: Request): void {
    const dialogRef = this.dialog.open(ResponseComponent, {
      width: '550px',
      data: { email: request.email }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.refreshService.refresh();
      }
    });
  }


  formatDate(date: string): string {
    const formattedDate = new Date(date);
    return this.datePipe.transform(formattedDate, 'HH:mm dd/MM/yyyy') || '';
  }
}
