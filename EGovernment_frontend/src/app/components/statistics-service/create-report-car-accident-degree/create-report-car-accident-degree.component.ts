import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DegreeOfAccident } from 'src/app/models/statisics/degreeOfAccident';
import { ReportCarAccidentDegree } from 'src/app/models/statisics/reportCarAccidentDegree';
import { RefreshService } from 'src/app/services/refresh.service';
import { ReportCarAccidentDegreeService } from 'src/app/services/statistics/reportCarAccidentDegree.service';


@Component({
  selector: 'app-create-report-car-accident-degree',
  templateUrl: './create-report-car-accident-degree.component.html',
  styleUrls: ['./create-report-car-accident-degree.component.css']
})
export class CreateReportCarAccidentDegreeComponent implements OnInit {
  reportForm: FormGroup;
  degrees: DegreeOfAccident[] = Object.values(DegreeOfAccident);
  
  constructor(
    private fb: FormBuilder,
    private reportService: ReportCarAccidentDegreeService,
    private refreshService: RefreshService,
    private snackBar: MatSnackBar
  ) {
    this.reportForm = this.fb.group({
      title: ['', Validators.required],
      degree: ['', Validators.required],
      year: ['', [Validators.required, Validators.pattern(/^\d{4}$/)]]
    });
  }

  ngOnInit(): void {
  }

  onSubmit(): void {
    if (this.reportForm.invalid) {
      this.markAllAsTouched();
      return;
    }

    const newReport: ReportCarAccidentDegree = this.reportForm.value;
    this.reportService.create(newReport).subscribe({
      next: () => {
        this.snackBar.open('Report created successfully.', 'Close', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.onCancel();
        this.refreshService.refresh();
      },
      error: () => {
        this.snackBar.open('Failed to create report. Please try again later.', 'Close', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  onCancel(): void {
    this.reportForm.reset();
    this.reportForm.patchValue({
      degree: '' 
    });
  }

  private markAllAsTouched(): void {
    this.reportForm.markAllAsTouched();
  }

}
