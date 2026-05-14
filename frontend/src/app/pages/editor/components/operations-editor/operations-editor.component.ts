import { CommonModule } from '@angular/common';
import { Component, DestroyRef, inject, signal } from '@angular/core';
import { takeUntilDestroyed, toObservable } from '@angular/core/rxjs-interop';
import { FormsModule } from '@angular/forms';
import { MonacoEditorModule } from 'ngx-monaco-editor-v2';
import { DialogService, DynamicDialogRef } from 'primeng/dynamicdialog';
import { LspService } from '../../../../core/service/lsp-client';
import { MessageService } from 'primeng/api';
import { catchError, debounceTime, distinctUntilChanged, of, skip, tap } from 'rxjs';
import { NewFileModalComponent } from '../../../../shared/components/new-file-modal/new-file-modal.component';
import { OpenFileModalComponent } from '../../../../shared/components/open-file-modal/open-file-modal.component';
import { EditorService } from '../../../../shared/service/editor.service';
import { PreviewService } from '../../../../shared/service/preview.service';
import { Toast } from 'primeng/toast';

@Component({
  selector: 'app-operations-editor',
  imports: [CommonModule, FormsModule, MonacoEditorModule, Toast],
  providers: [MessageService, DialogService],
  templateUrl: './operations-editor.component.html',
  styleUrl: './operations-editor.component.css',
})
export class OperationsEditorComponent {
  protected readonly operationsFileSelected = signal('No file');
  protected readonly editorContent = signal('');
  private readonly editorContent$ = toObservable(this.editorContent);
  private editorInstance: any = null;
  private lastValidation: { valid: boolean; errors: any[] } = { valid: true, errors: [] };

  protected ref: DynamicDialogRef | null = null;
  protected readonly EDITOR_OPTIONS = {
    language: 'codesuggestion',
    theme: 'vs-dark',
    fontSize: 14,
    minimap: { enabled: false },
    automaticLayout: true,
    quickSuggestions: true,
    suggestOnTriggerCharacters: true,
    fixedOverflowWidgets: true,
  };

  private readonly lspClient = inject(LspService);
  private readonly dialogService = inject(DialogService);
  private readonly editorService = inject(EditorService);
  private readonly messageService = inject(MessageService);
  private readonly destroyRef = inject(DestroyRef);
  private readonly previewService = inject(PreviewService);

  ngOnInit() {
    this.editorContent$
      .pipe(
        skip(1),
        debounceTime(2000),
        distinctUntilChanged(),
        takeUntilDestroyed(this.destroyRef),
      )
      .subscribe(() => {
        this.onSave();
      });
  }

  async onEditorInit(editor: any): Promise<void> {
    this.editorInstance = editor;
    const monaco = (window as any).monaco;
    monaco.languages.register({ id: 'codesuggestion' });

    try {
      await this.lspClient.connect();
      this.lspClient.openDocument(this.editorContent());

      monaco.languages.registerCompletionItemProvider('codesuggestion', {
        triggerCharacters: [' '],
        provideCompletionItems: async (model: any, position: any) => {
          const items = await this.lspClient.getCompletions(
            position.lineNumber - 1,
            position.column - 1,
          );
          if (!items || items.length === 0) return { suggestions: [] };

          return {
            suggestions: items.map((item: any) => {
              const isSnippet = item.insertTextFormat === 2;
              return {
                label: item.label,
                kind: item.kind,
                insertText: item.insertText || item.label,
                insertTextRules: isSnippet ? 4 : undefined,
                detail: item.detail || '',
              };
            }),
          };
        },
      });
    } catch (e) {
      console.error('Could not connect to LSP', e);
    }
  }

  onContentChange(value: string): void {
    this.lspClient.updateDocument(value);
    this.runValidation(value);
  }

  private async runValidation(text: string) {
    if (!this.editorInstance) return;
    const monaco = (window as any).monaco;
    const model = this.editorInstance.getModel();
    if (!model) return;

    const result = await this.lspClient.validate(text);
    this.lastValidation = result;
    const markers = result.errors.map((e: any) => ({
      severity: monaco.MarkerSeverity.Error,
      message: e.msg,
      startLineNumber: e.line + 1,
      startColumn: e.col + 1,
      endLineNumber: e.endLine + 1,
      endColumn: e.endCol + 1,
    }));
    monaco.editor.setModelMarkers(model, 'dsl-validation', markers);
    this.previewService.hasSyntaxErrors.set(result.errors.length > 0);
  }

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
          .createOperationsFile(fileName)
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
      data: { modalType: 'operations' },
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
          this.operationsFileSelected.set(fileContent.file);
          this.editorContent.set(fileContent.content);

          this.previewService.operationsFile.set(fileContent.file);
        }
      });
  }

  async onSave() {
    if (!this.lastValidation.valid) {
      const first = this.lastValidation.errors[0];
      this.showToast(
        'error',
        'Syntax Error',
        `Line ${first.line + 1}, col ${first.col + 1}: ${first.msg}`,
      );
      // return;
    }

    this.editorService
      .saveFileContent('operations', this.operationsFileSelected(), this.editorContent())
      .pipe(
        tap((val) => {
          console.log(val);
          this.previewService.triggerRefresh();
        }),
      )
      .subscribe();
  }

  showToast(type: 'success' | 'error' | 'info', title: string, details: string) {
    this.messageService.add({ severity: type, summary: title, detail: details });
  }
}
