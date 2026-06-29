import { ChangeDetectionStrategy, Component, inject, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { filter, from, Subject, switchMap } from 'rxjs';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { TranslatePipe } from '@ngx-translate/core';

import { PropertyFormComponent } from '../form/property-form.component';
import { PropertyReadService } from '@core/property/property-read.service';

@Component({
  selector: 'rm-property-view',
  standalone: true,
  imports: [PropertyFormComponent, MatButtonModule, MatIconModule, TranslatePipe],
  templateUrl: './property-view.component.html',
  styleUrl: './property-view.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class PropertyViewComponent implements OnInit {
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly route = inject(ActivatedRoute);
  private readonly router = inject(Router);

  readonly selectedPropertySignal = this.propertyReadService.selectedPropertySignal;
  readonly errorSignal = this.propertyReadService.errorSignal;

  readonly goToList$ = new Subject<void>();
  readonly goToEdit$ = new Subject<void>();

  private propertyId!: number;

  constructor() {
    this.goToList$
      .pipe(
        switchMap(() => from(this.router.navigate(['/app/properties']))),
        takeUntilDestroyed(),
      )
      .subscribe();

    this.goToEdit$
      .pipe(
        filter(() => !!this.propertyId),
        switchMap(() => from(this.router.navigate(['/app/properties/edit', this.propertyId]))),
        takeUntilDestroyed(),
      )
      .subscribe();
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
