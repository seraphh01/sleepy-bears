import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WorstResultsComponent } from './worst-results.component';

describe('WorstResultsComponent', () => {
  let component: WorstResultsComponent;
  let fixture: ComponentFixture<WorstResultsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ WorstResultsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(WorstResultsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
