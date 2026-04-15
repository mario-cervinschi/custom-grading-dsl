import { Injectable, OnDestroy, signal } from '@angular/core';
import { environment } from '../../../environments/environment';

export interface CompletionItem {
  label: string;
  insertText?: string;
  kind?: number;
  detail?: string;
  sortText?: string;
  documentation?: string | { kind: string; value: string };
}

interface JsonRpcMessage {
  jsonrpc: '2.0';
  id?: number;
  method?: string;
  params?: any;
  result?: any;
  error?: any;
}

type PendingRequest = {
  resolve: (result: any) => void;
  reject: (err: any) => void;
};

@Injectable({
  providedIn: 'root',
})
export class LspService implements OnDestroy {
  private readonly ws = signal<WebSocket | null>(null);
  private readonly msgId = signal(1);
  private readonly pending = signal(new Map<number, PendingRequest>());
  private readonly documentVersion = signal(1);
  private readonly initialized = signal(false);

  private readonly wsUrl = signal(environment.lspWsUrl);
  private readonly languageId = signal(environment.lspLanguageId);

  private readonly documentUri: string;

  constructor() {
    this.documentUri = `file:///browser/document.${this.languageId()}`;
  }

  // init websocket connection
  connect(): Promise<void> {
    return new Promise((resolve, reject) => {
      if (this.ws() && this.ws()?.readyState === WebSocket.OPEN) {
        resolve();
        return;
      }

      this.ws.set(new WebSocket(this.wsUrl()));

      this.ws()!.onopen = async () => {
        console.log('Websocket connected');
        try {
          await this.initialize();
          this.notify('initialized', {});
          this.initialized.set(true);
          console.log('Successfully initialized connection with suggestion engine');
          resolve();
        } catch (e) {
          reject(e);
        }
      };

      this.ws()!.onmessage = (event) => this.handleMessage(event.data);

      this.ws()!.onerror = (err) => {
        console.error('Websocket error:', err);
        reject(err);
      };

      this.ws()!.onclose = () => {
        console.log('Websocket closed');
        this.initialized.set(false);
      };
    });
  }

  openDocument(text: string): void {
    if (!this.initialized) return;
    this.documentVersion.set(1);
    this.notify('textDocument/didOpen', {
      textDocument: {
        uri: this.documentUri,
        languageId: this.languageId(),
        version: this.documentVersion(),
        text,
      },
    });
  }

  updateDocument(text: string): void {
    if (!this.initialized()) return;
    this.documentVersion.update((val) => val + 1);
    this.notify('textDocument/didChange', {
      textDocument: { uri: this.documentUri, version: this.documentVersion() },
      contentChanges: [{ text }],
    });
  }

  async getCompletions(line: number, character: number): Promise<CompletionItem[]> {
    if (!this.initialized()) return [];

    try {
      const result = await this.request('textDocument/completion', {
        textDocument: { uri: this.documentUri },
        position: { line, character },
      });

      if (!result) return [];
      const items: any[] = Array.isArray(result) ? result : (result.items ?? []);
      return items as CompletionItem[];
    } catch (e) {
      console.error('Error on completing...:', e);
      return [];
    }
  }

  private initialize(): Promise<any> {
    return this.request('initialize', {
      processId: null,
      clientInfo: { name: 'angular-lsp-client', version: '1.0.0' },
      rootUri: null,
      capabilities: {
        textDocument: {
          completion: {
            completionItem: {
              snippetSupport: true,
              documentationFormat: ['markdown', 'plaintext'],
            },
          },
          synchronization: { didSave: false, willSave: false },
        },
        workspace: {},
      },
    });
  }

  private request(method: string, params: any): Promise<any> {
    const id = this.msgId();
    this.msgId.update((v) => v + 1);
    const message: JsonRpcMessage = { jsonrpc: '2.0', id, method, params };

    return new Promise((resolve, reject) => {
      this.pending.update((currMap) => {
        const newMap = new Map(currMap);
        newMap.set(id, { resolve, reject });
        return newMap;
      });
      this.send(message);
    });
  }

  private notify(method: string, params: any): void {
    this.send({ jsonrpc: '2.0', method, params });
  }

  private send(message: JsonRpcMessage): void {
    if (!this.ws() || this.ws()?.readyState !== WebSocket.OPEN) return;
    this.ws()?.send(JSON.stringify(message));
  }

  private handleMessage(data: string): void {
    try {
      const jsonStart = data.indexOf('{');
      const jsonStr = jsonStart >= 0 ? data.substring(jsonStart) : data;
      const message: JsonRpcMessage = JSON.parse(jsonStr);

      if (message.id !== undefined && this.pending().has(message.id)) {
        const handler = this.pending().get(message.id)!;
        this.pending.update((currentMap) => {
          const newMap = new Map(currentMap); 
          newMap.delete(message.id!);
          return newMap;
        });
        message.error ? handler.reject(message.error) : handler.resolve(message.result);
      }
    } catch (e) {
      console.warn('invalid message:', data);
    }
  }

  ngOnDestroy(): void {
    this.ws()?.close();
  }
}
