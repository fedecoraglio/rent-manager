import { computed, Injectable, signal } from '@angular/core';

@Injectable({ providedIn: 'root' })
export class LoadingService {
  private readonly activeRequestsWritable = signal(0);

  readonly isLoading = computed(() => this.activeRequestsWritable() > 0);

  show(): void {
    this.activeRequestsWritable.update((value) => value + 1);
  }

  hide(): void {
    this.activeRequestsWritable.update((value) => Math.max(value - 1, 0));
  }
}
