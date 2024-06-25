import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Delict } from 'src/app/models/traffic-police/delict';
import { DelictService } from 'src/app/services/traffic-police/delictService';
import { UserService } from 'src/app/services/auth/user.service';

@Component({
  selector: 'app-delict-details',
  templateUrl: './delict-details.component.html',
  styleUrls: ['./delict-details.component.css']
})
export class DelictDetailsComponent implements OnInit {
  delict: Delict | undefined;
  imagesUrl: any[] = [];
  images: any[] = [];
  currentSlideIndex = 0;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private delictService: DelictService,
    private userService: UserService
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
          this.ngOnInit();
        },
        error => {
          console.error('Error paying fine:', error);
        }
      );
    }
  }

}
