import { Routes } from '@angular/router';
import { Editor } from './pages/editor/editor.component';

export const routes: Routes = [
    { path: '', redirectTo: 'editor', pathMatch: "full" },
    { path: 'editor', component: Editor },
];
