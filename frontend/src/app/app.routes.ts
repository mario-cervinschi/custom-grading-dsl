import { Routes } from '@angular/router';
import { Editor } from './pages/editor/editor.component';
import { InputChecker } from './pages/input-checker/input-checker.component';

export const routes: Routes = [
    { path: '', redirectTo: 'editor', pathMatch: "full" },
    { path: 'editor', component: Editor },
    // { path: 'input', component: InputChecker },
];
