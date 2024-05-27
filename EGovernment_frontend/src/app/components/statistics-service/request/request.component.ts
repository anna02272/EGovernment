import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CategoryPerson } from 'src/app/models/statisics/categoryPerson';
import { Request } from 'src/app/models/statisics/request';
import { RequestService } from 'src/app/services/statistics/request.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-request',
  templateUrl: './request.component.html',
  styleUrls: ['./request.component.css']
})
export class RequestComponent implements OnInit {
  requestForm: FormGroup;
  categoriesPerson: CategoryPerson[] = Object.values(CategoryPerson);

  constructor(
    private fb: FormBuilder,
    private requestService: RequestService,
    private snackBar: MatSnackBar
  ) {
    this.requestForm = this.fb.group({
      name: ['', Validators.required],
      lastname: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      phone_number: ['', Validators.required],
      category: ['', Validators.required],
      question: ['', Validators.required]
    });
  }

  ngOnInit(): void {}

  onSubmit(): void {
    if (this.requestForm.invalid) {
      this.markAllAsTouched();
      return;
    }

    const newRequest: Request = this.requestForm.value;
    this.requestService.create(newRequest).subscribe({
      next: response => {
        this.snackBar.open('Request created successfully.', 'Close', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.onCancel();
      },
      error: err => {
        this.snackBar.open('Failed to create request. Please try again later.', 'Close', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  onCancel(): void {
    this.requestForm.reset();
    this.requestForm.patchValue({
      category: '' 
    });
  }

  private markAllAsTouched(): void {
    this.requestForm.markAllAsTouched();
  }

}
