import { HttpErrorResponse, HttpInterceptorFn } from '@angular/common/http';
import { inject } from '@angular/core';
import { catchError, throwError } from 'rxjs';

import { NotificationService } from '@core/notification/notification.service';
import { ApiErrorResponse } from '@core/error/api-error.model';

export const apiErrorInterceptor: HttpInterceptorFn = (request, next) => {
  const notificationService = inject(NotificationService);

  return next(request).pipe(
    catchError((error: HttpErrorResponse) => {
      const apiError = error.error as Partial<ApiErrorResponse> | null;

      if (apiError?.message) {
        notificationService.showError(apiError.message);
      } else {
        notificationService.showError('Unexpected error');
      }

      return throwError(() => error);
    }),
  );
};
