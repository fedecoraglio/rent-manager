import { ChangeDetectionStrategy, Component, inject, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { filter, from, Subject, switchMap } from 'rxjs';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { OwnerFormComponent } from '../form/owner-form.component';
import { OwnerReadService } from '@core/owner/owner-read.service';
import { OwnerWriteService } from '@feature/owner/service/owner-write.service';

@Component({
  selector: 'rm-owner-edit',
  standalone: true,
  imports: [MatButtonModule, MatIconModule, TranslatePipe, OwnerFormComponent],
  templateUrl: './owner-edit.component.html',
  styleUrl: './owner-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class OwnerEditComponent implements OnInit {
  private readonly ownerCoreService = inject(OwnerReadService);
  private readonly ownerService = inject(OwnerWriteService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly searchById$ = new Subject<void>();
  readonly goToList$ = new Subject<void>();
  readonly update$ = new Subject<void>();
  readonly isLoadingSignal = this.ownerCoreService.isLoadingSignal;
  readonly errorSignal = this.ownerCoreService.errorSignal;
  readonly selectedOwnerSignal = this.ownerCoreService.selectedOwnerSignal;

  private ownerId!: number;

  @ViewChild('ownerForm') ownerForm!: OwnerFormComponent;

  constructor() {
    this.searchById$
      .pipe(
        switchMap(() => this.ownerCoreService.get$(this.ownerId)),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/owners']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.update$
      .pipe(
        filter(() => this.ownerForm.isValid && !this.isLoadingSignal()),
        switchMap(() => this.ownerService.update$(this.ownerId, this.ownerForm.value)),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
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
