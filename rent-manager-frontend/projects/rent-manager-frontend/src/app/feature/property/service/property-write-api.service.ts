import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { CreatePropertyRequest, UpdatePropertyRequest } from '../model/property.model';
import { Property } from '@core/property/property.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable()
export class PropertyWriteApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/properties');
  }

  create$(request: CreatePropertyRequest): Observable<Property> {
    return this.http
      .post<ApiResponse<Property>>(this.baseUrl, request)
      .pipe(map((response) => response.data));
  }

  update$(id: number, request: UpdatePropertyRequest): Observable<Property> {
    return this.http
      .put<ApiResponse<Property>>(`${this.baseUrl}/${id}`, request)
      .pipe(map((response) => response.data));
  }
}
