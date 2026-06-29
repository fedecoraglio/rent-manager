import { ChangeDetectionStrategy, Component, inject, signal } from '@angular/core';
import { NonNullableFormBuilder, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Subject, switchMap } from 'rxjs';

import { AuthService } from '@core/auth/auth.service';
import { TranslatePipe } from '@ngx-translate/core';

@Component({
  selector: 'rm-login',
  standalone: true,
  imports: [
    TranslatePipe,
    ReactiveFormsModule,
    MatButtonModule,
    MatCardModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatProgressSpinnerModule,
  ],
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  private readonly fb = inject(NonNullableFormBuilder);
  private readonly authService = inject(AuthService);
  private readonly router = inject(Router);

  readonly doLogin$ = new Subject<void>();
  readonly loading = signal(false);
  readonly loginError = signal<string | null>(null);
  readonly hidePassword = signal(true);

  readonly form = this.fb.group({
    email: ['', [Validators.required, Validators.email]],
    password: ['', [Validators.required]],
  });

  constructor() {
    this.doLogin$
      .pipe(
        switchMap(() => this.authService.login(this.form.getRawValue())),
        takeUntilDestroyed(),
      )
      .subscribe({
        next: () => {
          this.loading.set(false);
          void this.router.navigate(['/app']);
        },
        error: () => {
          this.loading.set(false);
          this.loginError.set('login.error');
        },
      });
  }

  submit(): void {
    if (this.form.invalid || this.loading()) {
      this.form.markAllAsTouched();
      return;
    }

    this.loading.set(true);
    this.loginError.set(null);
    this.doLogin$.next();
  }
}
