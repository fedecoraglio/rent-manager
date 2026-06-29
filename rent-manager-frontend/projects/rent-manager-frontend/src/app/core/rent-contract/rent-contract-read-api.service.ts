import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { RentContract } from './rent-contract.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class RentContractReadApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/rental-contracts');
  }

  list$(propertyId: number, page: number, limit: number): Observable<ApiResponse<RentContract[]>> {
    let params = new HttpParams().set('page', page).set('limit', limit);

    if (propertyId > 0) {
      params = params.set('property_id', propertyId);
    }

    return this.http
      .get<ApiResponse<RentContract[]>>(this.baseUrl, { params })
      .pipe(map((response) => response));
  }

  getById$(id: number): Observable<ApiResponse<RentContract>> {
    return this.http
      .get<ApiResponse<RentContract>>(`${this.baseUrl}/${id}`)
      .pipe(map((response) => response));
  }
}
