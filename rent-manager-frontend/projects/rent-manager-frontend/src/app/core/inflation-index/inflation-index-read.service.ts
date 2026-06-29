import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable, tap } from 'rxjs';

import { InflationIndexReadApiService } from './inflation-index-read-api.service';
import { InflationIndex } from '@core/inflation-index/inflation-index.model';

@Injectable({ providedIn: 'root' })
export class InflationIndexReadService {
  private readonly inflationIndexReadApiService = inject(InflationIndexReadApiService);

  private readonly inflationIndexesWritableSignal = signal<InflationIndex[]>([]);

  readonly inflationIndexesSignal = computed(() => this.inflationIndexesWritableSignal());
  readonly selectedInflationIndexSignal = signal<InflationIndex | null>(null);
  readonly isLoadingSignal = signal(false);
  readonly errorSignal = signal<string | null>(null);

  get$(id: number): Observable<InflationIndex> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.inflationIndexReadApiService.getById$(id).pipe(
      tap((item) => this.selectedInflationIndexSignal.set(item)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }

  list$(page = 1, limit = 20): Observable<InflationIndex[]> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.inflationIndexReadApiService.list$(page, limit).pipe(
      tap((items) => {
        console.log('items', items);
        this.inflationIndexesWritableSignal.set(items);
      }),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }
}
