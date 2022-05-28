import { ComponentFixture, TestBed } from '@angular/core/testing';

import { YearPerformanceComponent } from './year-performance.component';

describe('YearPerformanceComponent', () => {
  let component: YearPerformanceComponent;
  let fixture: ComponentFixture<YearPerformanceComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ YearPerformanceComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(YearPerformanceComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
