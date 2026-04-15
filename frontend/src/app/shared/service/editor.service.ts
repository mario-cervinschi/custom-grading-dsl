import { inject, Injectable } from '@angular/core';
import { ApiClientService } from './api-client.service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class EditorService {
  private readonly apiClient = inject(ApiClientService);

  createOperationsFile(fileName: string): Observable<any> {
    const payload = {
      file_name: fileName,
    };

    return this.apiClient.post('/operations/new', payload);
  }

  createDataFile(fileName: string): Observable<any> {
    const payload = {
      file_name: fileName,
    };

    return this.apiClient.post('/data/new', payload);
  }

  getOperationsFiles(): Observable<any> {
    return this.apiClient.get('/operations/list');
  }

  getDataFiles(): Observable<any> {
    return this.apiClient.get('/data/list');
  }

  readOperationsFile(name: string) {
    const payload = {
      name,
    };

    return this.apiClient.post('/operations/read', payload);
  }

  readDataFile(name: string) {
    const payload = {
      name,
    };

    return this.apiClient.post('/data/read', payload);
  }

  saveFileContent(type: 'operations' | 'data', name: string, content: string) {
    const payload = {
      type,
      name,
      content,
    };

    return this.apiClient.post('/file/save', payload);
  }

  getPreview(operationsFile: string, dataFile: string) {
    const payload = {
      operations_file: operationsFile,
      data_file: dataFile,
    }

    return this.apiClient.post('/evaluate/all', payload);
  }
}
