import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, map, Observable, tap } from 'rxjs';

import { PropertyReadApiService } from './property-read-api.service';
import { Property, PropertySummary } from '@core/property/property.model';

@Injectable({ providedIn: 'root' })
export class PropertyReadService {
  private readonly propertyReadApiService = inject(PropertyReadApiService);

  private readonly propertiesWritable = signal<Property[]>([]);
  private readonly propertiesSummariesWritable = signal<PropertySummary[]>([]);
  private readonly selectedPropertyWritable = signal<Property | null>(null);
  private readonly isLoadingWritable = signal(false);
  private readonly errorWritable = signal<string | null>(null);

  readonly propertiesSignal = computed(() => this.propertiesWritable());
  readonly propertiesSummariesSignal = computed(() => this.propertiesSummariesWritable());
  readonly selectedPropertySignal = computed(() => this.selectedPropertyWritable());
  readonly isLoadingSignal = computed(() => this.isLoadingWritable());
  readonly errorSignal = computed(() => this.errorWritable());

  get$(id: number): Observable<Property> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.propertyReadApiService.getById$(id).pipe(
      tap((property) => this.selectedPropertyWritable.set(property)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }

  list$(page = 1, limit = 10): Observable<Property[]> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.propertyReadApiService.list$(page, limit).pipe(
      tap((properties) => this.propertiesWritable.set(properties)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }

  listSummary$(page = 1, limit = 10): Observable<PropertySummary[]> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.propertyReadApiService.listSummaries$(page, limit).pipe(
      map((response) => response.data),
      tap((properties) => this.propertiesSummariesWritable.set(properties)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }

  search$(value: string, page = 1, limit = 10): Observable<Property[]> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.propertyReadApiService.search$(value, page, limit).pipe(
      tap((properties) => this.propertiesWritable.set(properties)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }
}
