import { HttpErrorResponse } from '@angular/common/http';
import { computed, inject, Injectable, signal } from '@angular/core';
import { catchError, EMPTY, map, Observable, take, tap } from 'rxjs';

import { RentContractReadApiService } from '@core/rent-contract/rent-contract-read-api.service';
import { RentContract } from '@core/rent-contract/rent-contract.model';

@Injectable({ providedIn: 'root' })
export class RentContractReadService {
  private readonly api = inject(RentContractReadApiService);

  private readonly rentContractsWritable = signal<RentContract[]>([]);
  private readonly selectedRentContractWritable = signal<RentContract | null>(null);

  readonly rentContractsSignal = computed(() => this.rentContractsWritable());
  readonly selectedRentContractSignal = computed(() => this.selectedRentContractWritable());

  list$(propertyId = 0, page = 1, limit = 10): Observable<RentContract[]> {
    return this.api.list$(propertyId, page, limit).pipe(
      take(1),
      tap((apiResponse) => this.rentContractsWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  get$(id: number): Observable<RentContract> {
    return this.api.getById$(id).pipe(
      take(1),
      tap((apiResponse) => this.selectedRentContractWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }
}
