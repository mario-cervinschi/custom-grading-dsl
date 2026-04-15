import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { DividerModule } from 'primeng/divider';
import { FloatLabelModule } from 'primeng/floatlabel';
import { InputTextModule } from 'primeng/inputtext';
import { SummaryComponent } from '../summary/summary.component';
import { ProgressBarModule } from 'primeng/progressbar';
import { MessageModule } from 'primeng/message';
import { SelectButtonModule } from 'primeng/selectbutton';
import { GradesService } from '../../shared/service/grades.service';

@Component({
  selector: 'app-input-checker',
  imports: [
    DividerModule,
    FormsModule,
    FloatLabelModule,
    ButtonModule,
    CardModule,
    InputTextModule,
    SummaryComponent,
    ProgressBarModule,
    MessageModule,
    SelectButtonModule,
  ],
  templateUrl: './input-checker.component.html',
  styleUrl: './input-checker.component.css',
})
export class InputChecker {
  studentIdSearch: string = '';
  studentIdTitle: any = null;
  studentData: any = null;
  showInputError = false;
  searching: boolean = false;
  errorMsg: string = '';

  languageOptions = [
    { label: 'EN', value: 'en', flagCode: 'GB' },
    { label: 'RO', value: 'ro', flagCode: 'RO' },
  ];

  selectedLanguage: string = 'en';

  constructor(private gradesService: GradesService) {}

  onReturn() {
    this.studentIdTitle = null;
    this.studentData = null;
  }

  getGrades() {
    if (this.studentIdSearch.trim() === '') {
      this.showInputError = true;
      this.errorMsg = 'Student ID cannot be empty.';
      return;
    }

    this.gradesService.checkGrades(this.studentIdSearch).subscribe({
      next: (response) => {
        this.studentIdTitle = this.studentIdSearch;
        this.studentData = response;
        console.log(response);
      },
      error: (error) => {
        this.showInputError = true;
        this.errorMsg = error.error.error;
      },
    });
  }
}
