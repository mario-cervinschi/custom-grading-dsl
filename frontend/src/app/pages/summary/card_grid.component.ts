import { CommonModule } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { BadgeModule } from 'primeng/badge';
import { OverlayBadgeModule } from 'primeng/overlaybadge';

@Component({
  selector: 'card-grid',
  standalone: true,
  imports: [CommonModule, BadgeModule, OverlayBadgeModule],
  template: `
    <div class="relative w-full">
      <p-overlaybadge
        value="i"
        severity="secondary"
        badgeSize="small"
        styleClass="cursor-pointer"
        [badgeDisabled]="!explanations"
        (click)="toggleExplanation()"
      >
        <div
          class="bg-[#111] border border-gray-800 p-2 transition-all duration-300 relative cursor-default"
          (click)="$event.stopPropagation()"
          [ngClass]="isExpanded && explanations ? 'rounded-t border-b-transparent' : 'rounded'"
        >
          <div class="flex justify-between text-[10px] text-gray-500 mb-1 gap-4">
            <span>{{ eText }}: {{ formatValue(eValue) }}</span>
            <span>{{ rText }}: {{ formatValue(rValue) }}</span>
          </div>

          <div class="border-t border-gray-800 mt-1 pt-1 flex justify-between items-center">
            <span class="text-xs font-bold" [ngClass]="finalClasses">{{ finalText }}</span>
            <span class="font-bold" [ngClass]="getStatusColor(finalValue)">
              {{ formatValue(finalValue) }}
            </span>
          </div>
        </div>
      </p-overlaybadge>

      <div class="slide-wrapper" [class.expanded]="isExpanded && explanations">
        <div class="slide-inner">
          <div
            class="bg-[#1a1a1a] border border-gray-800 border-t-0 p-2 rounded-b text-[10px] text-gray-400 shadow-inner -mt-[1px]"
          >
            <div class="font-semibold text-gray-500 mb-1 border-b border-gray-700 pb-1">
              Detalii calcule:
            </div>

            @if(explanations){
            <ng-container>
              <div class="flex flex-col justify-between py-0.5 hover:bg-white/5 px-1 rounded">
                <span class="text-gray-300 font-mono">{{ explanations["original"] }}</span>
                <span class="text-gray-300 font-mono">{{ explanations["substituted"] }}</span>
                <span class="text-gray-300 font-mono">{{ explanations["result"] }}</span>
              </div>
            </ng-container>
            } @else {
            <div class="text-gray-600 italic">Nu exista detalii.</div>
            }
          </div>
        </div>
      </div>
    </div>
  `,
  styles: [
    `
      .slide-wrapper {
        display: grid;
        grid-template-rows: 0fr;
        transition: grid-template-rows 0.3s ease-out;
      }

      .slide-wrapper.expanded {
        grid-template-rows: 1fr;
      }

      .slide-inner {
        overflow: hidden;
        min-height: 0;
      }
    `,
  ],
})
export class CardGridComponent implements OnInit {
  @Input() eText: string = '';
  @Input() rText: string = '';
  @Input() eValue: number | string | null = null;
  @Input() rValue: number | string | null = null;
  @Input() finalText: string = '';
  @Input() finalValue: string | number | null = null;
  @Input() explanations: { [key: string]: any } | null = null;
  @Input() finalClasses: string = 'text-yellow-600';

  isExpanded: boolean = false;

  ngOnInit(): void {}

  toggleExplanation() {
    if (this.explanations) {
      this.isExpanded = !this.isExpanded;
    }
  }

  formatValue(val: any): string {
    if (val === null || val === undefined || val === '') {
      return '-';
    }
    if (typeof val === 'number') {
      return val.toFixed(2);
    }
    return val.toString();
  }

  getStatusColor(val: any): string {
    if (val === 'absent' || val === 'toolow' || val === false) return 'text-red-500';
    if (val === null || val === undefined) return 'text-gray-600';
    if (typeof val === 'number' && val < 5) return 'text-red-500';
    return 'text-green-400';
  }
}
