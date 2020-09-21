import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ControlButtonsComponent } from './control-buttons.component';

describe('ControlButtonsComponent', () => {
  let component: ControlButtonsComponent;
  let fixture: ComponentFixture<ControlButtonsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ControlButtonsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ControlButtonsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
