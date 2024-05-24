import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportDelictsComponent } from './report-delicts.component';

describe('ReportDelictsComponent', () => {
  let component: ReportDelictsComponent;
  let fixture: ComponentFixture<ReportDelictsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportDelictsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReportDelictsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
