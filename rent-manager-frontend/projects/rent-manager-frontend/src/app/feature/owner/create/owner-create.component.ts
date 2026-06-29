import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';

import { MatButtonModule } from '@angular/material/button';

import { OwnerFormComponent } from '../form/owner-form.component';
import { OwnerReadService } from '@core/owner/owner-read.service';
import { TranslatePipe } from '@ngx-translate/core';
import { OwnerWriteService } from '@feature/owner/service/owner-write.service';

@Component({
  selector: 'rm-owner-create',
  standalone: true,
  imports: [MatButtonModule, TranslatePipe, OwnerFormComponent],
  templateUrl: './owner-create.component.html',
  styleUrl: './owner-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerCreateComponent {
  private readonly ownerCoreService = inject(OwnerReadService);
  private readonly ownerService = inject(OwnerWriteService);
  private readonly router = inject(Router);

  readonly create$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly isLoadingSignal = this.ownerCoreService.isLoadingSignal;
  @ViewChild('ownerForm') ownerForm!: OwnerFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/owners']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.create$
      .pipe(
        filter(() => this.ownerForm.isValid && !this.isLoadingSignal()),
        switchMap(() => this.ownerService.save$(this.ownerForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
