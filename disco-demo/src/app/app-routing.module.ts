import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';


const routes: Routes = [
  {
    path: '',
    loadChildren: () => import('./pages/authentication/login/login.module').then((m) => m.LoginModule),
  },
  {
    path: 'register',
    loadChildren: () => import('./pages/authentication/register/register.module').then(m => m.RegisterModule),
  },
  {
    path: '',
    component: AdminComponent,
    children: [
      {
        path: '',
        loadChildren: () => import('./pages/dashboard/dashboard.module').then(m => m.DashboardModule),
        pathMatch: 'full'
      },
      {
        path: 'apps/todo/:filter',
        loadChildren: () => import('./todo/todo.module').then(m => m.TodoModule)
      },
      { path: 'tables', loadChildren: () => import('./tables/tables.module').then(m => m.TablesModule) },
      { path: 'forms', loadChildren: () => import('./forms/forms.module').then(m => m.FormModule) },
      {
        path: 'materials',
        loadChildren: () => import('./materials/materials.module').then(m => m.MaterialsModule)
      },
      { path: 'pages', loadChildren: () => import('./pages/pages.module').then(m => m.PagesModule) },
      {
        path: 'analysis',
        loadChildren: () => import('./analysis/analysis.module').then(m => m.AnalysisModule)
      },
      {
        path: 'page-layouts',
        loadChildren: () => import('./page-layouts/page-layouts.module').then(m => m.PageLayoutsModule)
      },
      {
        path: 'kgms',
        loadChildren: () => import('./kgms/kgms.module').then(m => m.KGMSModule),
      },
      {
        path: 'block',
        loadChildren: () => import('./block/block.module').then(m => m.BlockMSModule),
      }
    ]
  },
  { path: 'home', loadChildren: () => import('./home/home.module').then(m => m.HomeModule) },
  {
    path: 'apps/navigation',
    loadChildren: () => import('./navigation/navigation.module').then(m => m.NavigationModule),
  },
  {
    path: 'mail', loadChildren: () => import('./mail/mail.module').then(m => m.MailModule),
  },
  {
    path: 'apps/chats',
    loadChildren: () => import('./chats/chats.module').then(m => m.ChatsModule),
  },
  {
    path: 'apm',
    loadChildren: () => import('./apm/apm.module').then(m => m.ApmModule),
    // canActivate: [AuthGuard]
  },
  {
    path: 'shield',
    loadChildren: () => import('./shield/shield.module').then(m => m.ShieldModule),
    // canActivate: [AuthGuard]
  },

];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule { }
