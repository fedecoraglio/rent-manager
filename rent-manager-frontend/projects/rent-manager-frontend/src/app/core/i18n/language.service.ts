import { inject, Injectable } from '@angular/core';
import { TranslateService } from '@ngx-translate/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

type Language = 'es' | 'en';

const LANGUAGE_KEY = 'rent_manager_language';

@Injectable({ providedIn: 'root' })
export class LanguageService {
  private readonly translateService = inject(TranslateService);

  constructor() {
    this.translateService.setFallbackLang('es').pipe(takeUntilDestroyed()).subscribe();
  }

  initialize(): void {
    const language = this.getStoredLanguage();

    this.translateService.use(language);
  }

  changeLanguage(language: Language): void {
    localStorage.setItem(LANGUAGE_KEY, language);
    this.translateService.use(language);
  }

  private getStoredLanguage(): Language {
    const language = localStorage.getItem(LANGUAGE_KEY);

    if (language === 'en' || language === 'es') {
      return language;
    }

    return 'es';
  }
}
