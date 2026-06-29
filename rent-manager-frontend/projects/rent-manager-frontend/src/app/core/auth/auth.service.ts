import { inject, Injectable } from '@angular/core';
import { Observable, Subject, switchMap, tap } from 'rxjs';

import { AuthApiService } from './auth-api.service';
import { LoginRequest, LoginResponse } from './auth.models';
import { TokenStorageService } from './token-storage.service';
import { RoleService } from '@core/role/role.service';

@Injectable({ providedIn: 'root' })
export class AuthService {
  private readonly authApi = inject(AuthApiService);
  private readonly tokenStorage = inject(TokenStorageService);
  private readonly roleService = inject(RoleService);

  private readonly loadInitData$ = new Subject<void>();

  constructor() {
    this.loadInitData$.pipe(switchMap(() => this.roleService.listRoles$())).subscribe();
  }

  login(request: LoginRequest): Observable<LoginResponse> {
    return this.authApi.login(request).pipe(
      tap((response) => this.tokenStorage.setToken(response.token)),
      //tap(() => this.loadInitData$.next()),
    );
  }

  logout(): void {
    this.tokenStorage.clear();
  }

  isAuthenticated(): boolean {
    return !!this.tokenStorage.getToken();
  }
}
