import { ChangeDetectionStrategy, Component, inject, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { RentContractReadService } from '@core/rent-contract/rent-contract-read.service';
import { RentContractWriteService } from '../service/rent-contract-write.service';
import { RentContractFormComponent } from '../form/rent-contract-form.component';

@Component({
  selector: 'rm-rent-contract-edit',
  standalone: true,
  imports: [RentContractFormComponent, MatButtonModule, MatIconModule, TranslatePipe],
  templateUrl: './rent-contract-edit.component.html',
  styleUrl: './rent-contract-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractEditComponent implements OnInit {
  private readonly readService = inject(RentContractReadService);
  private readonly writeService = inject(RentContractWriteService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly selectedRentContractSignal = this.readService.selectedRentContractSignal;
  readonly isLoadingSignal = this.writeService.isLoadingSignal;
  readonly errorSignal = this.writeService.errorSignal;

  readonly goToList$ = new Subject<void>();
  readonly update$ = new Subject<void>();

  private rentContractId!: number;

  @ViewChild('rentContractForm') rentContractForm!: RentContractFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/rent-contracts']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.update$
      .pipe(
        filter(() => this.rentContractForm.isValid),
        switchMap(() =>
          this.writeService.update$(this.rentContractId, this.rentContractForm.value),
        ),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }

  ngOnInit(): void {
    this.rentContractId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.rentContractId) {
      this.goToList$.next();
      return;
    }

    this.readService.get$(this.rentContractId).subscribe();
  }
}
