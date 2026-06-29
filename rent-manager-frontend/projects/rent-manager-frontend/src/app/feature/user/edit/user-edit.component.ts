import { ChangeDetectionStrategy, Component, inject, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { UserFormComponent } from '../form/user-form.component';
import { UserService } from '../service/user.service';

@Component({
  selector: 'rm-user-edit',
  standalone: true,
  imports: [UserFormComponent, MatButtonModule, MatIconModule, TranslatePipe],
  templateUrl: './user-edit.component.html',
  styleUrl: './user-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserEditComponent implements OnInit {
  private readonly userService = inject(UserService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly searchUserById$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly update$ = new Subject<void>();
  readonly isLoadingSignal = this.userService.isLoadingSignal;
  readonly errorSignal = this.userService.errorSignal;
  readonly selectedUserSignal = this.userService.selectedUserSignal;

  private userId!: number;

  @ViewChild('userForm') userForm!: UserFormComponent;

  constructor() {
    this.searchUserById$
      .pipe(
        switchMap(() => this.userService.get$(this.userId)),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/users']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.update$
      .pipe(
        filter(() => this.userForm.isValid && !this.isLoadingSignal()),
        switchMap(() => this.userService.update$(this.userId, this.userForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }

  ngOnInit(): void {
    this.userId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.userId) {
      this.goToList$.next();
      return;
    }

    this.searchUserById$.next();
  }
}
