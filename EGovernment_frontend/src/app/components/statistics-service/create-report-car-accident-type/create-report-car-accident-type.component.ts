import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { CarAccidentType } from 'src/app/models/statisics/carAccidentType';
import { ReportCarAccidentType } from 'src/app/models/statisics/reportCarAccidentType';
import { RefreshService } from 'src/app/services/refresh.service';
import { ReportCarAccidentTypeService } from 'src/app/services/statistics/reportCarAccidentTypeService.service';

@Component({
  selector: 'app-create-report-car-accident-type',
  templateUrl: './create-report-car-accident-type.component.html',
  styleUrls: ['./create-report-car-accident-type.component.css']
})
export class CreateReportCarAccidentTypeComponent  implements OnInit {
  reportForm: FormGroup;
  types: CarAccidentType[] = Object.values(CarAccidentType);

  constructor(
    private fb: FormBuilder,
    private reportService: ReportCarAccidentTypeService,
    private refreshService: RefreshService,
    private snackBar: MatSnackBar
  ) {
    this.reportForm = this.fb.group({
      title: ['', Validators.required],
      type: ['', Validators.required],
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

    const newReport: ReportCarAccidentType = this.reportForm.value;
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
      type: '' 
    });
  }

  private markAllAsTouched(): void {
    this.reportForm.markAllAsTouched();
  }

}

