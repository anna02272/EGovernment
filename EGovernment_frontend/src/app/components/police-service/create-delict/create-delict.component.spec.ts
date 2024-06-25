import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CreateDelictComponent } from './create-delict.component';

describe('CreateDelictComponent', () => {
  let component: CreateDelictComponent;
  let fixture: ComponentFixture<CreateDelictComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CreateDelictComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CreateDelictComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});