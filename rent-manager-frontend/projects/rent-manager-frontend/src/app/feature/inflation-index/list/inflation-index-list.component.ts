import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { from, Subject, switchMap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { InflationIndexReadService } from '@core/inflation-index/inflation-index-read.service';
import { MatProgressSpinner } from '@angular/material/progress-spinner';
import { InflationIndexListTableComponent } from '@feature/inflation-index/list/table/inflation-index-list-table.component';

@Component({
  selector: 'rm-inflation-index-list',
  standalone: true,
  imports: [
    MatButtonModule,
    MatIconModule,
    TranslatePipe,
    MatProgressSpinner,
    InflationIndexListTableComponent,
  ],
  templateUrl: './inflation-index-list.component.html',
  styleUrl: './inflation-index-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InflationIndexListComponent {
  private readonly router = inject(Router);
  private readonly inflationIndexReadService = inject(InflationIndexReadService);

  readonly inflationIndexesSignal = this.inflationIndexReadService.inflationIndexesSignal;
  readonly isLoadingSignal = this.inflationIndexReadService.isLoadingSignal;

  readonly create$ = new Subject<void>();

  constructor() {
    this.inflationIndexReadService.list$().pipe(takeUntilDestroyed()).subscribe();

    this.create$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/inflation-indexes/create']))),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  view(id: number): void {
    void this.router.navigate(['/app/inflation-indexes/view', id]);
  }

  edit(id: number): void {
    void this.router.navigate(['/app/inflation-indexes/edit', id]);
  }
}
