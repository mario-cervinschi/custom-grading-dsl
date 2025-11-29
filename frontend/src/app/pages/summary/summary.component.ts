import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { CardGridComponent } from './card_grid.component';

@Component({
  selector: 'student-summary',
  standalone: true,
  imports: [CommonModule, CardGridComponent, CardModule, ButtonModule],
  templateUrl: './summary.component.html',
})
export class SummaryComponent {
  @Input() studentData: any = null;
  @Output() onGoBack: EventEmitter<any> = new EventEmitter();

  convertStringToNumber(str: string): string | number {
    const num = +str;
    if (isNaN(num) || !isFinite(num)) {
      return str;
    } else {
      return num;
    }
  }

  getExplanationsFromStudentData(variable: string) {
    return this.studentData?.data[variable].explanation;
  }

  getResultsFromStudentData(variable: string) {
    return this.convertStringToNumber(this.studentData?.data[variable].result);
  }

  goBack() {
    this.onGoBack.emit();
  }

  getStatusColor(val: any): string {
    if (val === 'absent' || val === 'toolow' || val === false) return 'text-red-500';
    if (typeof val === 'number' && val < 5) return 'text-red-500';
    if (val === true) return 'text-green-400';
    return 'text-green-400';
  }

  formatNumber(val: any): string | number {
    if (typeof val === 'number') {
      return val.toFixed(2);
    }
    return val;
  }
}
