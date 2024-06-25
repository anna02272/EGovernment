import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AllDelictsComponent } from './all-delicts.component';

describe('AllDelictsComponent', () => {
  let component: AllDelictsComponent;
  let fixture: ComponentFixture<AllDelictsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AllDelictsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AllDelictsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});