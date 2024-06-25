import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DelictDetailsComponent } from './delict-details.component';

describe('DelictDetailsComponent', () => {
  let component: DelictDetailsComponent;
  let fixture: ComponentFixture<DelictDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DelictDetailsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DelictDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});