import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { CreateOwnerRequest, UpdateOwnerRequest } from '@feature/owner/model/owner.model';
import { Owner } from '@core/owner/owner.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class OwnerWriteApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/owners');
  }

  create(request: CreateOwnerRequest): Observable<Owner> {
    return this.http
      .post<ApiResponse<Owner>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  update(id: number, request: UpdateOwnerRequest): Observable<Owner> {
    return this.http
      .put<ApiResponse<Owner>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }
}
