import { Component, OnInit, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { FloatLabelModule } from 'primeng/floatlabel';
import { FormsModule } from '@angular/forms';
import { GradesService } from './service/grades.service';
import { DividerModule } from 'primeng/divider';
import { SummaryComponent } from './pages/summary/summary.component';
import { ProgressBarModule } from 'primeng/progressbar';
import { MessageModule } from 'primeng/message';
import { SelectButtonModule } from 'primeng/selectbutton';

@Component({
  selector: 'app-root',
  imports: [
    RouterOutlet,
    DividerModule,
    FormsModule,
    FloatLabelModule,
    ButtonModule,
    CardModule,
    InputTextModule,
    SummaryComponent,
    ProgressBarModule,
    MessageModule,
    SelectButtonModule
  ],
  templateUrl: './app.html',
  styleUrl: './app.css',
})
export class App{
  studentIdSearch: string = '';
  studentIdTitle: any = null;
  studentData: any = null;
  showInputError = false;
  searching: boolean = false;
  errorMsg: string = "";

  languageOptions = [
    { label: 'EN', value: 'en', flagCode: 'GB' },
    { label: 'RO', value: 'ro', flagCode: 'RO' }
  ];

  selectedLanguage: string = 'en';

  constructor(private gradesService: GradesService) {}

  onReturn() {
    this.studentIdTitle = null;
    this.studentData = null;
  }

  getGrades() {
    if(this.studentIdSearch.trim() === '') {
      this.showInputError = true;
      this.errorMsg = "Student ID cannot be empty.";
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
