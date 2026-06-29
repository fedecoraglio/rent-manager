import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { CreateTenantRequest, UpdateTenantRequest } from '../model/tenant.model';
import { Tenant } from '@core/tenant/tenant-core.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable()
export class TenantWriteApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/tenants');
  }

  create$(request: CreateTenantRequest): Observable<Tenant> {
    return this.http
      .post<ApiResponse<Tenant>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  update$(id: number, request: UpdateTenantRequest): Observable<Tenant> {
    return this.http
      .put<ApiResponse<Tenant>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }
}
