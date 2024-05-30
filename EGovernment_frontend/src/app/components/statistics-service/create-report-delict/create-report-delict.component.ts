import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { DelictType } from 'src/app/models/statisics/delictType';
import { ReportDelict } from 'src/app/models/statisics/reportDelict';
import { RefreshService } from 'src/app/services/refresh.service';
import { ReportDelicTypeService } from 'src/app/services/statistics/reportDelicType.service';

@Component({
  selector: 'app-create-report-delict',
  templateUrl: './create-report-delict.component.html',
  styleUrls: ['./create-report-delict.component.css']
})
export class CreateReportDelictComponent implements OnInit {
    reportForm: FormGroup;
    types: DelictType[] = Object.values(DelictType);
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
    constructor(
      private fb: FormBuilder,
      private reportService: ReportDelicTypeService,
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
  
      const newReport: ReportDelict = this.reportForm.value;
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
  