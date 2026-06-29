import { ChangeDetectionStrategy, Component, computed, inject } from '@angular/core';
import { TranslatePipe } from '@ngx-translate/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

import { PropertyReadService } from '@core/property/property-read.service';
import { PropertySummaryCardComponent } from '@ui/property-summary-card/property-summary-card.component';
import { DashboardMapper } from './dashboard.mapper';

@Component({
  selector: 'rm-dashboard',
  standalone: true,
  imports: [PropertySummaryCardComponent, TranslatePipe],
  templateUrl: './dashboard.component.html',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DashboardComponent {
  private readonly propertyReadService = inject(PropertyReadService);

  readonly isLoading = this.propertyReadService.isLoadingSignal;

  readonly properties = computed(() =>
    this.propertyReadService.propertiesSummariesSignal().map(DashboardMapper.toCardData),
  );

  constructor() {
    this.propertyReadService.listSummary$().pipe(takeUntilDestroyed()).subscribe();
  }
}
