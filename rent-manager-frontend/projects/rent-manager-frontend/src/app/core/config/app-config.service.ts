import { HttpClient } from '@angular/common/http';
import { inject, Injectable, signal } from '@angular/core';
import { firstValueFrom } from 'rxjs';

import { AppConfig } from './app-config.model';

@Injectable({ providedIn: 'root' })
export class AppConfigService {
  private readonly http = inject(HttpClient);
  private readonly configWritable = signal<AppConfig | null>(null);

  readonly configSignal = this.configWritable.asReadonly();

  async load(): Promise<void> {
    const config = await firstValueFrom(this.http.get<AppConfig>('/assets/config.json'));

    this.configWritable.set(config);
  }

  get apiBaseUrl(): string {
    const config = this.configWritable();

    if (!config) {
      throw new Error('App config was not loaded');
    }

    return config.apiBaseUrl;
  }
}
