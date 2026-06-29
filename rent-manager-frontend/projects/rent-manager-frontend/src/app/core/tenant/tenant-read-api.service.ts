import { HttpClient, HttpParams } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { Tenant } from '@core/tenant/tenant-core.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class TenantReadApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/tenants');
  }

  list$(page: number, limit: number): Observable<Tenant[]> {
    const params = new HttpParams().set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<Tenant[]>>(this.baseUrl, { params })
      .pipe(map((response) => response.data));
  }

  search$(value: string, page: number, limit: number): Observable<Tenant[]> {
    const params = new HttpParams().set('value', value).set('page', page).set('limit', limit);

    return this.http
      .get<ApiResponse<Tenant[]>>(`${this.baseUrl}/search`, { params })
      .pipe(map((response) => response.data));
  }

  getById$(id: number): Observable<Tenant> {
    return this.http
      .get<ApiResponse<Tenant>>(`${this.baseUrl}/${id}`)
      .pipe(map((response) => response.data));
  }
}
