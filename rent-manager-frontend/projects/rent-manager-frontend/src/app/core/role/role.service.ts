import { computed, inject, Injectable, signal } from '@angular/core';
import { RoleApiService } from '@core/role/role-api.service';
import { Role } from '@core/role/role.model';
import { catchError, EMPTY, map, Observable, take, tap } from 'rxjs';
import { HttpErrorResponse } from '@angular/common/http';

@Injectable({ providedIn: 'root' })
export class RoleService {
  private readonly roleApiService = inject(RoleApiService);

  private readonly rolesWritable = signal<Role[]>([]);

  readonly rolesSignal = computed(() => this.rolesWritable());

  readonly rolesById = computed(() => {
    return new Map(this.rolesWritable().map((role) => [role.id, role]));
  });

  listRoles$(page = 1, limit = 10): Observable<Role[]> {
    return this.roleApiService.listRoles$(page, limit).pipe(
      take(1),
      tap((apiResponse) => {
        this.rolesWritable.set(apiResponse.data);
      }),
      map((apiResponse) => apiResponse.data),
      catchError((e: HttpErrorResponse) => {
        console.error(e);
        return EMPTY;
      }),
    );
  }

  getRoleById(id: number): Role | null {
    return this.rolesById().get(id) ?? null;
  }

  getRoleNameById(id: number): string {
    return this.getRoleById(id)?.name ?? '-';
  }
}
