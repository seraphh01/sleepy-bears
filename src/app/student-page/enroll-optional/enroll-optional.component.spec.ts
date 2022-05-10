import { ComponentFixture, TestBed } from '@angular/core/testing';

import { EnrollOptionalComponent } from './enroll-optional.component';

describe('EnrollOptionalComponent', () => {
  let component: EnrollOptionalComponent;
  let fixture: ComponentFixture<EnrollOptionalComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ EnrollOptionalComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(EnrollOptionalComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
