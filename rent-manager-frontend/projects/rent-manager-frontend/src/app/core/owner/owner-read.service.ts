import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable, tap } from 'rxjs';

import { OwnerReadApiService } from './owner-read-api.service';
import { Owner } from '@core/owner/owner.model';

@Injectable({ providedIn: 'root' })
export class OwnerReadService {
  private readonly ownerCoreApiService = inject(OwnerReadApiService);

  private readonly ownersWritableSignal = signal<Owner[]>([]);
  readonly ownersSignal = computed(() => this.ownersWritableSignal());
  readonly selectedOwnerSignal = signal<Owner | null>(null);
  readonly isLoadingSignal = signal(false);
  readonly errorSignal = signal<string | null>(null);

  get$(id: number): Observable<Owner> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.ownerCoreApiService.getById(id).pipe(
      tap((item) => this.selectedOwnerSignal.set(item)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }

  list$(page = 1, limit = 10): Observable<Owner[]> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.ownerCoreApiService.list(page, limit).pipe(
      tap((items) => this.ownersWritableSignal.set(items)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }

  search$(value: string, page = 1, limit = 10): Observable<Owner[]> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.ownerCoreApiService.search(value, page, limit).pipe(
      tap((items) => this.ownersWritableSignal.set(items)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }
}
