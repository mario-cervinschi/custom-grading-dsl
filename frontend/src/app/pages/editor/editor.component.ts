import { Component } from '@angular/core';
import { ToastModule } from 'primeng/toast';
import { Splitter } from 'primeng/splitter';
import { DataInputTableComponent } from './components/data-input-table/data-input-table.component';
import { PreviewPanelComponent } from './components/preview-panel/preview-panel.component';
import { OperationsEditorComponent } from './components/operations-editor/operations-editor.component';
import { MessageService } from 'primeng/api';
import { DialogService } from 'primeng/dynamicdialog';

@Component({
  selector: 'app-editor',
  imports: [
    Splitter,
    ToastModule,
    OperationsEditorComponent,
    DataInputTableComponent,
    PreviewPanelComponent,
  ],
  providers: [MessageService, DialogService],
  templateUrl: './editor.component.html',
  styleUrl: './editor.component.css',
})
export class Editor {}
