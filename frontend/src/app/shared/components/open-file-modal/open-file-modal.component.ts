import { Component, computed, DestroyRef, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { DynamicDialogConfig, DynamicDialogRef } from 'primeng/dynamicdialog';
import { InputTextModule } from 'primeng/inputtext';
import { catchError, of, take, tap } from 'rxjs';
import { EditorService } from '../../service/editor.service';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
  imports: [FormsModule, InputTextModule, ButtonModule],
  templateUrl: './open-file-modal.component.html',
})
export class OpenFileModalComponent implements OnInit {
  protected readonly searchTerm = signal('');
  protected readonly files = signal<any[]>([]);

  protected readonly ref = inject(DynamicDialogRef);
  private readonly destroyRef = inject(DestroyRef);

  private readonly config = inject(DynamicDialogConfig);
  public readonly modalType = this.config.data.modalType as 'data' | 'operations';

  private readonly editorService = inject(EditorService);

  filteredFiles = computed(() =>
    this.files().filter((f) => {
      return f.toLowerCase().includes(this.searchTerm().toLowerCase());
    }),
  );

  ngOnInit() {
    if (this.modalType === 'data') {
      this.editorService
        .getDataFiles()
        .pipe(take(1))
        .subscribe((val) => {
          this.files.set(val);
        });
    } else if (this.modalType === 'operations') {
      this.editorService
        .getOperationsFiles()
        .pipe(take(1))
        .subscribe((val) => {
          this.files.set(val);
        });
    } else {
    }
  }

  selectFile(file: any) {
    if (this.modalType === 'operations') {
      this.editorService
        .readOperationsFile(file)
        .pipe(
          takeUntilDestroyed(this.destroyRef),
          tap((val: any) => this.ref.close({ content: val.content, file })),
          catchError(() => of(this.ref.close())),
        )
        .subscribe();
    } else if (this.modalType === 'data') {
      this.editorService
        .readDataFile(file)
        .pipe(
          takeUntilDestroyed(this.destroyRef),
          tap((val: any) => this.ref.close({ ...val, file })),
          catchError(() => of(this.ref.close())),
        )
        .subscribe();
    } else {
    }
  }
}
