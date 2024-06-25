import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PoliceHeaderComponent } from './police-header.component';

describe('PoliceHeaderComponent', () => {
  let component: PoliceHeaderComponent;
  let fixture: ComponentFixture<PoliceHeaderComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PoliceHeaderComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PoliceHeaderComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
