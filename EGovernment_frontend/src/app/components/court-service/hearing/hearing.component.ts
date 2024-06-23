import { ActivatedRoute, Router } from '@angular/router';
import { Component, OnInit } from '@angular/core';

import { AuthService } from 'src/app/services/auth/auth.service';
import { Court } from '../../../models/court/court';
import { CourtService } from '../../../services/court/court.service';
import { Hearing } from '../../../models/court/hearing';
import { HearingService } from '../../../services/court/hearing.service';
import { HttpErrorResponse } from '@angular/common/http';
import { Schedule } from '../../../models/court/schedule';
import { ScheduleService } from '../../../services/court/schedule.service';

@Component({
  selector: 'app-hearing',
  templateUrl: './hearing.component.html',
  styleUrls: ['./hearing.component.css']
})
export class HearingComponent implements OnInit {
  hearingData: Partial<Hearing> = {};
  scheduleData: Partial<Schedule> = {};
  courts: Court[] = [];
  selectedCourts: string[] = [];
  subjectId: string | null = null;

  constructor(
    private hearingService: HearingService,
    private scheduleService: ScheduleService,
    private courtService: CourtService,
    private router: Router,
    private activatedRoute: ActivatedRoute,
    private authService: AuthService
  ) { }

  ngOnInit(): void {
    this.activatedRoute.paramMap.subscribe(params => {
      this.subjectId = params.get('subjectId');
    });

    this.loadCourts();
  }

  loadCourts(): void {
    this.courtService.getAllCourts().subscribe(
      courts => {
        this.courts = courts;
      },
      error => {
        console.error('Greška pri učitavanju sudova:', error);
      }
    );
  }

  onSubmit(): void {
    if (this.hearingData.date && this.subjectId) {
      const newHearing: Hearing = {
        subject_id: this.subjectId,
        date: this.hearingData.date,
        
      };

      const token = this.authService.getToken();

      if (!token) {
        console.error('JWT token nije dostupan.');
        return;
      }

      const headers = {
        Authorization: `Bearer ${token}`
      };

      const options = {
        headers: headers
      };

      this.hearingService.createHearing(newHearing, options).subscribe(
        (response: any) => {
          const hearingId = response.id;
          if (hearingId && this.selectedCourts.length > 0 && this.scheduleData.start_time && this.scheduleData.end_time) {
            this.selectedCourts.forEach(courtId => {
              const newSchedule: Schedule = {
                hearing_id: hearingId,
                court_id: courtId,
                start_time: String(this.scheduleData.start_time),
                end_time: String(this.scheduleData.end_time)
              };

              this.scheduleService.createSchedule(newSchedule).subscribe(
                scheduleResponse => {
                  console.log('Hearing successfully scheduled for court:', scheduleResponse);
                },
                (error: HttpErrorResponse) => {
                  console.error('Greška pri kreiranju rasporeda:', error);
                }
              );
            });

            this.router.navigate(['/subject-details']);
          }
        },
        (error: HttpErrorResponse) => {
          console.error('Greška pri kreiranju slučaja:', error);
        }
      );
    }
  }
}
