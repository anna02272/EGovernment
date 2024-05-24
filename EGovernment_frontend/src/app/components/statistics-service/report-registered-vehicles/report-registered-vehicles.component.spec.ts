import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ReportRegisteredVehiclesComponent } from './report-registered-vehicles.component';

describe('ReportRegisteredVehiclesComponent', () => {
  let component: ReportRegisteredVehiclesComponent;
  let fixture: ComponentFixture<ReportRegisteredVehiclesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ReportRegisteredVehiclesComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReportRegisteredVehiclesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
