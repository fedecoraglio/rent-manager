import { ChangeDetectionStrategy, Component, inject, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';

import { MatButtonModule } from '@angular/material/button';
import { TranslatePipe } from '@ngx-translate/core';

import { InflationIndexReadService } from '@core/inflation-index/inflation-index-read.service';
import { InflationIndexWriteService } from '@feature/inflation-index/service/inflation-index-write.service';
import { InflationIndexFormComponent } from '../form/inflation-index-form.component';

@Component({
  selector: 'rm-inflation-index-edit',
  standalone: true,
  imports: [MatButtonModule, TranslatePipe, InflationIndexFormComponent],
  templateUrl: './inflation-index-edit.component.html',
  styleUrl: './inflation-index-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InflationIndexEditComponent {
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly inflationIndexReadService = inject(InflationIndexReadService);
  private readonly inflationIndexWriteService = inject(InflationIndexWriteService);

  readonly inflationIndexSignal = this.inflationIndexReadService.selectedInflationIndexSignal;

  readonly isLoadingSignal = this.inflationIndexReadService.isLoadingSignal;

  readonly save$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();

  readonly inflationIndexId = Number(this.route.snapshot.paramMap.get('id'));

  @ViewChild('inflationIndexForm')
  inflationIndexForm!: InflationIndexFormComponent;

  constructor() {
    this.inflationIndexReadService.get$(this.inflationIndexId).subscribe();

    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/inflation-indexes']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.save$
      .pipe(
        filter(() => this.inflationIndexForm.isValid && !this.isLoadingSignal()),
        switchMap(() =>
          this.inflationIndexWriteService.update$(
            this.inflationIndexId,
            this.inflationIndexForm.value,
          ),
        ),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }
}
