import { ChangeDetectionStrategy, Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { TranslatePipe } from '@ngx-translate/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { UserFormComponent } from '../form/user-form.component';
import { UserService } from '../service/user.service';

@Component({
  selector: 'rm-user-view',
  standalone: true,
  imports: [MatButtonModule, MatIconModule, UserFormComponent, TranslatePipe],
  templateUrl: './user-view.component.html',
  styleUrl: './user-view.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserViewComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  readonly userService = inject(UserService);

  readonly searchUserById$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly goToEdit$ = new Subject<void>();
  readonly isLoadingSignal = this.userService.isLoadingSignal;

  private userId!: number;

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

    this.goToEdit$
      .pipe(
        filter(() => !!this.userId),
        switchMap(() => from(this.router.navigate(['/app/users/edit', this.userId]))),
        takeUntilDestroyed(),
      )
      .subscribe();
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
