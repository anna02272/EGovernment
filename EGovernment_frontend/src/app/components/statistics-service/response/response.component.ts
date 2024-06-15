import { Component, Inject } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { ResponseService } from 'src/app/services/statistics/response.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-response',
  templateUrl: './response.component.html',
  styleUrls: ['./response.component.css']
})
export class ResponseComponent {
  responseForm: FormGroup;
  selectedFile: File | null = null;

  constructor(
    private fb: FormBuilder,
    private responseService: ResponseService,
    private snackBar: MatSnackBar,
    public dialogRef: MatDialogRef<ResponseComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {
    this.responseForm = this.fb.group({
      text: ['', Validators.required],
      accepted: [false],
      send_to: [data.email, [Validators.required, Validators.email]]
    });
  }

  onSubmit(): void {
    if (this.responseForm.invalid) {
      this.markAllAsTouched();
      return;
    }

    const formData = new FormData();
    formData.append('text', this.responseForm.get('text')?.value);
    formData.append('accepted', this.responseForm.get('accepted')?.value);
    formData.append('send_to', this.responseForm.get('send_to')?.value);
    if (this.selectedFile) {
      formData.append('attachment', this.selectedFile, this.selectedFile.name);
    }

    this.responseService.create(formData).subscribe({
      next: () => {
        this.snackBar.open('Response created successfully.', 'Close', {
          duration: 3000,
          panelClass: ['success-snackbar']
        });
        this.dialogRef.close();
      },
      error: () => {
        this.snackBar.open('Failed to create response. Please try again later.', 'Close', {
          duration: 3000,
          panelClass: ['error-snackbar']
        });
      }
    });
  }

  onCancel(): void {
    this.dialogRef.close();
  }

  onFileSelected(event: any): void {
    this.selectedFile = event.target.files[0];
  }

  private markAllAsTouched(): void {
    this.responseForm.markAllAsTouched();
  }
}
