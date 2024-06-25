import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllCarAccidentsComponent } from './all-car-accidents.component';

describe('AllCarAccidentsComponent', () => {
  let component: AllCarAccidentsComponent;
  let fixture: ComponentFixture<AllCarAccidentsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllCarAccidentsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllCarAccidentsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});