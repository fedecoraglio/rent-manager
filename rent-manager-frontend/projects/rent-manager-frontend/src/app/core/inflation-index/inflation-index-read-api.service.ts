import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { InflationIndex } from '@core/inflation-index/inflation-index.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class InflationIndexReadApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/inflation-indexes');
  }

  list$(page: number, limit: number): Observable<InflationIndex[]> {
    const params = new HttpParams().set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<InflationIndex[]>>(this.baseUrl, { params })
      .pipe(map((response) => response.data));
  }

  getById$(id: number): Observable<InflationIndex> {
    return this.http
      .get<ApiResponse<InflationIndex>>(`${this.baseUrl}/${id}`)
      .pipe(map((response) => response.data));
  }
}
