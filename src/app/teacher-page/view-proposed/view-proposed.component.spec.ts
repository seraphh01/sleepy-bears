import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ViewProposedComponent } from './view-proposed.component';

describe('ViewProposedComponent', () => {
  let component: ViewProposedComponent;
  let fixture: ComponentFixture<ViewProposedComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ViewProposedComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ViewProposedComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
