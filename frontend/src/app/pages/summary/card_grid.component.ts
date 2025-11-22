import { CommonModule } from '@angular/common';
import { Component, Input } from '@angular/core';

@Component({
  selector: 'card-grid',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div class="bg-[#111] border border-gray-800 p-2 rounded">
      <div class="flex justify-between text-[10px] text-gray-500 mb-1">
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
  `,
})
export class CardGridComponent {
  @Input() eText: string = '';
  @Input() rText: string = '';

  @Input() eValue: number | string | null = null;
  @Input() rValue: number | string | null = null;

  @Input() finalText: string = '';
  @Input() finalValue: string | number | null = null;
  @Input() finalClasses: string = 'text-yellow-600';

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
