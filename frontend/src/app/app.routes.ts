import { Routes } from '@angular/router';
import { LandingPage } from '../landing-page/landing-page';
import { MainPage } from '../main-page/main-page';

export const routes: Routes = [
    {
        path: '',
        component: LandingPage
    },
    {
        path: 'main',
        component: MainPage
    }
];
