import { HttpErrorResponse } from '@angular/common/http';
import { computed, inject, Injectable, signal } from '@angular/core';
import { catchError, EMPTY, map, Observable, take, tap } from 'rxjs';

import { ContractCatalogApiService } from './contract-catalog-api.service';
import {
  ContractStatus,
  InterestCalculationType,
  RentAdjustmentType,
} from './contract-catalog.model';

@Injectable({ providedIn: 'root' })
export class ContractCatalogService {
  private readonly api = inject(ContractCatalogApiService);

  private readonly contractStatusesWritable = signal<ContractStatus[]>([]);
  private readonly interestCalculationTypesWritable = signal<InterestCalculationType[]>([]);
  private readonly rentAdjustmentTypesWritable = signal<RentAdjustmentType[]>([]);

  readonly contractStatusesSignal = computed(() => this.contractStatusesWritable());
  readonly interestCalculationTypesSignal = computed(() => this.interestCalculationTypesWritable());
  readonly rentAdjustmentTypesSignal = computed(() => this.rentAdjustmentTypesWritable());

  listContractStatuses$(): Observable<ContractStatus[]> {
    return this.api.listContractStatuses$().pipe(
      take(1),
      tap((apiResponse) => this.contractStatusesWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  listInterestCalculationTypes$(): Observable<InterestCalculationType[]> {
    return this.api.listInterestCalculationTypes$().pipe(
      take(1),
      tap((apiResponse) => this.interestCalculationTypesWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  listRentAdjustmentTypes$(): Observable<RentAdjustmentType[]> {
    return this.api.listRentAdjustmentTypes$().pipe(
      take(1),
      tap((apiResponse) => this.rentAdjustmentTypesWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }
}
