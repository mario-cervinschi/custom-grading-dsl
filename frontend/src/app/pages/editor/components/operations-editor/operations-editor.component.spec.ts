import { ComponentFixture, TestBed } from '@angular/core/testing';

import { OperationsEditorComponent } from './operations-editor.component';

describe('OperationsEditorComponent', () => {
  let component: OperationsEditorComponent;
  let fixture: ComponentFixture<OperationsEditorComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [OperationsEditorComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(OperationsEditorComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
