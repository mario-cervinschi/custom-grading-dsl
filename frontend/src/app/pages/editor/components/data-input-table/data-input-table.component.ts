import { CommonModule } from '@angular/common';
import { Component, DestroyRef, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Dialog } from 'primeng/dialog';
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';
import { TableModule } from 'primeng/table';
import { MessageService } from 'primeng/api';
import { NewFileModalComponent } from '../../../../shared/components/new-file-modal/new-file-modal.component';
import { catchError, of, tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { OpenFileModalComponent } from '../../../../shared/components/open-file-modal/open-file-modal.component';
import { EditorService } from '../../../../shared/service/editor.service';
import { PreviewService } from '../../../../shared/service/preview.service';
import { Toast } from 'primeng/toast';

@Component({
  selector: 'app-data-input-table',
  imports: [CommonModule, FormsModule, TableModule, Dialog, Toast],
  providers: [MessageService, DialogService],
  templateUrl: './data-input-table.component.html',
  styleUrl: './data-input-table.component.css',
})
export class DataInputTableComponent {
  protected readonly dataFileSelected = signal('No file');
  protected readonly cols = signal<any[]>([]);
  protected readonly data = signal<any[]>([]);
  protected readonly newRowData = signal<any>({});

  protected readonly newInputDialog = signal(false);
  protected readonly newColumnDialog = signal(false);
  protected readonly newColumnName = signal('');

  protected ref: DynamicDialogRef | null = null;
  private readonly dialogService = inject(DialogService);
  private readonly editorService = inject(EditorService);
  private readonly messageService = inject(MessageService);
  private readonly destroyRef = inject(DestroyRef);
  private readonly previewService = inject(PreviewService);

  openNewFileDialog() {
    this.ref = this.dialogService.open(NewFileModalComponent, {
      header: 'Create New File',
      width: '400px',
      closable: true,
      dismissableMask: true,
      closeOnEscape: true,
      contentStyle: { 'background-color': '#1e1e1e', color: '#ffffff' },
      baseZIndex: 10000,
    });

    this.ref?.onClose.subscribe((fileName: string) => {
      if (fileName && fileName !== '') {
        this.showToast('info', 'Information', `Creating file...`);
        this.editorService
          .createDataFile(fileName)
          .pipe(
            tap((val) => {
              if (val.success)
                this.showToast('success', 'Created', `File "${fileName}" created successfully.`);
            }),
            catchError((err) =>
              of(this.showToast('error', 'Error', `More details: ${err.error.error}`)),
            ),
            takeUntilDestroyed(this.destroyRef),
          )
          .subscribe();
      } else if (fileName === '') {
        this.showToast('error', 'Error', `File name cannot be empty`);
      }
    });
  }

  openOpenFileDialog() {
    this.ref = this.dialogService.open(OpenFileModalComponent, {
      header: 'Open Existing File',
      width: '600px',
      data: { modalType: 'data' },
      closable: true,
      dismissableMask: true,
      closeOnEscape: true,
      contentStyle: { 'max-height': '500px', overflow: 'auto', 'background-color': '#1e1e1e' },
      baseZIndex: 10000,
    });

    this.ref?.onClose
      .pipe(
        catchError((err) =>
          of(this.showToast('error', 'Error', `More details: ${err.error.error}`)),
        ),
      )
      .subscribe((fileContent: any) => {
        if (fileContent) {
          this.dataFileSelected.set(fileContent.file);
          this.cols.set(fileContent.columns || []);
          this.data.set(fileContent.data || []);
          this.previewService.dataFile.set(fileContent.file);
        }
      });
  }

  onSave() {
    this.showToast('info', 'Information', `Saving file...`);
    const stringDataToSave = this.convertDataToINI(this.data());

    this.editorService
      .saveFileContent('data', this.dataFileSelected(), stringDataToSave)
      .pipe(
        tap(() => {
          this.showToast('success', 'Saved', 'File saved successfully.');

          this.previewService.triggerRefresh();
        }),
        catchError((err) =>
          of(this.showToast('error', 'Error', `More details: ${err.error.error}`)),
        ),
      )
      .subscribe();
  }

  openInputDialog() {
    this.newRowData.set({});
    if (!this.cols() || this.cols().length === 0) {
      this.cols.set([{ field: 'id', header: 'ID' }]);
    }
    this.cols().forEach((col) => {
      this.newRowData()[col.field] = '';
    });
    this.newInputDialog.set(true);
  }

  saveNewEntry() {
    const rowId = this.newRowData()['id']?.trim();
    if (!rowId) {
      this.showToast('error', 'Error', 'ID is required.');
      return;
    }

    if (!this.data()) {
      this.data.set([]);
    }

    const exists = this.data().find((r) => r.id === rowId);
    if (exists) {
      this.showToast('error', 'Error', 'An entry with this ID already exists!');
      return;
    }

    this.data.update((old) => [{ ...this.newRowData() }, ...(old || [])]);
    this.newInputDialog.set(false);
  }

  openAddColumnDialog() {
    this.newColumnName.set('');
    this.newColumnDialog.set(true);
  }

  addColumn() {
    const name = this.newColumnName().trim();
    if (!name) return;

    const field = name.toLowerCase().replace(/\s+/g, '_');
    if (this.cols().some((c) => c.field === field)) {
      this.showToast('error', 'Error', 'Column already exists');
      return;
    }

    this.cols.set([...this.cols(), { field, header: name }]);
    this.data.set(this.data().map((row) => ({ ...row, [field]: '' })));
    this.newColumnDialog.set(false);
  }

  private convertDataToINI(tableData: any[]): string {
    if (!tableData || !Array.isArray(tableData)) return '';
    let iniString = '';
    for (const row of tableData) {
      if (!row || !row.id) continue;
      iniString += `[${row.id}]\n`;
      for (const key in row) {
        if (key !== 'id' && row[key] !== undefined && row[key] !== null && row[key] !== '') {
          iniString += `${key} = ${row[key]}\n`;
        }
      }
      iniString += '\n';
    }
    return iniString.trim();
  }

  showToast(type: 'success' | 'error' | 'info', title: string, details: string) {
    this.messageService.add({ severity: type, summary: title, detail: details });
  }
}
