import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { Subject, switchMap, tap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginatorModule, PageEvent } from '@angular/material/paginator';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { TranslatePipe } from '@ngx-translate/core';
import { OwnerReadService } from '@core/owner/owner-read.service';
import { OwnerListTableComponent } from '@feature/owner/list/table/owner-list-table.component';

@Component({
  selector: 'rm-owner-list',
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
    TranslatePipe,
    OwnerListTableComponent,
  ],
  templateUrl: './owner-list.component.html',
  styleUrl: './owner-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerListComponent implements OnInit {
  private readonly router = inject(Router);
  private readonly ownerCoreService = inject(OwnerReadService);

  readonly errorSignal = this.ownerCoreService.errorSignal;
  readonly isLoadingSignal = this.ownerCoreService.isLoadingSignal;
  readonly ownersSignal = this.ownerCoreService.ownersSignal;
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
            return this.ownerCoreService.search$(value, this.pageSignal(), this.limitSignal());
          }

          return this.ownerCoreService.list$(this.pageSignal(), this.limitSignal());
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

  view(id: number): void {
    void this.router.navigate(['/app/owners/view', id]);
  }

  edit(id: number): void {
    void this.router.navigate(['/app/owners/edit', id]);
  }
}
