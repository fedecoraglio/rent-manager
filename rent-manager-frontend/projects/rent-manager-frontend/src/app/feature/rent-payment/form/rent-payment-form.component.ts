import { AfterViewInit, ChangeDetectionStrategy, Component, Input, OnDestroy } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { filter, ReplaySubject, Subject, takeUntil } from 'rxjs';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { TranslatePipe } from '@ngx-translate/core';

import { RentPayment, RentPaymentSuggestion } from '@core/rent-payment/rent-payment.model';
import { RentPaymentFormValue } from '../model/rent-payment-model';
import { RentPaymentFormField } from './field/rent-payment-form-field';

@Component({
  selector: 'rm-rent-payment-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatInputModule,
    TranslatePipe,
  ],
  templateUrl: './rent-payment-form.component.html',
  styleUrl: './rent-payment-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentPaymentFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();

  readonly fields = RentPaymentFormField;

  readonly formGroup = new FormGroup({
    [RentPaymentFormField.RentalContractId]: new FormControl(0, Validators.required),
    [RentPaymentFormField.Period]: new FormControl('', Validators.required),
    [RentPaymentFormField.DueDate]: new FormControl('', Validators.required),
    [RentPaymentFormField.PaymentDate]: new FormControl('', Validators.required),
    [RentPaymentFormField.BaseAmount]: new FormControl(0, [Validators.required, Validators.min(1)]),
    [RentPaymentFormField.SuggestedAdjustmentPercentage]: new FormControl(0),
    [RentPaymentFormField.AppliedAdjustmentPercentage]: new FormControl(0),
    [RentPaymentFormField.TotalAmount]: new FormControl(0, [
      Validators.required,
      Validators.min(1),
    ]),
    [RentPaymentFormField.IsPaid]: new FormControl(true),
    [RentPaymentFormField.Notes]: new FormControl(''),
  });

  readonly rentPayment$ = new ReplaySubject<RentPayment | null>(1);
  readonly suggestion$ = new ReplaySubject<RentPaymentSuggestion | null>(1);

  @Input() isReadonly = false;

  @Input() set rentPayment(rentPayment: RentPayment | null) {
    this.rentPayment$.next(rentPayment);
  }

  @Input() set suggestion(suggestion: RentPaymentSuggestion | null) {
    this.suggestion$.next(suggestion);
  }

  get totalAmountField() {
    return this.formGroup.get(RentPaymentFormField.TotalAmount);
  }

  get paidAmountField() {
    return this.formGroup.get(RentPaymentFormField.PaidAmount);
  }

  get paymentDateField() {
    return this.formGroup.get(RentPaymentFormField.PaymentDate);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): RentPaymentFormValue {
    return {
      ...(this.formGroup.getRawValue() as RentPaymentFormValue),
    };
  }

  ngAfterViewInit(): void {
    this.suggestion$
      .pipe(filter(Boolean), takeUntil(this.destroy$))
      .subscribe((suggestion: RentPaymentSuggestion) => {
        this.formGroup.patchValue({
          rental_contract_id: suggestion.rental_contract_id,
          period: suggestion.period,
          due_date: suggestion.due_date,
          payment_date: suggestion.payment_date,
          base_amount: suggestion.base_amount,
          suggested_adjustment_percentage: suggestion.suggested_adjustment_percentage,
          applied_adjustment_percentage: suggestion.suggested_adjustment_percentage,
          total_amount: suggestion.suggested_total_amount,
          is_paid: true,
        });
      });

    this.rentPayment$
      .pipe(filter(Boolean), takeUntil(this.destroy$))
      .subscribe((rentPayment: RentPayment) => {
        this.formGroup.patchValue({
          rental_contract_id: rentPayment.rental_contract_id,
          period: rentPayment.period,
          due_date: rentPayment.due_date,
          payment_date: rentPayment.payment_date ?? '',
          base_amount: rentPayment.base_amount,
          suggested_adjustment_percentage: rentPayment.suggested_adjustment_percentage,
          applied_adjustment_percentage: rentPayment.applied_adjustment_percentage,
          total_amount: rentPayment.total_amount,
          is_paid: rentPayment.is_paid,
          notes: rentPayment.notes,
        });

        if (this.isReadonly) {
          this.formGroup.disable();
        }

        this.formGroup.updateValueAndValidity();
      });
  }

  ngOnDestroy(): void {
    this.destroy$.next();
    this.destroy$.complete();
  }
}
