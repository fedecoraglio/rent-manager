import { ChangeDetectionStrategy, Component, inject, OnInit, ViewChild } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { PropertyFormComponent } from '../form/property-form.component';
import { PropertyWriteService } from '../service/property-write.service';
import { PropertyReadService } from '@core/property/property-read.service';

@Component({
  selector: 'rm-property-edit',
  standalone: true,
  imports: [PropertyFormComponent, MatButtonModule, MatIconModule, TranslatePipe],
  templateUrl: './property-edit.component.html',
  styleUrl: './property-edit.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PropertyEditComponent implements OnInit {
  private readonly propertyWriteService = inject(PropertyWriteService);
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly selectedPropertySignal = this.propertyReadService.selectedPropertySignal;
  readonly isLoadingSignal = this.propertyReadService.isLoadingSignal;
  readonly errorSignal = this.propertyReadService.errorSignal;

  readonly goToList$ = new Subject<void>();
  readonly update$ = new Subject<void>();

  private propertyId!: number;

  @ViewChild('propertyForm') propertyForm!: PropertyFormComponent;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/properties']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.update$
      .pipe(
        filter(() => this.propertyForm.isValid),
        switchMap(() =>
          this.propertyWriteService.update$(this.propertyId, this.propertyForm.value),
        ),
        takeUntilDestroyed(),
      )
      .subscribe(() => {
        this.goToList$.next();
      });
  }

  ngOnInit(): void {
    this.propertyId = Number(this.route.snapshot.paramMap.get('id'));

    if (!this.propertyId) {
      this.goToList$.next();
      return;
    }

    this.propertyReadService.get$(this.propertyId).subscribe();
  }
}
