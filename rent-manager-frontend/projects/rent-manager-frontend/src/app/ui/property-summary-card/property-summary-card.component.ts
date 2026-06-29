import { CurrencyPipe } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { TranslatePipe } from '@ngx-translate/core';

import { PropertySummaryCardData } from './property-summary-card.model';
import { Router } from '@angular/router';
import { from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

@Component({
  selector: 'rm-property-summary-card',
  standalone: true,
  imports: [CurrencyPipe, TranslatePipe],
  templateUrl: './property-summary-card.component.html',
})
export class PropertySummaryCardComponent {
  readonly propertySummary = input.required<PropertySummaryCardData>();
  private readonly router = inject(Router);
  private readonly navigateToProperty$ = new Subject<void>();

  constructor() {
    this.navigateToProperty$
      .pipe(
        switchMap(() =>
          from(
            this.router.navigate([
              '/app/rent-contracts/view',
              this.propertySummary().rentalContractId,
            ]),
          ),
        ),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  gotToViewProperty() {
    this.navigateToProperty$.next();
  }
}
