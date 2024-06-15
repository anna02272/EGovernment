import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateReportRegisteredVehiclesComponent } from './create-report-registered-vehicles.component';

describe('CreateReportRegisteredVehiclesComponent', () => {
  let component: CreateReportRegisteredVehiclesComponent;
  let fixture: ComponentFixture<CreateReportRegisteredVehiclesComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [CreateReportRegisteredVehiclesComponent]
    });
    fixture = TestBed.createComponent(CreateReportRegisteredVehiclesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
