import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';

import {
  ContractStatus,
  InterestCalculationType,
  RentAdjustmentType,
} from './contract-catalog.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class ContractCatalogApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/contract-catalogs');
  }

  listContractStatuses$(): Observable<ApiResponse<ContractStatus[]>> {
    return this.http
      .get<ApiResponse<ContractStatus[]>>(`${this.baseUrl}/statuses`)
      .pipe(map((response) => response));
  }

  listInterestCalculationTypes$(): Observable<ApiResponse<InterestCalculationType[]>> {
    return this.http
      .get<ApiResponse<InterestCalculationType[]>>(`${this.baseUrl}/interest-calculation-types`)
      .pipe(map((response) => response));
  }

  listRentAdjustmentTypes$(): Observable<ApiResponse<RentAdjustmentType[]>> {
    return this.http
      .get<ApiResponse<RentAdjustmentType[]>>(`${this.baseUrl}/rent-adjustment-types`)
      .pipe(map((response) => response));
  }
}
