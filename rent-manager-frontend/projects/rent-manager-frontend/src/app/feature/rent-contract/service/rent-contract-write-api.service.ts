import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { RentContract } from '@core/rent-contract/rent-contract.model';
import { CreateRentContractRequest, UpdateRentContractRequest } from '../model/rent-contract.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable()
export class RentContractWriteApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/rental-contracts');
  }

  create$(request: CreateRentContractRequest): Observable<RentContract> {
    return this.http
      .post<ApiResponse<RentContract>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  update$(id: number, request: UpdateRentContractRequest): Observable<RentContract> {
    return this.http
      .put<ApiResponse<RentContract>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }
}
