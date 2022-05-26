import { ComponentFixture, TestBed } from '@angular/core/testing';

import { BestResultsComponent } from './best-results.component';

describe('BestResultsComponent', () => {
  let component: BestResultsComponent;
  let fixture: ComponentFixture<BestResultsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ BestResultsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(BestResultsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
