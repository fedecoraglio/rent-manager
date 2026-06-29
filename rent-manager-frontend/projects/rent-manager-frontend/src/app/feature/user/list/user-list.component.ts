import {
  ChangeDetectionStrategy,
  Component,
  computed,
  inject,
  OnInit,
  signal,
} from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { filter, Subject, switchMap, tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginatorModule, PageEvent } from '@angular/material/paginator';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { TranslatePipe } from '@ngx-translate/core';

import { UserListTableComponent } from '@feature/user/list/table/user-list-table.component';
import { UserService } from '@feature/user/service/user.service';
import { RoleService } from '@core/role/role.service';

@Component({
  selector: 'rm-user-list',
  standalone: true,
  imports: [
    FormsModule,
    RouterLink,
    MatButtonModule,
    MatIconModule,
    MatPaginatorModule,
    MatFormFieldModule,
    MatInputModule,
    MatProgressSpinnerModule,
    UserListTableComponent,
    TranslatePipe,
  ],
  templateUrl: './user-list.component.html',
  styleUrl: './user-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class UserListComponent implements OnInit {
  private readonly router = inject(Router);
  private readonly userService = inject(UserService);
  private readonly roleService = inject(RoleService);

  readonly errorSignal = this.userService.errorSignal;
  readonly isLoadingSignal = this.userService.isLoadingSignal;
  readonly usersSignal = computed(() =>
    this.userService.ownersSignal().map((user) => {
      if (user != null) {
        user.role_name = this.roleService.getRoleById(user.role_id)?.name ?? '';
      }
      return user;
    }),
  );
  readonly searchValueSignal = signal('');
  readonly pageSignal = signal(1);
  readonly limitSignal = signal(10);

  readonly load$ = new Subject<void>();
  readonly search$ = new Subject<void>();
  readonly clearSearch$ = new Subject<void>();
  readonly delete$ = new Subject<number>();

  constructor() {
    this.load$
      .pipe(
        switchMap(() => {
          const value = this.searchValueSignal().trim();

          if (value) {
            return this.userService.search$(value, this.pageSignal(), this.limitSignal());
          }

          return this.userService.list$(this.pageSignal(), this.limitSignal());
        }),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.search$
      .pipe(
        tap(() => this.pageSignal.set(1)),
        tap(() => this.load$.next()),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.clearSearch$
      .pipe(
        tap(() => this.searchValueSignal.set('')),
        tap(() => this.pageSignal.set(1)),
        tap(() => this.load$.next()),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.delete$
      .pipe(
        filter(() => window.confirm('¿Seguro que querés eliminar este usuario?')),
        switchMap((id) => this.userService.delete$(id)),
        tap(() => this.load$.next()),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  ngOnInit(): void {
    this.load$.next();
  }

  onPageChange(event: PageEvent): void {
    this.pageSignal.set(event.pageIndex + 1);
    this.limitSignal.set(event.pageSize);
    this.load$.next();
  }

  viewUser(id: number): void {
    void this.router.navigate(['/app/users/view', id]);
  }

  editUser(id: number): void {
    void this.router.navigate(['/app/users/edit', id]);
  }
}
