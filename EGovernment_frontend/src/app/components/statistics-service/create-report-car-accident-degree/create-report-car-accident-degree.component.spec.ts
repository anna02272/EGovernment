import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateReportCarAccidentDegreeComponent } from './create-report-car-accident-degree.component';

describe('CreateReportCarAccidentDegreeComponent', () => {
  let component: CreateReportCarAccidentDegreeComponent;
  let fixture: ComponentFixture<CreateReportCarAccidentDegreeComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateReportCarAccidentDegreeComponent]
    });
    fixture = TestBed.createComponent(CreateReportCarAccidentDegreeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
