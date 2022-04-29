import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChiefPageComponent } from './chief-page.component';

describe('ChiefPageComponent', () => {
  let component: ChiefPageComponent;
  let fixture: ComponentFixture<ChiefPageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChiefPageComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ChiefPageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
