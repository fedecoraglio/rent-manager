import { AfterViewInit, ChangeDetectionStrategy, Component, Input, OnDestroy } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { filter, ReplaySubject, Subject, takeUntil } from 'rxjs';

import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { TranslatePipe } from '@ngx-translate/core';

import { InflationIndex } from '@core/inflation-index/inflation-index.model';
import { InflationIndexFormValue } from '@feature/inflation-index/model/inflation-index.model';
import { InflationIndexFormField } from './field/inflation-index-form-field';

@Component({
  selector: 'rm-inflation-index-form',
  standalone: true,
  imports: [ReactiveFormsModule, MatFormFieldModule, MatInputModule, TranslatePipe],
  templateUrl: './inflation-index-form.component.html',
  styleUrl: './inflation-index-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class InflationIndexFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();

  readonly fields = InflationIndexFormField;

  readonly formGroup = new FormGroup({
    [InflationIndexFormField.Period]: new FormControl('', Validators.required),
    [InflationIndexFormField.Percentage]: new FormControl(0, [
      Validators.required,
      Validators.min(0),
    ]),
    [InflationIndexFormField.Source]: new FormControl(''),
    [InflationIndexFormField.Notes]: new FormControl(''),
  });

  readonly inflationIndex$ = new ReplaySubject<InflationIndex | null>(1);

  @Input() isReadonly = false;

  @Input() set inflationIndex(inflationIndex: InflationIndex | null) {
    this.inflationIndex$.next(inflationIndex);
  }

  get periodField() {
    return this.formGroup.get(InflationIndexFormField.Period);
  }

  get percentageField() {
    return this.formGroup.get(InflationIndexFormField.Percentage);
  }

  get sourceField() {
    return this.formGroup.get(InflationIndexFormField.Source);
  }

  get notesField() {
    return this.formGroup.get(InflationIndexFormField.Notes);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): InflationIndexFormValue {
    return {
      ...(this.formGroup.getRawValue() as InflationIndexFormValue),
    };
  }

  ngAfterViewInit() {
    this.inflationIndex$
      .pipe(filter(Boolean), takeUntil(this.destroy$))
      .subscribe((inflationIndex: InflationIndex) => {
        this.formGroup.patchValue({
          period: inflationIndex.period,
          percentage: inflationIndex.percentage,
          source: inflationIndex.source,
          notes: inflationIndex.notes,
        });

        if (this.isReadonly) {
          this.formGroup.disable();
        }

        this.formGroup.updateValueAndValidity();
      });
  }

  ngOnDestroy() {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
