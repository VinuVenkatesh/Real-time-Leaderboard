import { Component, inject } from '@angular/core';
import { FormGroup, FormControl, Validators, ReactiveFormsModule } from '@angular/forms';
import { MatDialogRef, MatDialogModule } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-register-modal',
  standalone: true,
  imports: [
    ReactiveFormsModule,
    MatDialogModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule
  ],
  templateUrl: './register-modal.html',
  styleUrls: ['./register-modal.css']
})
export class RegisterModal {
  private dialogRef = inject(MatDialogRef<RegisterModal>);

  // Reactive form
  registerForm = new FormGroup({
    username: new FormControl('', Validators.required),
    password: new FormControl('', Validators.required)
  });

  onRegister() {
    if (this.registerForm.invalid) {
      this.registerForm.markAllAsTouched(); // show errors
      return;
    }

    // close dialog with form data
    this.dialogRef.close(this.registerForm.value);
  }
}
