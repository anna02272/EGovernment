import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportCarAccidentDegreeComponent } from './report-car-accident-degree.component';

describe('ReportCarAccidentDegreeComponent', () => {
  let component: ReportCarAccidentDegreeComponent;
  let fixture: ComponentFixture<ReportCarAccidentDegreeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportCarAccidentDegreeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReportCarAccidentDegreeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
