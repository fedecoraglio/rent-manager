import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable, tap } from 'rxjs';

import { TenantReadApiService } from '@core/tenant/tenant-read-api.service';
import { Tenant } from '@core/tenant/tenant-core.model';

@Injectable({ providedIn: 'root' })
export class TenantReadService {
  private readonly tenantCoreApiService = inject(TenantReadApiService);

  private readonly tenantsWritable = signal<Tenant[]>([]);
  private readonly selectedTenantWritable = signal<Tenant | null>(null);
  private readonly isLoadingWritable = signal(false);
  private readonly errorWritable = signal<string | null>(null);

  readonly tenantsSignal = computed(() => this.tenantsWritable());
  readonly selectedTenantSignal = computed(() => this.selectedTenantWritable());
  readonly isLoadingSignal = computed(() => this.isLoadingWritable());
  readonly errorSignal = computed(() => this.errorWritable());

  get$(id: number): Observable<Tenant> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.tenantCoreApiService.getById$(id).pipe(
      tap((tenant) => this.selectedTenantWritable.set(tenant)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }

  list$(page = 1, limit = 10): Observable<Tenant[]> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.tenantCoreApiService.list$(page, limit).pipe(
      tap((tenants) => this.tenantsWritable.set(tenants)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }

  search$(value: string, page = 1, limit = 10): Observable<Tenant[]> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    return this.tenantCoreApiService.search$(value, page, limit).pipe(
      tap((tenants) => this.tenantsWritable.set(tenants)),
      finalize(() => this.isLoadingWritable.set(false)),
    );
  }
}
