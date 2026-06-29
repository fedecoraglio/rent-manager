import { inject, Injectable, signal } from '@angular/core';
import { finalize, Observable } from 'rxjs';

import { InflationIndex } from '@core/inflation-index/inflation-index.model';
import {
  InflationIndexCreateRequest,
  InflationIndexUpdateRequest,
} from '../model/inflation-index.model';
import { InflationIndexWriteApiService } from './inflation-index-write-api.service';

@Injectable({ providedIn: 'root' })
export class InflationIndexWriteService {
  private readonly inflationIndexWriteApiService = inject(InflationIndexWriteApiService);

  readonly isSavingSignal = signal(false);
  readonly errorSignal = signal<string | null>(null);

  save$(request: InflationIndexCreateRequest): Observable<InflationIndex> {
    this.isSavingSignal.set(true);
    this.errorSignal.set(null);

    return this.inflationIndexWriteApiService
      .create(request)
      .pipe(finalize(() => this.isSavingSignal.set(false)));
  }

  update$(id: number, request: InflationIndexUpdateRequest): Observable<InflationIndex> {
    this.isSavingSignal.set(true);
    this.errorSignal.set(null);

    return this.inflationIndexWriteApiService
      .update(id, request)
      .pipe(finalize(() => this.isSavingSignal.set(false)));
  }
}
