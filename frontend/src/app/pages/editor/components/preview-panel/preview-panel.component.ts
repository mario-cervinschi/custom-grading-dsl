import { Component, computed, inject, signal } from '@angular/core';
import { PreviewService } from '../../../../shared/service/preview.service';
import { EditorService } from '../../../../shared/service/editor.service';
import { combineLatest, debounceTime, filter } from 'rxjs';
import { toObservable } from '@angular/core/rxjs-interop';
import { CommonModule } from '@angular/common';
import { ScrollerModule } from 'primeng/scroller';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { PaginatorModule, PaginatorState } from 'primeng/paginator';

@Component({
  selector: 'app-preview-panel',
  imports: [CommonModule, ProgressSpinnerModule, ScrollerModule, PaginatorModule],
  templateUrl: './preview-panel.component.html',
  styleUrl: './preview-panel.component.css',
})
export class PreviewPanelComponent {
  private previewService = inject(PreviewService);
  private editorService = inject(EditorService);

  first = signal(0);
  rows = signal(1);

  visibleResults = computed(() => {
    const allResults = this.result()?.results || [];
    const startIndex = this.first();
    const endIndex = startIndex + this.rows();
    return allResults.slice(startIndex, endIndex);
  });

  onPageChange(event: PaginatorState) {
    this.first.set(event.first || 0);
    this.rows.set(event.rows || 1);

    document.querySelector('.overflow-y-auto')?.scrollTo(0, 0);
  }

  result = this.previewService.preview;
  loading = signal(false);

  constructor() {
    combineLatest([
      toObservable(this.previewService.operationsFile),
      toObservable(this.previewService.dataFile),
      toObservable(this.previewService.refreshTrigger),
    ])
      .pipe(
        debounceTime(500),
        filter(([ops, data, _trigger]) => !!ops && !!data),
      )
      .subscribe(([ops, data, _trigger]) => {
        this.loading.set(true);

        this.editorService.getPreview(ops!, data!).subscribe({
          next: (res: any) => {
            this.previewService.preview.set(res);
            this.loading.set(false);
          },
          error: () => {
            this.previewService.preview.set(null);
            this.loading.set(false);
          },
        });
      });
  }
}
