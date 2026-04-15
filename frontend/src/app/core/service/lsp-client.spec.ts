import { TestBed } from '@angular/core/testing';

import { LspClient } from './lsp-client';

describe('LspClient', () => {
  let service: LspClient;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(LspClient);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
