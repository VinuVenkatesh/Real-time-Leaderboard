import { Component, inject } from '@angular/core';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { Activity } from '../models/activity';
import {MatDatepickerModule} from '@angular/material/datepicker';
import { MAT_MOMENT_DATE_ADAPTER_OPTIONS } from '@angular/material-moment-adapter';
import { provideMomentDateAdapter } from '@angular/material-moment-adapter';
import {MatSelectModule} from '@angular/material/select';
@Component({
  selector: 'app-activity-modal',
  providers: [
    provideMomentDateAdapter(),
    {
      provide: MAT_MOMENT_DATE_ADAPTER_OPTIONS, 
      useValue: { useUtc: true }
    }
    ],
  imports: [    
    ReactiveFormsModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatDatepickerModule,
    MatSelectModule
  ],
  templateUrl: './activity-modal.html',
  styleUrl: './activity-modal.css',
})
export class ActivityModal {
  private dialogRef = inject(MatDialogRef<ActivityModal>);
  private activityModel: Activity = new Activity();
  
  activityForm = new FormGroup({
    activityName: new FormControl(this.activityModel.activityName, Validators.required),
    activityResult: new FormControl(this.activityModel.activityResult, Validators.required),
    activityDate: new FormControl<moment.Moment | null>(null, Validators.required),
  });
  onSubmit() {
    if (this.activityForm.invalid) {
      this.activityForm.markAllAsTouched(); // show errors
      return;
    }
    const formValue = this.activityForm.value;

    const activity : Activity = {
      activityName: formValue.activityName!,
      activityResult: formValue.activityResult!,
      activityDate: formValue.activityDate ? formValue.activityDate.format('YYYY-MM-DD') : ''
    }


    console.log(activity);
    this.dialogRef.close(); // close dialog with form data
  }
}
