import { ChangeDetectionStrategy, Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { TranslatePipe } from '@ngx-translate/core';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { OwnerReadService } from '@core/owner/owner-read.service';
import { OwnerFormComponent } from '@feature/owner/form/owner-form.component';

@Component({
  selector: 'rm-user-view',
  standalone: true,
  imports: [MatButtonModule, MatIconModule, TranslatePipe, OwnerFormComponent],
  templateUrl: './owner-view.component.html',
  styleUrl: './owner-view.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerViewComponent implements OnInit {
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);
  private readonly ownerReadService = inject(OwnerReadService);

  readonly errorSignal = this.ownerReadService.errorSignal;
  readonly selectedOwnerSignal = this.ownerReadService.selectedOwnerSignal;
  readonly searchById$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly goToEdit$ = new Subject<void>();
  readonly isLoadingSignal = this.ownerReadService.isLoadingSignal;

  private ownerId!: number;

  constructor() {
    this.searchById$
      .pipe(
        switchMap(() => this.ownerReadService.get$(this.ownerId)),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/owners']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToEdit$
      .pipe(
        filter(() => !!this.ownerId),
        switchMap(() => from(this.router.navigate(['/app/owners/edit', this.ownerId]))),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  ngOnInit(): void {
    this.ownerId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.ownerId) {
      this.goToList$.next();
      return;
    }

    this.searchById$.next();
  }
}
