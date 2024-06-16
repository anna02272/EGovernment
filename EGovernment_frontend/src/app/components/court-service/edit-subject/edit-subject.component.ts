import { ActivatedRoute, Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';

import { Subject } from '../../../models/court/subject';
import { SubjectService } from '../../../services/court/subject.service';
import { forkJoin } from 'rxjs';

@Component({
  selector: 'app-edit-subject',
  templateUrl: './edit-subject.component.html',
  styleUrls: ['./edit-subject.component.css']
})
export class EditSubjectComponent implements OnInit {
  subject: Subject | undefined;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private subjectService: SubjectService
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const subjectId = params.get('id');
      if (subjectId) {
        this.subjectService.getSubjectById(subjectId).subscribe(
          subject => {
            this.subject = subject;
          },
          error => {
            console.error('Error loading subject:', error);
          }
        );
      }
    });
  }

  saveChanges(): void {
    if (this.subject) {
      const { id, judgment, status, compromis } = this.subject;
      
      // Napravite tri odvojena zahteva za ažuriranje svake vrste podatka
      const updateStatus$ = this.subjectService.updateSubjectStatus(id, status);
      const updateJudgment$ = this.subjectService.updateSubjectJudgment(id, judgment);
      const updateCompromise$ = this.subjectService.updateSubjectCompromis(id, compromis);

      // Kombinujte ove observable i izvršite ih istovremeno
      forkJoin([updateStatus$, updateJudgment$, updateCompromise$]).subscribe(
        ([updatedStatus, updatedJudgment, updatedCompromise]) => {
          console.log('Successfully updated subject:', updatedStatus, updatedJudgment, updatedCompromise);
          this.router.navigate(['/subjectTab', id]); // Navigate back to the subject details
        },
        error => {
          console.error('Error updating subject:', error);
        }
      );
    }
  }

  cancel(): void {
    if (this.subject) {
      this.router.navigate(['/subjectTab', this.subject.id]); // Navigate back to the subject details
    }
  }
}
