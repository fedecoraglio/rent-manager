import { computed, inject, Injectable, signal } from '@angular/core';
import { finalize, Observable, tap } from 'rxjs';

import { CreateUserRequest, UpdateUserRequest, User, UserFormValue } from '../model/user.model';
import { UserApiService } from './user-api.service';

@Injectable()
export class UserService {
  private readonly userApi = inject(UserApiService);

  private readonly ownersWritableSignal = signal<User[]>([]);
  readonly ownersSignal = computed(() => this.ownersWritableSignal());
  readonly selectedUserSignal = signal<User | null>(null);
  readonly isLoadingSignal = signal(false);
  readonly errorSignal = signal<string | null>(null);

  save$(user: UserFormValue): Observable<User> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    const request: CreateUserRequest = {
      name: user.name,
      email: user.email,
      password: user.password,
      role_id: user.role_id,
    };

    return this.userApi.create(request).pipe(finalize(() => this.isLoadingSignal.set(false)));
  }

  update$(id: number, user: UserFormValue): Observable<User> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    const request: UpdateUserRequest = {
      name: user.name,
      email: user.email,
      role_id: user.role_id,
    };

    if (user.password?.trim()) {
      request.password_hash = user.password;
    }

    return this.userApi.update(id, request).pipe(finalize(() => this.isLoadingSignal.set(false)));
  }

  get$(id: number): Observable<User> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.userApi.getById(id).pipe(
      tap((user) => this.selectedUserSignal.set(user)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }

  list$(page = 1, limit = 10): Observable<User[]> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.userApi.list(page, limit).pipe(
      tap((users) => this.ownersWritableSignal.set(users)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }

  search$(value: string, page = 1, limit = 10): Observable<User[]> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.userApi.search(value, page, limit).pipe(
      tap((users) => this.ownersWritableSignal.set(users)),
      finalize(() => this.isLoadingSignal.set(false)),
    );
  }

  delete$(id: number): Observable<void> {
    this.isLoadingSignal.set(true);
    this.errorSignal.set(null);

    return this.userApi.delete(id).pipe(finalize(() => this.isLoadingSignal.set(false)));
  }
}
