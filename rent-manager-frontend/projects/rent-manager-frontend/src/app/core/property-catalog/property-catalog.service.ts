import { HttpErrorResponse } from '@angular/common/http';
import { computed, inject, Injectable, signal } from '@angular/core';
import { catchError, EMPTY, map, Observable, take, tap } from 'rxjs';

import { PropertyCatalogApiService } from './property-catalog-api.service';
import { PropertyStatus, PropertyType } from './property-catalog.model';

@Injectable({ providedIn: 'root' })
export class PropertyCatalogService {
  private readonly api = inject(PropertyCatalogApiService);

  private readonly propertyTypesWritable = signal<PropertyType[]>([]);
  private readonly propertyStatusesWritable = signal<PropertyStatus[]>([]);

  readonly propertyTypesSignal = computed(() => this.propertyTypesWritable());
  readonly propertyStatusesSignal = computed(() => this.propertyStatusesWritable());

  listPropertyTypes$(): Observable<PropertyType[]> {
    return this.api.listPropertyTypes$().pipe(
      take(1),
      tap((apiResponse) => this.propertyTypesWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  listPropertyStatuses$(): Observable<PropertyStatus[]> {
    return this.api.listPropertyStatuses$().pipe(
      take(1),
      tap((apiResponse) => this.propertyStatusesWritable.set(apiResponse.data)),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }
}
