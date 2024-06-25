import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { DelictCreate } from 'src/app/models/traffic-police/delictCreate';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { DelictType} from 'src/app/models/traffic-police/delictType';
import { DelictStatus} from 'src/app/models/traffic-police/delictStatus';
import { UserService } from 'src/app/services/auth/user.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-create-delict',
  templateUrl: './create-delict.component.html',
  styleUrls: ['./create-delict.component.css']
})
export class CreateDelictComponent {
  delict: DelictCreate = new DelictCreate(
   '', '', '', '', '', '',DelictType.Speeding, DelictStatus.FineAwarded, 0, 0
  );

  delictTypes = Object.values(DelictType);
  delictStatuses = Object.values(DelictStatus);
  imageFiles: File[] = [];
  responseId: any;

  constructor(private delictService: DelictService, private router: Router,  private userService: UserService, private snackBar: MatSnackBar,) {}


  getCurrentUserId() {
    return this.userService.currentUser?.user.id;
  }

  onFileChange(event: any): void {
    if (event.target.files && event.target.files.length) {
      this.imageFiles = Array.from(event.target.files);
    }
  }

  createDelict(): void {
    if (!this.delict.delict_type || !this.delict.delict_status) {
      alert('Please select both delict type and status.');
      return;
    }

    this.delictService.insertDelict(this.delict).subscribe(
      (response: any) => {
        this.responseId = response.id;
        this.uploadImages();
      },
      error => {
        console.error('Error creating delict:', error);
      }
    );
  }

  uploadImages(): void {
    if (this.imageFiles.length === 0) {
      this.openSnackBar("No images selected for upload!", "");
      console.warn('No images selected for upload.');
      return;
    }

    const formData = new FormData();
    for (const file of this.imageFiles) {
      formData.append('images', file);
    }

    this.delictService.uploadImages(this.responseId, formData).subscribe(
      () => {
        console.log('Images uploaded successfully!');
        this.openSnackBar("Images uploaded successfully!", "");
        this.router.navigate(['/saobracajnaPolicija']);  // Navigate to delicts page after successful upload
      },
      (error) => {
        console.error('Error uploading images:', error);
        this.openSnackBar("Error uploading images!", "");
      }
    );
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 2000,
    });
  }
}