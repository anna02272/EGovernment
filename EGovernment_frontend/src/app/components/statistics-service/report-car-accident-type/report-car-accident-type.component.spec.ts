import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportCarAccidentTypeComponent } from './report-car-accident-type.component';

describe('ReportCarAccidentTypeComponent', () => {
  let component: ReportCarAccidentTypeComponent;
  let fixture: ComponentFixture<ReportCarAccidentTypeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportCarAccidentTypeComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReportCarAccidentTypeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
