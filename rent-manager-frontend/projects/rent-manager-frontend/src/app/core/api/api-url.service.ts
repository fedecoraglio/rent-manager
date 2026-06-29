import { inject, Injectable } from '@angular/core';

import { AppConfigService } from '@core/config/app-config.service';

@Injectable({ providedIn: 'root' })
export class ApiUrlService {
  private readonly appConfig = inject(AppConfigService);

  build(path: string): string {
    const normalizedPath = path.startsWith('/') ? path : `/${path}`;

    return `${this.appConfig.apiBaseUrl}${normalizedPath}`;
  }
}
