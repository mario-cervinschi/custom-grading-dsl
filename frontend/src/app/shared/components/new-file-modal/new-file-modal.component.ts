import { Component, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DynamicDialogRef } from 'primeng/dynamicdialog';
import { InputTextModule } from 'primeng/inputtext';

@Component({
  imports: [FormsModule, InputTextModule, ButtonModule],
  templateUrl: './new-file-modal.component.html',
})
export class NewFileModalComponent {
  protected readonly fileName = signal('');

  private readonly ref = inject(DynamicDialogRef);

  create() {
    this.ref.close(this.fileName());
  }
}
