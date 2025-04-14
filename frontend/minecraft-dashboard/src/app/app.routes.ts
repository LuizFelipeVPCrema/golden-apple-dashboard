import { Routes } from '@angular/router';
import { ServerListComponent } from './pages/dashboard/server-list.component';

export const routes: Routes = [
    { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
    { path: 'dashboard', component: ServerListComponent },
];
