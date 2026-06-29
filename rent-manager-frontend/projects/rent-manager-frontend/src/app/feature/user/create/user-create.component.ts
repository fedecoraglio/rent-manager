import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

import { MatButtonModule } from '@angular/material/button';

import { UserFormComponent } from '../form/user-form.component';
import { UserService } from '../service/user.service';
import { TranslatePipe } from '@ngx-translate/core';

@Component({
  selector: 'rm-user-create',
  standalone: true,
  imports: [UserFormComponent, MatButtonModule, TranslatePipe],
  templateUrl: './user-create.component.html',
  styleUrl: './user-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserCreateComponent {
  private readonly userService = inject(UserService);
  private readonly router = inject(Router);

  readonly create$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly isLoadingSignal = this.userService.isLoadingSignal;
  @ViewChild('userForm') userForm!: UserFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/users']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.create$
      .pipe(
        filter(() => this.userForm.isValid && !this.isLoadingSignal()),
        switchMap(() => this.userService.save$(this.userForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
