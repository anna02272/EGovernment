import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Delict } from 'src/app/models/traffic-police/delict';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { UserService } from 'src/app/services/auth/user.service';
import { MatSnackBar } from '@angular/material/snack-bar';

@Component({
  selector: 'app-delict-details',
  templateUrl: './delict-details.component.html',
  styleUrls: ['./delict-details.component.css']
})
export class DelictDetailsComponent implements OnInit {
  delict: Delict | undefined;
  imagesUrl: any[] = [];
  images: any[] = [];
  imageFiles: File[] = [];
  currentSlideIndex = 0;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private delictService: DelictService,
    private userService: UserService,
    private snackBar: MatSnackBar,
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const delictId = params.get('id');
      if (delictId) {
        this.delictService.getDelictById(delictId).subscribe(
          delict => {
            this.delict = delict;
            this.loadImages(delict.id);
          },
          error => {
            console.error('Error loading delict:', error);
          }
        );
      }
    });
  }


  loadImages(delictId: string): void {
    this.delictService.getImagesUrls(delictId).subscribe(
      (images: any) => {
        this.imagesUrl = images;
        if (this.imagesUrl.length !== 0) {
          this.imagesUrl.forEach(imageName => this.loadImageContent(delictId, imageName));
        }
      },
      error => {
        console.error('Error getting delict images:', error);
      }
    );
  }

  loadImageContent(folderName: string, imageName: string): void {
    this.delictService.getImages(folderName, imageName).subscribe(
      (imageContent: any) => {
        this.displayImage(imageContent);
      },
      error => {
        console.error('Error getting images:', error);
      }
    );
  }

  displayImage(blob: Blob) {
    const reader = new FileReader();
    reader.onload = (e: any) => {
      this.images.push(e.target.result);
    };
    reader.readAsDataURL(blob);
  }

  onFileChange(event: any): void {
    if (event.target.files && event.target.files.length) {
      this.imageFiles = Array.from(event.target.files);
    }
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

    if (this.delict) {
      this.delictService.uploadImages(this.delict.id, formData).subscribe(
        () => {
          console.log('Images uploaded successfully!');
          this.openSnackBar("Images uploaded successfully!", "");
          this.ngOnInit(); // Reload the page after successful upload
        },
        (error) => {
          console.error('Error uploading images:', error);
          this.openSnackBar("Error uploading images!", "");
        }
      );
    } else {
      console.error('Delict is undefined, cannot upload images.');
    }
  }


  prevSlide() {
    if (this.currentSlideIndex > 0) {
      this.currentSlideIndex--;
    }
  }

  nextSlide() {
    if (this.currentSlideIndex < this.images.length - 1) {
      this.currentSlideIndex++;
    }
  }

  downloadPdf(): void {
    if (this.delict) {
      this.delictService.getPdfByDelictId(this.delict.id).subscribe(
        (pdfBlob: Blob) => {
          const url = window.URL.createObjectURL(pdfBlob);
          const a = document.createElement('a');
          a.href = url;
          a.download = `Delict_${this.delict?.id}.pdf`;  // Use optional chaining
          a.click();
          window.URL.revokeObjectURL(url);
        },
        error => {
          console.error('Error downloading PDF:', error);
        }
      );
    } else {
      console.error('Delict is undefined, cannot download PDF.');
    }
  }

  getRole() {
    return this.userService.currentUser?.user.userRole;
  }

  payFine() {
    if (this.delict) {
      this.delictService.updateDelictStatus(this.delict.id).subscribe(
        () => {
          console.log('Fine paid successfully');
          // Optionally refresh the delict details
          this.openSnackBar("Uspesno ste isplatili kaznu!", "");
          this.ngOnInit();
        },
        error => {
          console.error('Error paying fine:', error);
        }
      );
    }
  }

  openSnackBar(message: string, action: string) {
    this.snackBar.open(message, action, {
      duration: 3000,
    });
  }



}
