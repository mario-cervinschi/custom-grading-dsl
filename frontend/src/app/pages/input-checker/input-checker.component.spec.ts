import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InputChecker } from './input-checker.component';

describe('InputChecker', () => {
  let component: InputChecker;
  let fixture: ComponentFixture<InputChecker>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [InputChecker]
    })
      .compileComponents();

    fixture = TestBed.createComponent(InputChecker);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
