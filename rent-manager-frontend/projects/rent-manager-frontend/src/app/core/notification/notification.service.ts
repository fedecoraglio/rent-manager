import { computed, Injectable, signal } from '@angular/core';

export type NotificationType = 'error' | 'success' | 'info' | 'warning';

export interface AppNotification {
  type: NotificationType;
  message: string;
}

@Injectable({ providedIn: 'root' })
export class NotificationService {
  private readonly notificationWritable = signal<AppNotification | null>(null);

  readonly notificationSignal = computed(() => this.notificationWritable());

  private timerId?: ReturnType<typeof setTimeout>;

  showError(message: string): void {
    this.show({
      type: 'error',
      message,
    });
  }

  showSuccess(message: string): void {
    this.show({
      type: 'success',
      message,
    });
  }

  showInfo(message: string): void {
    this.show({
      type: 'info',
      message,
    });
  }

  clear(): void {
    if (this.timerId) {
      clearTimeout(this.timerId);
    }

    this.notificationWritable.set(null);
  }

  private show(notification: AppNotification): void {
    if (this.timerId) {
      clearTimeout(this.timerId);
    }

    this.notificationWritable.set(notification);

    this.timerId = setTimeout(() => {
      this.notificationWritable.set(null);
    }, 5000);
  }
}
