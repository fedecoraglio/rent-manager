import {
  AfterViewInit,
  ChangeDetectionStrategy,
  Component,
  inject,
  Input,
  OnDestroy,
} from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { filter, ReplaySubject, Subject, takeUntil } from 'rxjs';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { TranslatePipe } from '@ngx-translate/core';

import { ContractCatalogService } from '@core/contract-catalog/contract-catalog.service';
import { PropertyReadService } from '@core/property/property-read.service';
import { RentContract } from '@core/rent-contract/rent-contract.model';
import { TenantReadService } from '@core/tenant/tenant-read.service';
import { RentContractFormValue } from '../model/rent-contract.model';
import { RentContractFormField } from './field/rent-contract-form-field';

@Component({
  selector: 'rm-rent-contract-form',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    TranslatePipe,
  ],
  templateUrl: './rent-contract-form.component.html',
  styleUrl: './rent-contract-form.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class RentContractFormComponent implements OnDestroy, AfterViewInit {
  private readonly destroy$ = new Subject<void>();
  private readonly propertyReadService = inject(PropertyReadService);
  private readonly tenantReadService = inject(TenantReadService);
  private readonly contractCatalogService = inject(ContractCatalogService);

  readonly propertiesSignal = this.propertyReadService.propertiesSignal;
  readonly tenantsSignal = this.tenantReadService.tenantsSignal;
  readonly contractStatusesSignal = this.contractCatalogService.contractStatusesSignal;
  readonly interestCalculationTypesSignal =
    this.contractCatalogService.interestCalculationTypesSignal;
  readonly rentAdjustmentTypesSignal = this.contractCatalogService.rentAdjustmentTypesSignal;
  readonly fields = RentContractFormField;

  readonly formGroup = new FormGroup({
    [RentContractFormField.PropertyId]: new FormControl(0, Validators.required),
    [RentContractFormField.TenantId]: new FormControl(0, Validators.required),
    [RentContractFormField.StatusId]: new FormControl(0, Validators.required),
    [RentContractFormField.InterestCalculationTypeId]: new FormControl(0, Validators.required),
    [RentContractFormField.AdjustmentTypeId]: new FormControl(0, Validators.required),
    [RentContractFormField.StartDate]: new FormControl('', Validators.required),
    [RentContractFormField.EndDate]: new FormControl('', Validators.required),
    [RentContractFormField.MonthlyAmount]: new FormControl(0, [
      Validators.required,
      Validators.min(1),
    ]),
    [RentContractFormField.DepositAmount]: new FormControl(0, Validators.min(0)),
    [RentContractFormField.Currency]: new FormControl('ARS', Validators.required),
    [RentContractFormField.DueDay]: new FormControl(10, [
      Validators.required,
      Validators.min(1),
      Validators.max(31),
    ]),
    [RentContractFormField.DailyInterestPercentage]: new FormControl(0, Validators.min(0)),
    [RentContractFormField.AdjustmentFrequencyMonths]: new FormControl(4, [
      Validators.required,
      Validators.min(1),
    ]),
    [RentContractFormField.Notes]: new FormControl(''),
  });
  readonly rentContract$ = new ReplaySubject<RentContract | null>(1);

  @Input() isReadonly = false;

  @Input() set rentContract(rentContract: RentContract | null) {
    this.rentContract$.next(rentContract);
  }

  get propertyIdField() {
    return this.formGroup.get(RentContractFormField.PropertyId);
  }

  get tenantIdField() {
    return this.formGroup.get(RentContractFormField.TenantId);
  }

  get statusIdField() {
    return this.formGroup.get(RentContractFormField.StatusId);
  }

  get interestCalculationTypeIdField() {
    return this.formGroup.get(RentContractFormField.InterestCalculationTypeId);
  }

  get adjustmentTypeIdField() {
    return this.formGroup.get(RentContractFormField.AdjustmentTypeId);
  }

  get startDateField() {
    return this.formGroup.get(RentContractFormField.StartDate);
  }

  get endDateField() {
    return this.formGroup.get(RentContractFormField.EndDate);
  }

  get monthlyAmountField() {
    return this.formGroup.get(RentContractFormField.MonthlyAmount);
  }

  get currencyField() {
    return this.formGroup.get(RentContractFormField.Currency);
  }

  get dueDayField() {
    return this.formGroup.get(RentContractFormField.DueDay);
  }

  get adjustmentFrequencyMonthsField() {
    return this.formGroup.get(RentContractFormField.AdjustmentFrequencyMonths);
  }

  get isValid() {
    this.formGroup.markAllAsTouched();
    return this.formGroup.valid;
  }

  get value(): RentContractFormValue {
    return {
      ...(this.formGroup.getRawValue() as RentContractFormValue),
    };
  }

  ngAfterViewInit(): void {
    this.rentContract$
      .pipe(filter(Boolean), takeUntil(this.destroy$))
      .subscribe((rentContract: RentContract) => {
        this.formGroup.patchValue({
          property_id: rentContract.property_id,
          tenant_id: rentContract.tenant_id,
          status_id: rentContract.status_id,
          interest_calculation_type_id: rentContract.interest_calculation_type_id,
          adjustment_type_id: rentContract.adjustment_type_id,
          start_date: rentContract.start_date,
          end_date: rentContract.end_date,
          monthly_amount: rentContract.monthly_amount,
          deposit_amount: rentContract.deposit_amount,
          currency: rentContract.currency,
          due_day: rentContract.due_day,
          daily_interest_percentage: rentContract.daily_interest_percentage,
          adjustment_frequency_months: rentContract.adjustment_frequency_months,
          notes: rentContract.notes,
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
