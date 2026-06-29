import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { RentContractFormComponent } from '../form/rent-contract-form.component';
import { RentContractWriteService } from '../service/rent-contract-write.service';

@Component({
  selector: 'rm-rent-contract-create',
  standalone: true,
  imports: [RentContractFormComponent, MatButtonModule, TranslatePipe],
  templateUrl: './rent-contract-create.component.html',
  styleUrl: './rent-contract-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractCreateComponent {
  private readonly writeService = inject(RentContractWriteService);
  private readonly router = inject(Router);

  readonly create$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly isLoadingSignal = this.writeService.isLoadingSignal;

  @ViewChild('rentContractForm') rentContractForm!: RentContractFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/rent-contracts']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.create$
      .pipe(
        filter(() => this.rentContractForm.isValid),
        switchMap(() => this.writeService.save$(this.rentContractForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
