import { ActivatedRoute, Router } from '@angular/router'; // Dodali smo Router
import { Component, OnInit } from '@angular/core';

import { Schedule } from 'src/app/models/court/schedule';
import { ScheduleService } from 'src/app/services/court/schedule.service';
import { Subject } from '../../../models/court/subject';
import { SubjectService } from '../../../services/court/subject.service';

@Component({
  selector: 'app-subject-tab',
  templateUrl: './subject-tab.component.html',
  styleUrls: ['./subject-tab.component.css']
})
export class SubjectTabComponent implements OnInit {
  subject: Subject | undefined;
  violationDetails: any;
  schedule: Schedule | undefined;
  isEditing: boolean = false;



  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private subjectService: SubjectService,
    private scheduleService: ScheduleService

  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe(params => {
      const hearingId = params.get('hearingId');
      if (hearingId) {
        console.log(hearingId)
        this.loadScheduleByHearingId(hearingId);
      }
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
  loadScheduleByHearingId(hearingId: string): void {
    this.scheduleService.getScheduleByHearingId(hearingId).subscribe(
      schedule => {
        this.schedule = schedule;
        console.log('Schedule:', this.schedule);
      },
      error => {
        console.error('Error loading schedule:', error);
      }
    );
  }

  editSubject(): void {
    if (this.subject) {
      this.router.navigate(['editSubject', this.subject.id]);
    }
  }
}
