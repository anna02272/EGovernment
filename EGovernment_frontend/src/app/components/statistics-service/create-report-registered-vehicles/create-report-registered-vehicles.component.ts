import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Category } from 'src/app/models/statisics/category';
import { ReportRegisteredVehicle } from 'src/app/models/statisics/reportRegisteredVehicle';
import { RefreshService } from 'src/app/services/refresh.service';
import { ReportRegisteredVehiclesService } from 'src/app/services/statistics/reportRegisteredVehicles.service';

@Component({
  selector: 'app-create-report-registered-vehicles',
  templateUrl: './create-report-registered-vehicles.component.html',
  styleUrls: ['./create-report-registered-vehicles.component.css']
})
export class CreateReportRegisteredVehiclesComponent implements OnInit {
  reportForm: FormGroup;
  categories: Category[] = Object.values(Category);
  constructor(
    private fb: FormBuilder,
    private reportService: ReportRegisteredVehiclesService,
    private refreshService: RefreshService,
    private snackBar: MatSnackBar
  ) {
    this.reportForm = this.fb.group({
      title: ['', Validators.required],
      category: ['', Validators.required],
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

    const newReport: ReportRegisteredVehicle = this.reportForm.value;
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
      category: '' 
    });
  }

  private markAllAsTouched(): void {
    this.reportForm.markAllAsTouched();
  }

}
