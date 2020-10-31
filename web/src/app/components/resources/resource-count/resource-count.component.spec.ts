import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ResourceCountComponent } from './resource-count.component';

describe('ResourceCountComponent', () => {
  let component: ResourceCountComponent;
  let fixture: ComponentFixture<ResourceCountComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ResourceCountComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ResourceCountComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
