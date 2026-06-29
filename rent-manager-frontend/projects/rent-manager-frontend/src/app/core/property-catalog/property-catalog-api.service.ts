import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';

import { ApiResponse } from '@core/api/api.model';
import { PropertyStatus, PropertyType } from './property-catalog.model';
import { ApiUrlService } from '@core/api/api-url.service';

@Injectable({ providedIn: 'root' })
export class PropertyCatalogApiService {
  private readonly http = inject(HttpClient);
  private readonly apiUrl = inject(ApiUrlService);

  private get baseUrl(): string {
    return this.apiUrl.build('/property-catalogs');
  }

  listPropertyTypes$(): Observable<ApiResponse<PropertyType[]>> {
    return this.http
      .get<ApiResponse<PropertyType[]>>(`${this.baseUrl}/types`)
      .pipe(map((response) => response));
  }

  listPropertyStatuses$(): Observable<ApiResponse<PropertyStatus[]>> {
    return this.http
      .get<ApiResponse<PropertyStatus[]>>(`${this.baseUrl}/statuses`)
      .pipe(map((response) => response));
  }
}
