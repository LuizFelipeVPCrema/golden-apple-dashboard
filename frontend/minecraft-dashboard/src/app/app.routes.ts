import { Routes } from '@angular/router';
import { ServerListComponent } from './pages/dashboard/server-list.component';
import { ServerCreatorComponent } from './pages/server-creator/server-creator.component';
import { ServerDetailsComponent } from './pages/server-details/server-details.component';

export const routes: Routes = [
    { path: '', redirectTo: 'dashboard', pathMatch: 'full' },
    { path: 'dashboard', component: ServerListComponent },
    { path: 'create-server', component: ServerCreatorComponent},
    { path: 'server/:id', component: ServerDetailsComponent }
];
