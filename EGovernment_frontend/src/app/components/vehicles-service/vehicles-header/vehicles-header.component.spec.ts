import { ComponentFixture, TestBed } from '@angular/core/testing';

import { VehiclesHeaderComponent } from './vehicles-header.component';

describe('VehiclesHeaderComponent', () => {
  let component: VehiclesHeaderComponent;
  let fixture: ComponentFixture<VehiclesHeaderComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [VehiclesHeaderComponent]
    });
    fixture = TestBed.createComponent(VehiclesHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
