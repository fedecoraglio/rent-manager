import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { PropertyReadService } from '@core/property/property-read.service';
import { TenantReadService } from '@core/tenant/tenant-read.service';

@Component({
  selector: 'rm-rent-contract',
  standalone: true,
  imports: [RouterOutlet],
  templateUrl: './rent-contract.component.html',
  styleUrl: './rent-contract.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractComponent {
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly tenantReadService = inject(TenantReadService);

  constructor() {
    this.propertyReadService.list$().pipe(takeUntilDestroyed()).subscribe();
    this.tenantReadService.list$().pipe(takeUntilDestroyed()).subscribe();
  }
}
