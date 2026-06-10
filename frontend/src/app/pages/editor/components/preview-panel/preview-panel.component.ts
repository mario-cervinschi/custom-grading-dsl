import { Component, computed, inject, signal } from '@angular/core';
import { PreviewService } from '../../../../shared/service/preview.service';
import { EditorService } from '../../../../shared/service/editor.service';
import { combineLatest, debounceTime, filter, switchMap, tap, catchError, of } from 'rxjs';
import { takeUntilDestroyed, toObservable } from '@angular/core/rxjs-interop';
import { CommonModule } from '@angular/common';
import { ScrollerModule } from 'primeng/scroller';
import { ProgressSpinnerModule } from 'primeng/progressspinner';
import { PaginatorModule, PaginatorState } from 'primeng/paginator';
import { DestroyRef } from '@angular/core';

@Component({
  selector: 'app-preview-panel',
  imports: [CommonModule, ProgressSpinnerModule, ScrollerModule, PaginatorModule],
  templateUrl: './preview-panel.component.html',
  styleUrl: './preview-panel.component.css',
})
export class PreviewPanelComponent {
  private previewService = inject(PreviewService);
  private editorService = inject(EditorService);
  private destroyRef = inject(DestroyRef);

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
  hasSyntaxErrors = this.previewService.hasSyntaxErrors;
  operationsFile = this.previewService.operationsFile;
  dataFile = this.previewService.dataFile;
  loading = signal(false);

  constructor() {
    combineLatest([
      toObservable(this.previewService.operationsFile),
      toObservable(this.previewService.dataFile),
      toObservable(this.previewService.refreshTrigger),
      toObservable(this.previewService.hasSyntaxErrors),
    ])
      .pipe(
        debounceTime(500),
        filter(([ops, data, _trigger, hasErrors]) => !!ops && !!data && !hasErrors),
        tap(() => this.loading.set(true)),
        switchMap(([ops, data]) =>
          this.editorService.getPreview(ops!, data!).pipe(
            catchError(() => of(null)),
            tap(() => this.loading.set(false)),
          ),
        ),
        takeUntilDestroyed(this.destroyRef),
      )
      .subscribe((res: any) => {
        console.log(res);
        this.previewService.preview.set(res);
      });
  }
}
