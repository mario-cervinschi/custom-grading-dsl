import { Injectable } from '@angular/core';
import { ApiClientService } from './api-client.service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class GradesService {
  constructor(private apiClient: ApiClientService) {}

  checkGrades(studentId: string): Observable<any> {
    const payload = {
      student_id: studentId
    };

    return this.apiClient.post('/grade', payload);
  }
}
