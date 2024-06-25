import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DelictComponent } from './delict.component';

describe('DelictComponent', () => {
  let component: DelictComponent;
  let fixture: ComponentFixture<DelictComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ DelictComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DelictComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});