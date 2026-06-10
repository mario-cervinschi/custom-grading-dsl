import { Injectable, signal } from '@angular/core';

export interface Explanation {
  original: string;
  substituted: string;
  result: string;
  description?: string;
}

export interface VariableData {
  result: any;
  explanation?: Explanation;
}

export interface Result {
  id: string;
  data: Record<string, VariableData>;
}

export interface EvaluationResponse {
  success: boolean;
  results: Result[];
}

@Injectable({
  providedIn: 'root',
})
export class PreviewService {
  operationsFile = signal<string | null>(null);
  dataFile = signal<string | null>(null);

  preview = signal<EvaluationResponse | null>(null);

  refreshTrigger = signal<number>(0);
  hasSyntaxErrors = signal<boolean>(false);

  triggerRefresh() {
    this.refreshTrigger.update(val => val + 1);
  }
}