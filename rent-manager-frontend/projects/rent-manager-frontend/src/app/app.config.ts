import { ApplicationConfig, provideAppInitializer, inject } from '@angular/core';
import { provideHttpClient, withInterceptors } from '@angular/common/http';
import { provideAnimations } from '@angular/platform-browser/animations';
import { provideRouter, withComponentInputBinding } from '@angular/router';
import { provideTranslateService } from '@ngx-translate/core';
import { provideTranslateHttpLoader } from '@ngx-translate/http-loader';

import { AppConfigService } from '@core/config/app-config.service';
import { authInterceptor } from '@core/auth/auth.interceptor';
import { apiErrorInterceptor } from '@core/error/api-error.interceptor';
import { loadingInterceptor } from '@core/loading/loading.interceptor';

import { routes } from './app.routes';

export const appConfig: ApplicationConfig = {
  providers: [
    provideAnimations(),

    provideRouter(routes, withComponentInputBinding()),

    provideHttpClient(withInterceptors([authInterceptor, apiErrorInterceptor, loadingInterceptor])),

    provideTranslateService({
      loader: provideTranslateHttpLoader({
        prefix: '/assets/i18n/',
        suffix: '.json',
      }),
      fallbackLang: 'es',
      lang: 'es',
    }),

    provideAppInitializer(() => {
      const appConfigService = inject(AppConfigService);
      return appConfigService.load();
    }),
  ],
};
