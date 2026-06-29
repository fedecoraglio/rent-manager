import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { InflationIndexReadService } from '@core/inflation-index/inflation-index-read.service';
import { InflationIndexWriteService } from '@feature/inflation-index/service/inflation-index-write.service';
import { InflationIndexFormComponent } from '../form/inflation-index-form.component';

@Component({
  selector: 'rm-inflation-index-create',
  standalone: true,
  imports: [MatButtonModule, TranslatePipe, InflationIndexFormComponent],
  templateUrl: './inflation-index-create.component.html',
  styleUrl: './inflation-index-create.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InflationIndexCreateComponent {
  private readonly inflationIndexReadService = inject(InflationIndexReadService);
  private readonly inflationIndexWriteService = inject(InflationIndexWriteService);
  private readonly router = inject(Router);

  readonly create$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly isLoadingSignal = this.inflationIndexReadService.isLoadingSignal;

  @ViewChild('inflationIndexForm')
  inflationIndexForm!: InflationIndexFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/inflation-indexes']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.create$
      .pipe(
        filter(() => this.inflationIndexForm.isValid && !this.isLoadingSignal()),
        switchMap(() => this.inflationIndexWriteService.save$(this.inflationIndexForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
