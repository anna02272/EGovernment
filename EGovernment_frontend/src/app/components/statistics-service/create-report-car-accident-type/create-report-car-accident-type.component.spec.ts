import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateReportCarAccidentTypeComponent } from './create-report-car-accident-type.component';

describe('CreateReportCarAccidentTypeComponent', () => {
  let component: CreateReportCarAccidentTypeComponent;
  let fixture: ComponentFixture<CreateReportCarAccidentTypeComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateReportCarAccidentTypeComponent]
    });
    fixture = TestBed.createComponent(CreateReportCarAccidentTypeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
