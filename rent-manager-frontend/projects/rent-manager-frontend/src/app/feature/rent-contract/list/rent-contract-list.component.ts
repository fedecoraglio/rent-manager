import { ChangeDetectionStrategy, Component, inject, OnInit, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { Router, RouterLink } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Subject, switchMap, tap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginatorModule, PageEvent } from '@angular/material/paginator';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { TranslatePipe } from '@ngx-translate/core';

import { PropertyReadService } from '@core/property/property-read.service';
import { RentContractReadService } from '@core/rent-contract/rent-contract-read.service';
import { RentContractListTableComponent } from './table/rent-contract-list-table.component';

@Component({
  selector: 'rm-rent-contract-list',
  standalone: true,
  imports: [
    FormsModule,
    RouterLink,
    MatButtonModule,
    MatFormFieldModule,
    MatIconModule,
    MatPaginatorModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    TranslatePipe,
    RentContractListTableComponent,
  ],
  templateUrl: './rent-contract-list.component.html',
  styleUrl: './rent-contract-list.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractListComponent implements OnInit {
  private readonly router = inject(Router);
  private readonly rentContractReadService = inject(RentContractReadService);
  private readonly propertyReadService = inject(PropertyReadService);

  readonly rentContractsSignal = this.rentContractReadService.rentContractsSignal;
  readonly propertiesSignal = this.propertyReadService.propertiesSignal;
  readonly propertyIdSignal = signal(0);
  readonly pageSignal = signal(1);
  readonly limitSignal = signal(10);
  readonly load$ = new Subject<void>();
  readonly clearFilter$ = new Subject<void>();

  constructor() {
    this.load$
      .pipe(
        switchMap(() =>
          this.rentContractReadService.list$(
            this.propertyIdSignal(),
            this.pageSignal(),
            this.limitSignal(),
          ),
        ),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.clearFilter$
      .pipe(
        tap(() => this.propertyIdSignal.set(0)),
        tap(() => this.pageSignal.set(1)),
        tap(() => this.load$.next()),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  ngOnInit(): void {
    this.load$.next();
  }

  onPropertyChange(propertyId: number): void {
    this.propertyIdSignal.set(propertyId);
    this.pageSignal.set(1);
    this.load$.next();
  }

  onPageChange(event: PageEvent): void {
    this.pageSignal.set(event.pageIndex + 1);
    this.limitSignal.set(event.pageSize);
    this.load$.next();
  }

  viewRentContract(id: number): void {
    void this.router.navigate(['/app/rent-contracts/view', id]);
  }

  editRentContract(id: number): void {
    void this.router.navigate(['/app/rent-contracts/edit', id]);
  }
}
