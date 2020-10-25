import { NgModule,Component, OnInit, Input, ViewChild, Type, ComponentFactoryResolver, AfterViewInit, ReflectiveInjector } from '@angular/core';
import { DragDropModule } from '@angular/cdk/drag-drop';
import { CalendarModule, DateAdapter } from 'angular-calendar';
import { adapterFactory } from 'angular-calendar/date-adapters/date-fns';
import { SharedModule } from '../../@stbui/shared/shared.module';

import { BlockMSRoutingModule } from './block.routing';

import { AgGridModule } from 'ag-grid-angular';
import { BlockMSComponent } from './block.component';
import { AddNodeStepComponent } from './step-addnode/step-addnode.component';
import { OnBlockStepComponent } from './step-onblock/step-onblock.component';
import { GraphStepComponent }   from './step-graph/step-graph.component';
import { AnalysisStepComponent } from './step-analysis/step-analysis.component';

@NgModule({
  imports: [
    SharedModule,
    BlockMSRoutingModule,
    AgGridModule.withComponents([])

  ],
  declarations: [
    BlockMSComponent,
    AddNodeStepComponent,
    OnBlockStepComponent,
    GraphStepComponent,
    AnalysisStepComponent
  ]
})
export class BlockMSModule { }
