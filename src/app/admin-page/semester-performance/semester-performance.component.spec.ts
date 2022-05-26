import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SemesterPerformanceComponent } from './semester-performance.component';

describe('SemesterPerformanceComponent', () => {
  let component: SemesterPerformanceComponent;
  let fixture: ComponentFixture<SemesterPerformanceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SemesterPerformanceComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SemesterPerformanceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
