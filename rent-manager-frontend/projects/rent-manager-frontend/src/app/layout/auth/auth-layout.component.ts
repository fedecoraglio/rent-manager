import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';

import { LanguageService } from '@core/i18n/language.service';

@Component({
  selector: 'rm-auth-layout',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './auth-layout.component.html',
  styleUrl: './auth-layout.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AuthLayoutComponent {
  private readonly languageService = inject(LanguageService);

  changeLanguage(language: 'es' | 'en'): void {
    this.languageService.changeLanguage(language);
  }
}
