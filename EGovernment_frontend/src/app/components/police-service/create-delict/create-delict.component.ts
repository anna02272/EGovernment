import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { DelictCreate } from 'src/app/models/traffic-police/delictCreate';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { DelictType } from 'src/app/models/traffic-police/delictType';
import { DelictStatus } from 'src/app/models/traffic-police/delictStatus';
import { UserService } from 'src/app/services/auth/user.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-create-delict',
  templateUrl: './create-delict.component.html',
  styleUrls: ['./create-delict.component.css']
})
export class CreateDelictComponent {
  delict: DelictCreate = new DelictCreate('', '', '', '', '', DelictType.Speeding, DelictStatus.FineAwarded, 0, 0);

  delictTypes = Object.values(DelictType);
  delictStatuses = Object.values(DelictStatus);
  imageFiles: File[] = [];
  responseId: any;

  constructor(private delictService: DelictService, private router: Router, private userService: UserService, private snackBar: MatSnackBar) {}

  getCurrentUserId() {
    return this.userService.currentUser?.user.id;
  }

  onFileChange(event: any): void {
    if (event.target.files && event.target.files.length) {
      this.imageFiles = Array.from(event.target.files);
    }
  }

  createDelict(): void {
    // Validate form fields
    if (!this.isFormValid()) {
      this.openSnackBar("Molimo popunite sva polja i unesite pozitivnu vrednost za 'Novčana kazna' i 'Broj kaznenih poena'.", "");
      return;
    }

    this.delictService.insertDelict(this.delict).subscribe(
      (response: any) => {
        this.responseId = response.id;
        this.uploadImages();
        this.openSnackBar("Uspešno ste kreirali prekršaj!", "");
        this.router.navigate(['/all-delicts']);
      },
      error => {
        console.error('Error creating delict:', error);
      }
    );
  }

  isFormValid(): boolean {
    return (
      !!this.delict.driver_identification_number &&
      !!this.delict.vehicle_licence_number &&
      !!this.delict.driver_email &&
      !!this.delict.location &&
      !!this.delict.description &&
      !!this.delict.delict_type &&
      !!this.delict.delict_status &&
      this.delict.price_of_fine >= 0 &&
      this.delict.number_of_penalty_points >= 0
    );
  }

  uploadImages(): void {
    if (this.imageFiles.length === 0) {
      this.openSnackBar("Niste izabrali slike za upload!", "");
      console.warn('No images selected for upload.');
      return;
    }

    const formData = new FormData();
    for (const file of this.imageFiles) {
      formData.append('images', file);
    }

    this.delictService.uploadImages(this.responseId, formData).subscribe(
      () => {
        //console.log('Slike uspešno učitane!');
        this.router.navigate(['/all-delicts']);
      },
      (error) => {
        console.error('Greška prilikom uploada slika:', error);
        this.openSnackBar("Greška prilikom uploada slika!", "");
      }
    );
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 3000,
    });
  }
}
