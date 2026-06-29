import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Subject, switchMap, tap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginatorModule, PageEvent } from '@angular/material/paginator';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { TranslatePipe } from '@ngx-translate/core';
import { TenantListTableComponent } from '@feature/tenant/list/table/tenant-list-table.component';
import { TenantReadService } from '@core/tenant/tenant-read.service';

@Component({
  selector: 'rm-tenant-list',
  standalone: true,
  imports: [
    FormsModule,
    RouterLink,
    MatButtonModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatPaginatorModule,
    MatProgressSpinnerModule,
    TranslatePipe,
    TenantListTableComponent,
  ],
  templateUrl: './tenant-list.component.html',
  styleUrl: './tenant-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TenantListComponent implements OnInit {
  private readonly router = inject(Router);
  private readonly tenantReadService = inject(TenantReadService);

  readonly tenantsSignal = this.tenantReadService.tenantsSignal;
  readonly isLoadingSignal = this.tenantReadService.isLoadingSignal;
  readonly errorSignal = this.tenantReadService.errorSignal;

  readonly searchValueSignal = signal('');
  readonly pageSignal = signal(1);
  readonly limitSignal = signal(10);

  readonly load$ = new Subject<void>();
  readonly search$ = new Subject<void>();
  readonly clearSearch$ = new Subject<void>();

  constructor() {
    this.load$
      .pipe(
        switchMap(() => {
          const value = this.searchValueSignal().trim();

          if (value) {
            return this.tenantReadService.search$(value, this.pageSignal(), this.limitSignal());
          }

          return this.tenantReadService.list$(this.pageSignal(), this.limitSignal());
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
  }

  ngOnInit(): void {
    this.load$.next();
  }

  onPageChange(event: PageEvent): void {
    this.pageSignal.set(event.pageIndex + 1);
    this.limitSignal.set(event.pageSize);
    this.load$.next();
  }

  viewTenant(id: number): void {
    void this.router.navigate(['/app/tenants/view', id]);
  }

  editTenant(id: number): void {
    void this.router.navigate(['/app/tenants/edit', id]);
  }
}
