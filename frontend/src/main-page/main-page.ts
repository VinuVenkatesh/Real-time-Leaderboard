import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
import { ActivityModal } from '../activity-modal/activity-modal';

@Component({
  selector: 'app-main-page',
  imports: [MatButtonModule, MatDialogModule],
  templateUrl: './main-page.html',
  styleUrl: './main-page.css',
})
export class MainPage {
  private dialog = inject(MatDialog);

  activity(){
    this.dialog.open(ActivityModal,{
      width: '600px',
      height: '400px'
    });
  }
}
