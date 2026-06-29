import { HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';

import { TokenStorageService } from './token-storage.service';

export const authInterceptor: HttpInterceptorFn = (request, next) => {
  const tokenStorage = inject(TokenStorageService);
  const token = tokenStorage.getToken();

  if (!token) {
    return next(request);
  }

  return next(
    request.clone({
      setHeaders: {
        Authorization: `Bearer ${token}`,
      },
    }),
  );
};
