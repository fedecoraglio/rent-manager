import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable } from 'rxjs';

import {
  CreateTenantRequest,
  TenantFormValue,
  UpdateTenantRequest,
} from '@feature/tenant/model/tenant.model';
import { TenantWriteApiService } from './tenant-write-api.service';
import { Tenant } from '@core/tenant/tenant-core.model';

@Injectable()
export class TenantWriteService {
  private readonly tenantApi = inject(TenantWriteApiService);
  private readonly isLoadingWritable = signal(false);
  private readonly errorWritable = signal<string | null>(null);

  readonly isLoadingSignal = computed(() => this.isLoadingWritable());
  readonly errorSignal = computed(() => this.errorWritable());

  save$(tenant: TenantFormValue): Observable<Tenant> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: CreateTenantRequest = { ...tenant };

    return this.tenantApi.create$(request).pipe(finalize(() => this.isLoadingWritable.set(false)));
  }

  update$(id: number, tenant: TenantFormValue): Observable<Tenant> {
    this.isLoadingWritable.set(true);
    this.errorWritable.set(null);

    const request: UpdateTenantRequest = { ...tenant };

    return this.tenantApi
      .update$(id, request)
      .pipe(finalize(() => this.isLoadingWritable.set(false)));
  }
}
