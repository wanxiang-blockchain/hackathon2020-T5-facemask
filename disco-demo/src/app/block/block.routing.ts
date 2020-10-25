import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BlockMSComponent } from './block.component';


const routes: Routes = [
  {
    path: '',
    component: BlockMSComponent
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})

export class BlockMSRoutingModule {
}
