import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DriverLicenceComponent } from './driver-licence.component';

describe('DriverLicenceComponent', () => {
  let component: DriverLicenceComponent;
  let fixture: ComponentFixture<DriverLicenceComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [DriverLicenceComponent]
    });
    fixture = TestBed.createComponent(DriverLicenceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
