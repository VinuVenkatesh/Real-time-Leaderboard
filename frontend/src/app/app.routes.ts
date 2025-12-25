import { Routes } from '@angular/router';
import { LandingPage } from '../landing-page/landing-page';
import { MainPage } from '../main-page/main-page';
import { AuthGuard } from '../guards/auth-guard';

export const routes: Routes = [
    {
        path: '',
        component: LandingPage
    },
    {
        path: 'main',
        component: MainPage,
        canActivate: [AuthGuard]
    }
];
