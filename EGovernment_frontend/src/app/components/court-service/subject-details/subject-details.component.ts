import { ActivatedRoute, Router } from '@angular/router'; // Dodali smo Router
import { Component, OnInit } from '@angular/core';

import { Subject } from '../../../models/court/subject';
import { SubjectService } from '../../../services/court/subject.service';

@Component({
  selector: 'app-subject-details',
  templateUrl: './subject-details.component.html',
  styleUrls: ['./subject-details.component.css']
})
export class SubjectDetailsComponent implements OnInit {
  subject: Subject | undefined;
  violationDetails: any;

  constructor(
    private route: ActivatedRoute,
    private router: Router, // Dodali smo Router
    private subjectService: SubjectService
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const subjectId = params.get('id');
      if (subjectId) {
        this.subjectService.getSubjectById(subjectId).subscribe(
          subject => {
            this.subject = subject;
            if (subject.violation_id) {
              this.loadViolationDetails(subject.violation_id);
            }
          },
          error => {
            console.error('Error loading subject:', error);
          }
        );
      }
    });
  }

  loadViolationDetails(violationId: string): void {
    this.subjectService.getViolationDetails(violationId).subscribe(
      details => {
        this.violationDetails = details;
        console.log('Violation details:', this.violationDetails);
      },
      error => {
        console.error('Error loading violation details:', error);
      }
    );
  }

  scheduleHearing(): void {
    if (this.subject) {
      this.router.navigate(['rociste', this.subject.id]);
    }
  }
}
