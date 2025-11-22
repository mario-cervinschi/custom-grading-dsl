import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { InputTextModule } from 'primeng/inputtext';
import { FloatLabelModule } from 'primeng/floatlabel';
import { FormsModule } from '@angular/forms';
import { GradesService } from './service/grades.service';
import { DividerModule } from 'primeng/divider';
import { SummaryComponent } from './pages/summary/summary.component';

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
  ],
  templateUrl: './app.html',
  styleUrl: './app.css',
})
export class App {
  studentId: string = '';
  studentData: any = null;

  constructor(private gradesService: GradesService) {}

  getGrades() {
    if (!this.studentId) {
      console.warn('no id');
      return;
    }

    console.log('id is: ', this.studentId);

    this.gradesService.checkGrades(this.studentId).subscribe({
      next: (response) => {
        this.studentData = response;
        console.log(response);
      },
      error: (error) => {
        console.error(error);
      },
    });
  }
}
