import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { InflationIndex } from '@core/inflation-index/inflation-index.model';
import {
  InflationIndexCreateRequest,
  InflationIndexUpdateRequest,
} from '../model/inflation-index.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class InflationIndexWriteApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/inflation-indexes');
  }

  create(request: InflationIndexCreateRequest): Observable<InflationIndex> {
    return this.http
      .post<ApiResponse<InflationIndex>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  update(id: number, request: InflationIndexUpdateRequest): Observable<InflationIndex> {
    return this.http
      .put<ApiResponse<InflationIndex>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }
}
